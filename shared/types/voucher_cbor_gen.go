package types

import (
	"fmt"
	"io"

	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

var _ = xerrors.Errorf

func (t *SignedVoucher) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{137}); err != nil {
		return err
	}

	// t.TimeLock (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.TimeLock))); err != nil {
		return err
	}

	// t.SecretPreimage ([]uint8) (slice)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(t.SecretPreimage)))); err != nil {
		return err
	}
	if _, err := w.Write(t.SecretPreimage); err != nil {
		return err
	}

	// t.Extra (types.ModVerifyParams) (struct)
	if err := t.Extra.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Lane (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Lane))); err != nil {
		return err
	}

	// t.Nonce (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Nonce))); err != nil {
		return err
	}

	// t.Amount (tokenamount.TokenAmount) (struct)
	if err := t.Amount.MarshalCBOR(w); err != nil {
		return err
	}

	// t.MinCloseHeight (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.MinCloseHeight))); err != nil {
		return err
	}

	// t.Merges ([]types.Merge) (slice)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(t.Merges)))); err != nil {
		return err
	}
	for _, v := range t.Merges {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}

	// t.Signature (types.Signature) (struct)
	if err := t.Signature.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *SignedVoucher) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 9 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.TimeLock (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.TimeLock = uint64(extra)
	// t.SecretPreimage ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.SecretPreimage: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}
	t.SecretPreimage = make([]byte, extra)
	if _, err := io.ReadFull(br, t.SecretPreimage); err != nil {
		return err
	}
	// t.Extra (types.ModVerifyParams) (struct)

	{

		pb, err := br.PeekByte()
		if err != nil {
			return err
		}
		if pb == cbg.CborNull[0] {
			var nbuf [1]byte
			if _, err := br.Read(nbuf[:]); err != nil {
				return err
			}
		} else {
			t.Extra = new(ModVerifyParams)
			if err := t.Extra.UnmarshalCBOR(br); err != nil {
				return err
			}
		}

	}
	// t.Lane (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.Lane = uint64(extra)
	// t.Nonce (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.Nonce = uint64(extra)
	// t.Amount (tokenamount.TokenAmount) (struct)

	{

		if err := t.Amount.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.MinCloseHeight (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.MinCloseHeight = uint64(extra)
	// t.Merges ([]types.Merge) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Merges: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}
	if extra > 0 {
		t.Merges = make([]Merge, extra)
	}
	for i := 0; i < int(extra); i++ {

		var v Merge
		if err := v.UnmarshalCBOR(br); err != nil {
			return err
		}

		t.Merges[i] = v
	}

	// t.Signature (types.Signature) (struct)

	{

		pb, err := br.PeekByte()
		if err != nil {
			return err
		}
		if pb == cbg.CborNull[0] {
			var nbuf [1]byte
			if _, err := br.Read(nbuf[:]); err != nil {
				return err
			}
		} else {
			t.Signature = new(Signature)
			if err := t.Signature.UnmarshalCBOR(br); err != nil {
				return err
			}
		}

	}
	return nil
}

func (t *ModVerifyParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{131}); err != nil {
		return err
	}

	// t.Actor (address.Address) (struct)
	if err := t.Actor.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Method (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Method))); err != nil {
		return err
	}

	// t.Data ([]uint8) (slice)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(t.Data)))); err != nil {
		return err
	}
	if _, err := w.Write(t.Data); err != nil {
		return err
	}
	return nil
}

func (t *ModVerifyParams) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Actor (address.Address) (struct)

	{

		if err := t.Actor.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.Method (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.Method = uint64(extra)
	// t.Data ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Data: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}
	t.Data = make([]byte, extra)
	if _, err := io.ReadFull(br, t.Data); err != nil {
		return err
	}
	return nil
}

func (t *Merge) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{130}); err != nil {
		return err
	}

	// t.Lane (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Lane))); err != nil {
		return err
	}

	// t.Nonce (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Nonce))); err != nil {
		return err
	}
	return nil
}

func (t *Merge) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Lane (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.Lane = uint64(extra)
	// t.Nonce (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.Nonce = uint64(extra)
	return nil
}
