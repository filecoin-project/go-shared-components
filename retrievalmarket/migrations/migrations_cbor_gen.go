// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package migrations

import (
	"fmt"
	"io"
	"sort"

	migrations "github.com/filecoin-project/go-fil-markets/piecestore/migrations"
	retrievalmarket "github.com/filecoin-project/go-fil-markets/retrievalmarket"
	multistore "github.com/filecoin-project/go-multistore"
	paych "github.com/filecoin-project/specs-actors/actors/builtin/paych"
	cid "github.com/ipfs/go-cid"
	peer "github.com/libp2p/go-libp2p-core/peer"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = sort.Sort

var lengthBufQuery0 = []byte{130}

func (t *Query0) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufQuery0); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.PayloadCID (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.PayloadCID); err != nil {
		return xerrors.Errorf("failed to write cid field t.PayloadCID: %w", err)
	}

	// t.QueryParams0 (migrations.QueryParams0) (struct)
	if err := t.QueryParams0.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *Query0) UnmarshalCBOR(r io.Reader) error {
	*t = Query0{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.PayloadCID (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.PayloadCID: %w", err)
		}

		t.PayloadCID = c

	}
	// t.QueryParams0 (migrations.QueryParams0) (struct)

	{

		if err := t.QueryParams0.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.QueryParams0: %w", err)
		}

	}
	return nil
}

var lengthBufQueryResponse0 = []byte{137}

func (t *QueryResponse0) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufQueryResponse0); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Status (retrievalmarket.QueryResponseStatus) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Status)); err != nil {
		return err
	}

	// t.PieceCIDFound (retrievalmarket.QueryItemStatus) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.PieceCIDFound)); err != nil {
		return err
	}

	// t.Size (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Size)); err != nil {
		return err
	}

	// t.PaymentAddress (address.Address) (struct)
	if err := t.PaymentAddress.MarshalCBOR(w); err != nil {
		return err
	}

	// t.MinPricePerByte (big.Int) (struct)
	if err := t.MinPricePerByte.MarshalCBOR(w); err != nil {
		return err
	}

	// t.MaxPaymentInterval (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.MaxPaymentInterval)); err != nil {
		return err
	}

	// t.MaxPaymentIntervalIncrease (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.MaxPaymentIntervalIncrease)); err != nil {
		return err
	}

	// t.Message (string) (string)
	if len(t.Message) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Message was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.Message))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Message)); err != nil {
		return err
	}

	// t.UnsealPrice (big.Int) (struct)
	if err := t.UnsealPrice.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *QueryResponse0) UnmarshalCBOR(r io.Reader) error {
	*t = QueryResponse0{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 9 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Status (retrievalmarket.QueryResponseStatus) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Status = retrievalmarket.QueryResponseStatus(extra)

	}
	// t.PieceCIDFound (retrievalmarket.QueryItemStatus) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.PieceCIDFound = retrievalmarket.QueryItemStatus(extra)

	}
	// t.Size (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Size = uint64(extra)

	}
	// t.PaymentAddress (address.Address) (struct)

	{

		if err := t.PaymentAddress.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.PaymentAddress: %w", err)
		}

	}
	// t.MinPricePerByte (big.Int) (struct)

	{

		if err := t.MinPricePerByte.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.MinPricePerByte: %w", err)
		}

	}
	// t.MaxPaymentInterval (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.MaxPaymentInterval = uint64(extra)

	}
	// t.MaxPaymentIntervalIncrease (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.MaxPaymentIntervalIncrease = uint64(extra)

	}
	// t.Message (string) (string)

	{
		sval, err := cbg.ReadStringBuf(br, scratch)
		if err != nil {
			return err
		}

		t.Message = string(sval)
	}
	// t.UnsealPrice (big.Int) (struct)

	{

		if err := t.UnsealPrice.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.UnsealPrice: %w", err)
		}

	}
	return nil
}

var lengthBufDealProposal0 = []byte{131}

