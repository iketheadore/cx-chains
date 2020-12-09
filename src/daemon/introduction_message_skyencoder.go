// Code generated by github.com/Skycoin/skyencoder. DO NOT EDIT.
package daemon

import (
	"errors"
	"math"

	"github.com/skycoin/skycoin/src/cipher/encoder"
)

// encodeSizeIntroductionMessage computes the size of an encoded object of type IntroductionMessage
func encodeSizeIntroductionMessage(obj *IntroductionMessage) uint64 {
	i0 := uint64(0)

	// obj.Mirror
	i0 += 4

	// obj.ListenPort
	i0 += 2

	// obj.ProtocolVersion
	i0 += 4

	// omitempty
	if len(obj.Extra) != 0 {

		// obj.Extra
		i0 += 4 + uint64(len(obj.Extra))

	}

	return i0
}

// encodeIntroductionMessage encodes an object of type IntroductionMessage to a buffer allocated to the exact size
// required to encode the object.
func encodeIntroductionMessage(obj *IntroductionMessage) ([]byte, error) {
	n := encodeSizeIntroductionMessage(obj)
	buf := make([]byte, n)

	if err := encodeIntroductionMessageToBuffer(buf, obj); err != nil {
		return nil, err
	}

	return buf, nil
}

// encodeIntroductionMessageToBuffer encodes an object of type IntroductionMessage to a []byte buffer.
// The buffer must be large enough to encode the object, otherwise an error is returned.
func encodeIntroductionMessageToBuffer(buf []byte, obj *IntroductionMessage) error {
	if uint64(len(buf)) < encodeSizeIntroductionMessage(obj) {
		return encoder.ErrBufferUnderflow
	}

	e := &encoder.Encoder{
		Buffer: buf[:],
	}

	// obj.Mirror
	e.Uint32(obj.Mirror)

	// obj.ListenPort
	e.Uint16(obj.ListenPort)

	// obj.ProtocolVersion
	e.Int32(obj.ProtocolVersion)

	// omitempty
	if len(obj.Extra) != 0 {

		// obj.Extra length check
		if uint64(len(obj.Extra)) > math.MaxUint32 {
			return errors.New("obj.Extra length exceeds math.MaxUint32")
		}

		// obj.Extra length
		e.Uint32(uint32(len(obj.Extra)))

		// obj.Extra copy
		e.CopyBytes(obj.Extra)

	}

	return nil
}

// decodeIntroductionMessage decodes an object of type IntroductionMessage from a buffer.
// Returns the number of bytes used from the buffer to decode the object.
// If the buffer not long enough to decode the object, returns encoder.ErrBufferUnderflow.
func decodeIntroductionMessage(buf []byte, obj *IntroductionMessage) (uint64, error) {
	d := &encoder.Decoder{
		Buffer: buf[:],
	}

	{
		// obj.Mirror
		i, err := d.Uint32()
		if err != nil {
			return 0, err
		}
		obj.Mirror = i
	}

	{
		// obj.ListenPort
		i, err := d.Uint16()
		if err != nil {
			return 0, err
		}
		obj.ListenPort = i
	}

	{
		// obj.ProtocolVersion
		i, err := d.Int32()
		if err != nil {
			return 0, err
		}
		obj.ProtocolVersion = i
	}

	{
		// obj.Extra

		if len(d.Buffer) == 0 {
			return uint64(len(buf) - len(d.Buffer)), nil
		}

		ul, err := d.Uint32()
		if err != nil {
			return 0, err
		}

		length := int(ul)
		if length < 0 || length > len(d.Buffer) {
			return 0, encoder.ErrBufferUnderflow
		}

		if length != 0 {
			obj.Extra = make([]byte, length)

			copy(obj.Extra[:], d.Buffer[:length])
			d.Buffer = d.Buffer[length:]
		}
	}

	return uint64(len(buf) - len(d.Buffer)), nil
}

// decodeIntroductionMessageExact decodes an object of type IntroductionMessage from a buffer.
// If the buffer not long enough to decode the object, returns encoder.ErrBufferUnderflow.
// If the buffer is longer than required to decode the object, returns encoder.ErrRemainingBytes.
func decodeIntroductionMessageExact(buf []byte, obj *IntroductionMessage) error {
	if n, err := decodeIntroductionMessage(buf, obj); err != nil {
		return err
	} else if n != uint64(len(buf)) {
		return encoder.ErrRemainingBytes
	}

	return nil
}
