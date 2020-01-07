package network_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-data-transfer/testutil"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/go-fil-components/retrievalmarket"
	"github.com/filecoin-project/go-fil-components/retrievalmarket/network"
	"github.com/filecoin-project/go-fil-components/shared/tokenamount"
	"github.com/filecoin-project/go-fil-components/shared_testutil"
)

type testReceiver struct {
	t *testing.T
	// dealStreamHandler  func(network.RetrievalDealStream)
	queryStreamHandler func(network.RetrievalQueryStream)
}

func (tr *testReceiver) HandleDealStream(s network.RetrievalDealStream) {
}
func (tr *testReceiver) HandleQueryStream(s network.RetrievalQueryStream) {
	defer s.Close()
	require.NotNil(tr.t, s)
	tr.queryStreamHandler(s)
}

func TestQueryStreamSendReceiveQuery(t *testing.T) {
	ctx := context.Background()
	td := shared_testutil.NewLibp2pTestData(ctx, t)
	nw1 := network.NewFromLibp2pHost(td.Host1)
	nw2 := network.NewFromLibp2pHost(td.Host2)
	require.NoError(t, td.Host1.Connect(ctx, peer.AddrInfo{ID: td.Host2.ID()}))

	testQueryReceived(ctx, t, nw1, nw2, td.Host2.ID())
}

func TestQueryStreamSendReceiveQueryResponse(t *testing.T) {
	ctx := context.Background()
	td := shared_testutil.NewLibp2pTestData(ctx, t)
	nw1 := network.NewFromLibp2pHost(td.Host1)
	nw2 := network.NewFromLibp2pHost(td.Host2)
	require.NoError(t, td.Host1.Connect(ctx, peer.AddrInfo{ID: td.Host2.ID()}))

	testQueryResponseReceived(ctx, t, nw1, nw2, td.Host2.ID())
}

func TestQueryStreamSendReceiveMultipleSuccessful(t *testing.T) {
	// send query, read in handler, send response back, read response
	ctxBg := context.Background()
	td := shared_testutil.NewLibp2pTestData(ctxBg, t)
	nw1 := network.NewFromLibp2pHost(td.Host1)
	nw2 := network.NewFromLibp2pHost(td.Host2)
	require.NoError(t, td.Host1.Connect(ctxBg, peer.AddrInfo{ID: td.Host2.ID()}))

	ctx, cancel := context.WithTimeout(ctxBg, 10*time.Second)
	defer cancel()

	// host1 will be getting a query response
	qrchan := make(chan retrievalmarket.QueryResponse)
	tr1 := &testReceiver{t: t, queryStreamHandler: func(s network.RetrievalQueryStream) {
		q, err := s.ReadQueryResponse()
		require.NoError(t, err)
		qrchan <- q
	}}
	require.NoError(t, nw1.SetDelegate(tr1))

	// host2 will be getting a query
	qchan := make(chan retrievalmarket.Query)
	tr2 := &testReceiver{t: t, queryStreamHandler: func(s network.RetrievalQueryStream) {
		q, err := s.ReadQuery()
		require.NoError(t, err)
		qchan <- q
	}}
	require.NoError(t, nw2.SetDelegate(tr2))

	assertQueryReceived(ctx, t, nw1, td.Host2.ID(), qchan)
	assertQueryResponseReceived(ctx, t, nw2, td.Host1.ID(), qrchan)
}

func TestQueryStreamSendReceiveOutOfOrderFails(t *testing.T) {
	// send query, read response in handler - fails
	// host2 tries to read QueryResponse instead of Query
	t.Run("sending a query and reading a response in handler fails", func(t *testing.T) {
		ctxBg := context.Background()
		td := shared_testutil.NewLibp2pTestData(ctxBg, t)
		nw1 := network.NewFromLibp2pHost(td.Host1)
		nw2 := network.NewFromLibp2pHost(td.Host2)
		require.NoError(t, td.Host1.Connect(ctxBg, peer.AddrInfo{ID: td.Host2.ID()}))

		tr := &testReceiver{t: t, queryStreamHandler: trivialQueryHandler}
		require.NoError(t, nw1.SetDelegate(tr))
		var readErr error
		errChan := make(chan error)
		tr2 := &testReceiver{t: t, queryStreamHandler: func(s network.RetrievalQueryStream) {
			_, err := s.ReadQueryResponse()
			if err != nil {
				errChan <- err
			}
		}}
		require.NoError(t, nw2.SetDelegate(tr2))
		qs1, err := nw1.NewQueryStream(td.Host2.ID())
		require.NoError(t, err)

		// send Query to host2, which tries to read a QueryResponse
		cid := testutil.GenerateCids(1)[0]
		q := retrievalmarket.NewQueryV0(cid.Bytes())
		require.NoError(t, qs1.WriteQuery(q))

		assertErrInChan(ctxBg, t, readErr, errChan, "cbor input had wrong number of fields")

	})

	t.Run("sending a QueryResponse and trying to read a Query in a handler fails", func(t *testing.T) {
		ctxBg := context.Background()
		td := shared_testutil.NewLibp2pTestData(ctxBg, t)
		nw1 := network.NewFromLibp2pHost(td.Host1)
		nw2 := network.NewFromLibp2pHost(td.Host2)
		require.NoError(t, td.Host1.Connect(ctxBg, peer.AddrInfo{ID: td.Host2.ID()}))

		tr := &testReceiver{t: t, queryStreamHandler: trivialQueryHandler}
		require.NoError(t, nw1.SetDelegate(tr))
		var readErr error
		errChan := make(chan error)
		// send response, read query in handler - fails
		tr2 := &testReceiver{t: t, queryStreamHandler: func(s network.RetrievalQueryStream) {
			_, err := s.ReadQuery()
			if err != nil {
				errChan <- err
			}
		}}
		require.NoError(t, nw2.SetDelegate(tr2))

		qs1, err := nw1.NewQueryStream(td.Host2.ID())
		require.NoError(t, err)

		require.NoError(t, qs1.WriteQueryResponse(makeTestQueryResponse()))
		assertErrInChan(ctxBg, t, readErr, errChan, "cbor input had wrong number of fields")

	})

}

