package sectorbuilder_test

import (
	"io"
	"io/ioutil"
	"math/rand"
	"testing"

	"github.com/filecoin-project/go-lotus/chain/address"
	"github.com/filecoin-project/go-lotus/lib/sectorbuilder"
	"github.com/filecoin-project/go-lotus/storage/sector"
)

func TestSealAndVerify(t *testing.T) {
	t.Skip("this is slow")
	dir, err := ioutil.TempDir("", "sbtest")
	if err != nil {
		t.Fatal(err)
	}

	addr, err := address.NewFromString("t1tct3nfaw2q543xtybxcyw4deyxmfwkjk43u4t5y")
	if err != nil {
		t.Fatal(err)
	}

	sb, err := sectorbuilder.New(&sectorbuilder.SectorBuilderConfig{
		SectorSize:  1024,
		SealedDir:   dir,
		StagedDir:   dir,
		MetadataDir: dir,
		Miner:       addr,
	})
	if err != nil {
		t.Fatal(err)
	}

	fi, err := ioutil.TempFile("", "sbtestfi")
	if err != nil {
		t.Fatal(err)
	}
	defer fi.Close()

	io.CopyN(fi, rand.New(rand.NewSource(42)), 1016)

	if _, err := sb.AddPiece("foo", 1016, fi.Name()); err != nil {
		t.Fatal(err)
	}

	store := sector.NewStore(sb)
	store.Service()
	ssinfo := <-store.Incoming()

	ok, err := sectorbuilder.VerifySeal(1024, ssinfo.CommR[:], ssinfo.CommD[:], ssinfo.CommRStar[:], addr, ssinfo.SectorID, ssinfo.Proof)
	if err != nil {
		t.Fatal(err)
	}

	if !ok {
		t.Fatal("proof failed to validate")
	}
}