func (t *DealProposal0) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufDealProposal0); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.PayloadCID (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.PayloadCID); err != nil {
		return xerrors.Errorf("failed to write cid field t.PayloadCID: %w", err)
	}

	// t.ID (retrievalmarket.DealID) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.ID)); err != nil {
		return err
	}

	// t.Params0 (migrations.Params0) (struct)
	if err := t.Params0.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *DealProposal0) UnmarshalCBOR(r io.Reader) error {
	*t = DealProposal0{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.PayloadCID (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.PayloadCID: %w", err)
		}

		t.PayloadCID = c

	}
	// t.ID (retrievalmarket.DealID) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.ID = retrievalmarket.DealID(extra)

	}
	// t.Params0 (migrations.Params0) (struct)

	{

		if err := t.Params0.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Params0: %w", err)
		}

	}
	return nil
}

var lengthBufDealResponse0 = []byte{132}

func (t *DealResponse0) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufDealResponse0); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Status (retrievalmarket.DealStatus) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Status)); err != nil {
		return err
	}

	// t.ID (retrievalmarket.DealID) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.ID)); err != nil {
		return err
	}

	// t.PaymentOwed (big.Int) (struct)
	if err := t.PaymentOwed.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Message (string) (string)
	if len(t.Message) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Message was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.Message))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Message)); err != nil {
		return err
	}
	return nil
}

func (t *DealResponse0) UnmarshalCBOR(r io.Reader) error {
	*t = DealResponse0{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 4 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Status (retrievalmarket.DealStatus) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Status = retrievalmarket.DealStatus(extra)

	}
	// t.ID (retrievalmarket.DealID) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.ID = retrievalmarket.DealID(extra)

	}
	// t.PaymentOwed (big.Int) (struct)

	{

		if err := t.PaymentOwed.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.PaymentOwed: %w", err)
		}

	}
	// t.Message (string) (string)

	{
		sval, err := cbg.ReadStringBuf(br, scratch)
		if err != nil {
			return err
		}

		t.Message = string(sval)
	}
	return nil
}

var lengthBufParams0 = []byte{134}

func (t *Params0) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufParams0); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Selector (typegen.Deferred) (struct)
	if err := t.Selector.MarshalCBOR(w); err != nil {
		return err
	}

	// t.PieceCID (cid.Cid) (struct)

	if t.PieceCID == nil {
		if _, err := w.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCidBuf(scratch, w, *t.PieceCID); err != nil {
			return xerrors.Errorf("failed to write cid field t.PieceCID: %w", err)
		}
	}

	// t.PricePerByte (big.Int) (struct)
	if err := t.PricePerByte.MarshalCBOR(w); err != nil {
		return err
	}

	// t.PaymentInterval (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.PaymentInterval)); err != nil {
		return err
	}

	// t.PaymentIntervalIncrease (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.PaymentIntervalIncrease)); err != nil {
		return err
	}

	// t.UnsealPrice (big.Int) (struct)
	if err := t.UnsealPrice.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *Params0) UnmarshalCBOR(r io.Reader) error {
	*t = Params0{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 6 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Selector (typegen.Deferred) (struct)

	{

		t.Selector = new(cbg.Deferred)

		if err := t.Selector.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("failed to read deferred field: %w", err)
		}
	}
	// t.PieceCID (cid.Cid) (struct)

	{

		b, err := br.ReadByte()
		if err != nil {
			return err
		}
		if b != cbg.CborNull[0] {
			if err := br.UnreadByte(); err != nil {
				return err
			}

			c, err := cbg.ReadCid(br)
			if err != nil {
				return xerrors.Errorf("failed to read cid field t.PieceCID: %w", err)
			}

			t.PieceCID = &c
		}

	}
	// t.PricePerByte (big.Int) (struct)

	{

		if err := t.PricePerByte.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.PricePerByte: %w", err)
		}

	}
	// t.PaymentInterval (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.PaymentInterval = uint64(extra)

	}
	// t.PaymentIntervalIncrease (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.PaymentIntervalIncrease = uint64(extra)

	}
	// t.UnsealPrice (big.Int) (struct)

	{

		if err := t.UnsealPrice.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.UnsealPrice: %w", err)
		}

	}
	return nil
}