func assertErrInChan(ctxBg context.Context, t *testing.T, readErr error, errChan chan error, errTxt string) {
	// wait for error to occur
	ctx, cancel := context.WithTimeout(ctxBg, 10*time.Second)
	defer cancel()
	select {
	case <-ctx.Done():
		t.Fatalf("expected error but got nothing")
	case readErr = <-errChan:
		cancel()
	}
	assert.EqualError(t, readErr, errTxt)
}

func TestDealStreamSendReceiveDealProposal(t *testing.T) {
	// send proposal, read in handler
}

func TestDealStreamSendReceiveDealResponse(t *testing.T) {
	// send response, read in handler
}

func TestDealStreamSendReceiveDealPayment(t *testing.T) {
	// send payment, read in handler
}

func TestDealStreamSendReceiveMultipleSuccessful(t *testing.T) {
	// send proposal, read in handler, send response back, read response, send payment, read farther in hander
}

func TestQueryStreamSendReceiveMultipleOutOfOrderFails(t *testing.T) {
	// send proposal, read response in handler - fails
	// send proposal, read payment in handler - fails
	// send response, read proposal in handler - fails
	// send response, read payment in handler - fails
	// send payment, read proposal in handler - fails
	// send payment, read deal in handler - fails
}

func makeTestQueryResponse() retrievalmarket.QueryResponse {
	return retrievalmarket.QueryResponse{
		Status:                     retrievalmarket.QueryResponseUnavailable,
		Size:                       66,
		PaymentAddress:             address.TestAddress,
		MinPricePerByte:            tokenamount.TokenAmount{Int: big.NewInt(77)},
		MaxPaymentInterval:         88,
		MaxPaymentIntervalIncrease: 99,
	}
}

func testQueryReceived(inCtx context.Context, t *testing.T, fromNetwork, toNetwork network.RetrievalMarketNetwork, toHost peer.ID) {
	// host1 gets no-op receiver
	tr := &testReceiver{t: t, queryStreamHandler: trivialQueryHandler}
	require.NoError(t, fromNetwork.SetDelegate(tr))

	// host2 gets receiver
	qchan := make(chan retrievalmarket.Query)
	tr2 := &testReceiver{t: t, queryStreamHandler: func(s network.RetrievalQueryStream) {
		readq, err := s.ReadQuery()
		require.NoError(t, err)
		qchan <- readq
	}}
	require.NoError(t, toNetwork.SetDelegate(tr2))

	// setup query stream host1 --> host 2
	assertQueryReceived(inCtx, t, fromNetwork, toHost, qchan)
}

func assertQueryReceived(inCtx context.Context, t *testing.T, fromNetwork network.RetrievalMarketNetwork, toHost peer.ID, qchan chan retrievalmarket.Query) {
	ctx, cancel := context.WithTimeout(inCtx, 10*time.Second)
	defer cancel()

	qs1, err := fromNetwork.NewQueryStream(toHost)
	require.NoError(t, err)

	// send query to host2
	cid := testutil.GenerateCids(1)[0]
	q := retrievalmarket.NewQueryV0(cid.Bytes())
	require.NoError(t, qs1.WriteQuery(q))

	var inq retrievalmarket.Query
	select {
	case <-ctx.Done():
		t.Fatal("msg not received")
	case inq = <-qchan:
		cancel()
	}
	require.NotNil(t, inq)
	assert.Equal(t, q.PieceCID, inq.PieceCID)
}

func testQueryResponseReceived(inCtx context.Context, t *testing.T,
	fromNetwork, toNetwork network.RetrievalMarketNetwork,
	toHost peer.ID) {
	// host1 gets no-op receiver
	tr := &testReceiver{t: t, queryStreamHandler: trivialQueryHandler}
	require.NoError(t, fromNetwork.SetDelegate(tr))

	// host2 gets receiver
	qchan := make(chan retrievalmarket.QueryResponse)
	tr2 := &testReceiver{t: t, queryStreamHandler: func(s network.RetrievalQueryStream) {
		q, err := s.ReadQueryResponse()
		require.NoError(t, err)
		qchan <- q
	}}
	require.NoError(t, toNetwork.SetDelegate(tr2))

	assertQueryResponseReceived(inCtx, t, fromNetwork, toHost, qchan)
}

func assertQueryResponseReceived(inCtx context.Context, t *testing.T,
	fromNetwork network.RetrievalMarketNetwork,
	toHost peer.ID,
	qchan chan retrievalmarket.QueryResponse) {
	ctx, cancel := context.WithTimeout(inCtx, 10*time.Second)
	defer cancel()

	// setup query stream host1 --> host 2
	qs1, err := fromNetwork.NewQueryStream(toHost)
	require.NoError(t, err)

	// send queryresponse to host2
	qr := makeTestQueryResponse()
	require.NoError(t, qs1.WriteQueryResponse(qr))

	// read queryresponse
	var inqr retrievalmarket.QueryResponse
	select {
	case <-ctx.Done():
		t.Fatal("msg not received")
	case inqr = <-qchan:
		cancel()
	}

	require.NotNil(t, inqr)
	assert.Equal(t, qr, inqr)
}

func trivialQueryHandler(s network.RetrievalQueryStream) {}