var lengthBufQueryParams0 = []byte{129}

func (t *QueryParams0) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufQueryParams0); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.PieceCID (cid.Cid) (struct)

	if t.PieceCID == nil {
		if _, err := w.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCidBuf(scratch, w, *t.PieceCID); err != nil {
			return xerrors.Errorf("failed to write cid field t.PieceCID: %w", err)
		}
	}

	return nil
}

func (t *QueryParams0) UnmarshalCBOR(r io.Reader) error {
	*t = QueryParams0{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.PieceCID (cid.Cid) (struct)

	{

		b, err := br.ReadByte()
		if err != nil {
			return err
		}
		if b != cbg.CborNull[0] {
			if err := br.UnreadByte(); err != nil {
				return err
			}

			c, err := cbg.ReadCid(br)
			if err != nil {
				return xerrors.Errorf("failed to read cid field t.PieceCID: %w", err)
			}

			t.PieceCID = &c
		}

	}
	return nil
}

var lengthBufDealPayment0 = []byte{131}

func (t *DealPayment0) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufDealPayment0); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.ID (retrievalmarket.DealID) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.ID)); err != nil {
		return err
	}

	// t.PaymentChannel (address.Address) (struct)
	if err := t.PaymentChannel.MarshalCBOR(w); err != nil {
		return err
	}

	// t.PaymentVoucher (paych.SignedVoucher) (struct)
	if err := t.PaymentVoucher.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *DealPayment0) UnmarshalCBOR(r io.Reader) error {
	*t = DealPayment0{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.ID (retrievalmarket.DealID) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.ID = retrievalmarket.DealID(extra)

	}
	// t.PaymentChannel (address.Address) (struct)

	{

		if err := t.PaymentChannel.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.PaymentChannel: %w", err)
		}

	}
	// t.PaymentVoucher (paych.SignedVoucher) (struct)

	{

		b, err := br.ReadByte()
		if err != nil {
			return err
		}
		if b != cbg.CborNull[0] {
			if err := br.UnreadByte(); err != nil {
				return err
			}
			t.PaymentVoucher = new(paych.SignedVoucher)
			if err := t.PaymentVoucher.UnmarshalCBOR(br); err != nil {
				return xerrors.Errorf("unmarshaling t.PaymentVoucher pointer: %w", err)
			}
		}

	}
	return nil
}

var lengthBufClientDealState0 = []byte{148}

func (t *ClientDealState0) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufClientDealState0); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.DealProposal0 (migrations.DealProposal0) (struct)
	if err := t.DealProposal0.MarshalCBOR(w); err != nil {
		return err
	}

	// t.StoreID (multistore.StoreID) (uint64)

	if t.StoreID == nil {
		if _, err := w.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(*t.StoreID)); err != nil {
			return err
		}
	}

	// t.ChannelID (datatransfer.ChannelID) (struct)
	if err := t.ChannelID.MarshalCBOR(w); err != nil {
		return err
	}

	// t.LastPaymentRequested (bool) (bool)
	if err := cbg.WriteBool(w, t.LastPaymentRequested); err != nil {
		return err
	}

	// t.AllBlocksReceived (bool) (bool)
	if err := cbg.WriteBool(w, t.AllBlocksReceived); err != nil {
		return err
	}

	// t.TotalFunds (big.Int) (struct)
	if err := t.TotalFunds.MarshalCBOR(w); err != nil {
		return err
	}

	// t.ClientWallet (address.Address) (struct)
	if err := t.ClientWallet.MarshalCBOR(w); err != nil {
		return err
	}

	// t.MinerWallet (address.Address) (struct)
	if err := t.MinerWallet.MarshalCBOR(w); err != nil {
		return err
	}

	// t.PaymentInfo (migrations.PaymentInfo0) (struct)
	if err := t.PaymentInfo.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Status (retrievalmarket.DealStatus) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Status)); err != nil {
		return err
	}

	// t.Sender (peer.ID) (string)
	if len(t.Sender) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Sender was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.Sender))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Sender)); err != nil {
		return err
	}

	// t.TotalReceived (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.TotalReceived)); err != nil {
		return err
	}

	// t.Message (string) (string)
	if len(t.Message) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Message was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.Message))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Message)); err != nil {
		return err
	}

	// t.BytesPaidFor (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.BytesPaidFor)); err != nil {
		return err
	}

	// t.CurrentInterval (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.CurrentInterval)); err != nil {
		return err
	}

	// t.PaymentRequested (big.Int) (struct)
	if err := t.PaymentRequested.MarshalCBOR(w); err != nil {
		return err
	}

	// t.FundsSpent (big.Int) (struct)
	if err := t.FundsSpent.MarshalCBOR(w); err != nil {
		return err
	}

	// t.UnsealFundsPaid (big.Int) (struct)
	if err := t.UnsealFundsPaid.MarshalCBOR(w); err != nil {
		return err
	}

	// t.WaitMsgCID (cid.Cid) (struct)

	if t.WaitMsgCID == nil {
		if _, err := w.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCidBuf(scratch, w, *t.WaitMsgCID); err != nil {
			return xerrors.Errorf("failed to write cid field t.WaitMsgCID: %w", err)
		}
	}

	// t.VoucherShortfall (big.Int) (struct)
	if err := t.VoucherShortfall.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *ClientDealState0) UnmarshalCBOR(r io.Reader) error {
	*t = ClientDealState0{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 20 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.DealProposal0 (migrations.DealProposal0) (struct)

	{

		if err := t.DealProposal0.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.DealProposal0: %w", err)
		}

	}
	// t.StoreID (multistore.StoreID) (uint64)

	{

		b, err := br.ReadByte()
		if err != nil {
			return err
		}
		if b != cbg.CborNull[0] {
			if err := br.UnreadByte(); err != nil {
				return err
			}
			maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
			if err != nil {
				return err
			}
			if maj != cbg.MajUnsignedInt {
				return fmt.Errorf("wrong type for uint64 field")
			}
			typed := multistore.StoreID(extra)
			t.StoreID = &typed
		}

	}
	// t.ChannelID (datatransfer.ChannelID) (struct)

	{

		if err := t.ChannelID.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.ChannelID: %w", err)
		}

	}
	// t.LastPaymentRequested (bool) (bool)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajOther {
		return fmt.Errorf("booleans must be major type 7")
	}
	switch extra {
	case 20:
		t.LastPaymentRequested = false
	case 21:
		t.LastPaymentRequested = true
	default:
		return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
	}
	// t.AllBlocksReceived (bool) (bool)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajOther {
		return fmt.Errorf("booleans must be major type 7")
	}
	switch extra {
	case 20:
		t.AllBlocksReceived = false
	case 21:
		t.AllBlocksReceived = true
	default:
		return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
	}
	// t.TotalFunds (big.Int) (struct)

	{

		if err := t.TotalFunds.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.TotalFunds: %w", err)
		}

	}
	// t.ClientWallet (address.Address) (struct)

	{

		if err := t.ClientWallet.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.ClientWallet: %w", err)
		}

	}
	// t.MinerWallet (address.Address) (struct)

	{

		if err := t.MinerWallet.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.MinerWallet: %w", err)
		}

	}
	// t.PaymentInfo (migrations.PaymentInfo0) (struct)

	{

		b, err := br.ReadByte()
		if err != nil {
			return err
		}
		if b != cbg.CborNull[0] {
			if err := br.UnreadByte(); err != nil {
				return err
			}
			t.PaymentInfo = new(PaymentInfo0)
			if err := t.PaymentInfo.UnmarshalCBOR(br); err != nil {
				return xerrors.Errorf("unmarshaling t.PaymentInfo pointer: %w", err)
			}
		}

	}
	// t.Status (retrievalmarket.DealStatus) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Status = retrievalmarket.DealStatus(extra)

	}
	// t.Sender (peer.ID) (string)

	{
		sval, err := cbg.ReadStringBuf(br, scratch)
		if err != nil {
			return err
		}

		t.Sender = peer.ID(sval)
	}
	// t.TotalReceived (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.TotalReceived = uint64(extra)

	}
	// t.Message (string) (string)

	{
		sval, err := cbg.ReadStringBuf(br, scratch)
		if err != nil {
			return err
		}

		t.Message = string(sval)
	}
	// t.BytesPaidFor (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.BytesPaidFor = uint64(extra)

	}
	// t.CurrentInterval (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.CurrentInterval = uint64(extra)

	}
	// t.PaymentRequested (big.Int) (struct)

	{

		if err := t.PaymentRequested.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.PaymentRequested: %w", err)
		}

	}
	// t.FundsSpent (big.Int) (struct)

	{

		if err := t.FundsSpent.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.FundsSpent: %w", err)
		}

	}
	// t.UnsealFundsPaid (big.Int) (struct)

	{

		if err := t.UnsealFundsPaid.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.UnsealFundsPaid: %w", err)
		}

	}
	// t.WaitMsgCID (cid.Cid) (struct)

	{

		b, err := br.ReadByte()
		if err != nil {
			return err
		}
		if b != cbg.CborNull[0] {
			if err := br.UnreadByte(); err != nil {
				return err
			}

			c, err := cbg.ReadCid(br)
			if err != nil {
				return xerrors.Errorf("failed to read cid field t.WaitMsgCID: %w", err)
			}

			t.WaitMsgCID = &c
		}

	}
	// t.VoucherShortfall (big.Int) (struct)

	{

		if err := t.VoucherShortfall.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.VoucherShortfall: %w", err)
		}

	}
	return nil
}

var lengthBufProviderDealState0 = []byte{138}

func (t *ProviderDealState0) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufProviderDealState0); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.DealProposal0 (migrations.DealProposal0) (struct)
	if err := t.DealProposal0.MarshalCBOR(w); err != nil {
		return err
	}

	// t.StoreID (multistore.StoreID) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.StoreID)); err != nil {
		return err
	}

	// t.ChannelID (datatransfer.ChannelID) (struct)
	if err := t.ChannelID.MarshalCBOR(w); err != nil {
		return err
	}

	// t.PieceInfo (migrations.PieceInfo0) (struct)
	if err := t.PieceInfo.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Status (retrievalmarket.DealStatus) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Status)); err != nil {
		return err
	}

	// t.Receiver (peer.ID) (string)
	if len(t.Receiver) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Receiver was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.Receiver))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Receiver)); err != nil {
		return err
	}

	// t.TotalSent (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.TotalSent)); err != nil {
		return err
	}

	// t.FundsReceived (big.Int) (struct)
	if err := t.FundsReceived.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Message (string) (string)
	if len(t.Message) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Message was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.Message))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Message)); err != nil {
		return err
	}

	// t.CurrentInterval (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.CurrentInterval)); err != nil {
		return err
	}

	return nil
}

func (t *ProviderDealState0) UnmarshalCBOR(r io.Reader) error {
	*t = ProviderDealState0{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 10 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.DealProposal0 (migrations.DealProposal0) (struct)

	{

		if err := t.DealProposal0.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.DealProposal0: %w", err)
		}

	}
	// t.StoreID (multistore.StoreID) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.StoreID = multistore.StoreID(extra)

	}
	// t.ChannelID (datatransfer.ChannelID) (struct)

	{

		if err := t.ChannelID.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.ChannelID: %w", err)
		}

	}
	// t.PieceInfo (migrations.PieceInfo0) (struct)

	{

		b, err := br.ReadByte()
		if err != nil {
			return err
		}
		if b != cbg.CborNull[0] {
			if err := br.UnreadByte(); err != nil {
				return err
			}
			t.PieceInfo = new(migrations.PieceInfo0)
			if err := t.PieceInfo.UnmarshalCBOR(br); err != nil {
				return xerrors.Errorf("unmarshaling t.PieceInfo pointer: %w", err)
			}
		}

	}
	// t.Status (retrievalmarket.DealStatus) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Status = retrievalmarket.DealStatus(extra)

	}
	// t.Receiver (peer.ID) (string)

	{
		sval, err := cbg.ReadStringBuf(br, scratch)
		if err != nil {
			return err
		}

		t.Receiver = peer.ID(sval)
	}
	// t.TotalSent (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.TotalSent = uint64(extra)

	}
	// t.FundsReceived (big.Int) (struct)

	{

		if err := t.FundsReceived.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.FundsReceived: %w", err)
		}

	}
	// t.Message (string) (string)

	{
		sval, err := cbg.ReadStringBuf(br, scratch)
		if err != nil {
			return err
		}

		t.Message = string(sval)
	}
	// t.CurrentInterval (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.CurrentInterval = uint64(extra)

	}
	return nil
}

var lengthBufPaymentInfo0 = []byte{130}

func (t *PaymentInfo0) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufPaymentInfo0); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.PayCh (address.Address) (struct)
	if err := t.PayCh.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Lane (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Lane)); err != nil {
		return err
	}

	return nil
}

func (t *PaymentInfo0) UnmarshalCBOR(r io.Reader) error {
	*t = PaymentInfo0{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.PayCh (address.Address) (struct)

	{

		if err := t.PayCh.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.PayCh: %w", err)
		}

	}
	// t.Lane (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Lane = uint64(extra)

	}
	return nil
}

var lengthBufRetrievalPeer0 = []byte{131}

func (t *RetrievalPeer0) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufRetrievalPeer0); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Address (address.Address) (struct)
	if err := t.Address.MarshalCBOR(w); err != nil {
		return err
	}

	// t.ID (peer.ID) (string)
	if len(t.ID) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.ID was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.ID))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.ID)); err != nil {
		return err
	}

	// t.PieceCID (cid.Cid) (struct)

	if t.PieceCID == nil {
		if _, err := w.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCidBuf(scratch, w, *t.PieceCID); err != nil {
			return xerrors.Errorf("failed to write cid field t.PieceCID: %w", err)
		}
	}

	return nil
}

func (t *RetrievalPeer0) UnmarshalCBOR(r io.Reader) error {
	*t = RetrievalPeer0{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Address (address.Address) (struct)

	{

		if err := t.Address.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Address: %w", err)
		}

	}
	// t.ID (peer.ID) (string)

	{
		sval, err := cbg.ReadStringBuf(br, scratch)
		if err != nil {
			return err
		}

		t.ID = peer.ID(sval)
	}
	// t.PieceCID (cid.Cid) (struct)

	{

		b, err := br.ReadByte()
		if err != nil {
			return err
		}
		if b != cbg.CborNull[0] {
			if err := br.UnreadByte(); err != nil {
				return err
			}

			c, err := cbg.ReadCid(br)
			if err != nil {
				return xerrors.Errorf("failed to read cid field t.PieceCID: %w", err)
			}

			t.PieceCID = &c
		}

	}
	return nil
}

var lengthBufAsk0 = []byte{132}

func (t *Ask0) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufAsk0); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.PricePerByte (big.Int) (struct)
	if err := t.PricePerByte.MarshalCBOR(w); err != nil {
		return err
	}

	// t.UnsealPrice (big.Int) (struct)
	if err := t.UnsealPrice.MarshalCBOR(w); err != nil {
		return err
	}

	// t.PaymentInterval (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.PaymentInterval)); err != nil {
		return err
	}

	// t.PaymentIntervalIncrease (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.PaymentIntervalIncrease)); err != nil {
		return err
	}

	return nil
}

func (t *Ask0) UnmarshalCBOR(r io.Reader) error {
	*t = Ask0{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 4 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.PricePerByte (big.Int) (struct)

	{

		if err := t.PricePerByte.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.PricePerByte: %w", err)
		}

	}
	// t.UnsealPrice (big.Int) (struct)

	{

		if err := t.UnsealPrice.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.UnsealPrice: %w", err)
		}

	}
	// t.PaymentInterval (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.PaymentInterval = uint64(extra)

	}
	// t.PaymentIntervalIncrease (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.PaymentIntervalIncrease = uint64(extra)

	}
	return nil
}
