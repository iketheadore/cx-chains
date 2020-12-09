// Code generated by github.com/Skycoin/skyencoder. DO NOT EDIT.
package coin

import (
	"errors"
	"math"

	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/cipher/encoder"
)

// encodeSizeBlockBody computes the size of an encoded object of type BlockBody
func encodeSizeBlockBody(obj *BlockBody) uint64 {
	i0 := uint64(0)

	// obj.Transactions
	i0 += 4
	for _, x := range obj.Transactions {
		i1 := uint64(0)

		// x.Length
		i1 += 4

		// x.Type
		i1++

		// x.InnerHash
		i1 += 32

		// x.Sigs
		i1 += 4
		{
			i2 := uint64(0)

			// x
			i2 += 65

			i1 += uint64(len(x.Sigs)) * i2
		}

		// x.In
		i1 += 4
		{
			i2 := uint64(0)

			// x
			i2 += 32

			i1 += uint64(len(x.In)) * i2
		}

		// x.Out
		i1 += 4
		{
			i2 := uint64(0)

			// x.Address.Version
			i2++

			// x.Address.Key
			i2 += 20

			// x.Coins
			i2 += 8

			// x.Hours
			i2 += 8

			// x.ProgramState
			// WARNING: x.Out[0].ProgramState manually changed from x.ProgramState
			// WARNING: This is not considering program states in different `Out`s with different lengths
			i2 += 4 + uint64(len(x.Out[0].ProgramState))

			i1 += uint64(len(x.Out)) * i2
		}

		// x.MainExpressions
		i1 += 4 + uint64(len(x.MainExpressions))

		i0 += i1
	}

	return i0
}

// encodeBlockBody encodes an object of type BlockBody to a buffer allocated to the exact size
// required to encode the object.
func encodeBlockBody(obj *BlockBody) ([]byte, error) {
	n := encodeSizeBlockBody(obj)
	buf := make([]byte, n)

	if err := encodeBlockBodyToBuffer(buf, obj); err != nil {
		return nil, err
	}

	return buf, nil
}

// encodeBlockBodyToBuffer encodes an object of type BlockBody to a []byte buffer.
// The buffer must be large enough to encode the object, otherwise an error is returned.
func encodeBlockBodyToBuffer(buf []byte, obj *BlockBody) error {
	if uint64(len(buf)) < encodeSizeBlockBody(obj) {
		return encoder.ErrBufferUnderflow
	}

	e := &encoder.Encoder{
		Buffer: buf[:],
	}

	// obj.Transactions maxlen check
	if len(obj.Transactions) > 65535 {
		return encoder.ErrMaxLenExceeded
	}

	// obj.Transactions length check
	if uint64(len(obj.Transactions)) > math.MaxUint32 {
		return errors.New("obj.Transactions length exceeds math.MaxUint32")
	}

	// obj.Transactions length
	e.Uint32(uint32(len(obj.Transactions)))

	// obj.Transactions
	for _, x := range obj.Transactions {

		// x.Length
		e.Uint32(x.Length)

		// x.Type
		e.Uint8(x.Type)

		// x.InnerHash
		e.CopyBytes(x.InnerHash[:])

		// x.Sigs maxlen check
		if len(x.Sigs) > 65535 {
			return encoder.ErrMaxLenExceeded
		}

		// x.Sigs length check
		if uint64(len(x.Sigs)) > math.MaxUint32 {
			return errors.New("x.Sigs length exceeds math.MaxUint32")
		}

		// x.Sigs length
		e.Uint32(uint32(len(x.Sigs)))

		// x.Sigs
		for _, x := range x.Sigs {

			// x
			e.CopyBytes(x[:])

		}

		// x.In maxlen check
		if len(x.In) > 65535 {
			return encoder.ErrMaxLenExceeded
		}

		// x.In length check
		if uint64(len(x.In)) > math.MaxUint32 {
			return errors.New("x.In length exceeds math.MaxUint32")
		}

		// x.In length
		e.Uint32(uint32(len(x.In)))

		// x.In
		for _, x := range x.In {

			// x
			e.CopyBytes(x[:])

		}

		// x.Out maxlen check
		if len(x.Out) > 65535 {
			return encoder.ErrMaxLenExceeded
		}

		// x.Out length check
		if uint64(len(x.Out)) > math.MaxUint32 {
			return errors.New("x.Out length exceeds math.MaxUint32")
		}

		// x.Out length
		e.Uint32(uint32(len(x.Out)))

		// x.Out
		for _, x := range x.Out {

			// x.Address.Version
			e.Uint8(x.Address.Version)

			// x.Address.Key
			e.CopyBytes(x.Address.Key[:])

			// x.Coins
			e.Uint64(x.Coins)

			// x.Hours
			e.Uint64(x.Hours)

			// x.ProgramState length check
			if uint64(len(x.ProgramState)) > math.MaxUint32 {
				return errors.New("x.ProgramState length exceeds math.MaxUint32")
			}

			// x.ProgramState length
			e.Uint32(uint32(len(x.ProgramState)))

			// x.ProgramState copy
			e.CopyBytes(x.ProgramState)

		}

		// x.MainExpressions length check
		if uint64(len(x.MainExpressions)) > math.MaxUint32 {
			return errors.New("x.MainExpressions length exceeds math.MaxUint32")
		}

		// x.MainExpressions length
		e.Uint32(uint32(len(x.MainExpressions)))

		// x.MainExpressions copy
		e.CopyBytes(x.MainExpressions)

	}

	return nil
}

// decodeBlockBody decodes an object of type BlockBody from a buffer.
// Returns the number of bytes used from the buffer to decode the object.
// If the buffer not long enough to decode the object, returns encoder.ErrBufferUnderflow.
func decodeBlockBody(buf []byte, obj *BlockBody) (uint64, error) {
	d := &encoder.Decoder{
		Buffer: buf[:],
	}

	{
		// obj.Transactions

		ul, err := d.Uint32()
		if err != nil {
			return 0, err
		}

		length := int(ul)
		if length < 0 || length > len(d.Buffer) {
			return 0, encoder.ErrBufferUnderflow
		}

		if length > 65535 {
			return 0, encoder.ErrMaxLenExceeded
		}

		if length != 0 {
			obj.Transactions = make([]Transaction, length)

			for z1 := range obj.Transactions {
				{
					// obj.Transactions[z1].Length
					i, err := d.Uint32()
					if err != nil {
						return 0, err
					}
					obj.Transactions[z1].Length = i
				}

				{
					// obj.Transactions[z1].Type
					i, err := d.Uint8()
					if err != nil {
						return 0, err
					}
					obj.Transactions[z1].Type = i
				}

				{
					// obj.Transactions[z1].InnerHash
					if len(d.Buffer) < len(obj.Transactions[z1].InnerHash) {
						return 0, encoder.ErrBufferUnderflow
					}
					copy(obj.Transactions[z1].InnerHash[:], d.Buffer[:len(obj.Transactions[z1].InnerHash)])
					d.Buffer = d.Buffer[len(obj.Transactions[z1].InnerHash):]
				}

				{
					// obj.Transactions[z1].Sigs

					ul, err := d.Uint32()
					if err != nil {
						return 0, err
					}

					length := int(ul)
					if length < 0 || length > len(d.Buffer) {
						return 0, encoder.ErrBufferUnderflow
					}

					if length > 65535 {
						return 0, encoder.ErrMaxLenExceeded
					}

					if length != 0 {
						obj.Transactions[z1].Sigs = make([]cipher.Sig, length)

						for z3 := range obj.Transactions[z1].Sigs {
							{
								// obj.Transactions[z1].Sigs[z3]
								if len(d.Buffer) < len(obj.Transactions[z1].Sigs[z3]) {
									return 0, encoder.ErrBufferUnderflow
								}
								copy(obj.Transactions[z1].Sigs[z3][:], d.Buffer[:len(obj.Transactions[z1].Sigs[z3])])
								d.Buffer = d.Buffer[len(obj.Transactions[z1].Sigs[z3]):]
							}

						}
					}
				}

				{
					// obj.Transactions[z1].In

					ul, err := d.Uint32()
					if err != nil {
						return 0, err
					}

					length := int(ul)
					if length < 0 || length > len(d.Buffer) {
						return 0, encoder.ErrBufferUnderflow
					}

					if length > 65535 {
						return 0, encoder.ErrMaxLenExceeded
					}

					if length != 0 {
						obj.Transactions[z1].In = make([]cipher.SHA256, length)

						for z3 := range obj.Transactions[z1].In {
							{
								// obj.Transactions[z1].In[z3]
								if len(d.Buffer) < len(obj.Transactions[z1].In[z3]) {
									return 0, encoder.ErrBufferUnderflow
								}
								copy(obj.Transactions[z1].In[z3][:], d.Buffer[:len(obj.Transactions[z1].In[z3])])
								d.Buffer = d.Buffer[len(obj.Transactions[z1].In[z3]):]
							}

						}
					}
				}

				{
					// obj.Transactions[z1].Out

					ul, err := d.Uint32()
					if err != nil {
						return 0, err
					}

					length := int(ul)
					if length < 0 || length > len(d.Buffer) {
						return 0, encoder.ErrBufferUnderflow
					}

					if length > 65535 {
						return 0, encoder.ErrMaxLenExceeded
					}

					if length != 0 {
						obj.Transactions[z1].Out = make([]TransactionOutput, length)

						for z3 := range obj.Transactions[z1].Out {
							{
								// obj.Transactions[z1].Out[z3].Address.Version
								i, err := d.Uint8()
								if err != nil {
									return 0, err
								}
								obj.Transactions[z1].Out[z3].Address.Version = i
							}

							{
								// obj.Transactions[z1].Out[z3].Address.Key
								if len(d.Buffer) < len(obj.Transactions[z1].Out[z3].Address.Key) {
									return 0, encoder.ErrBufferUnderflow
								}
								copy(obj.Transactions[z1].Out[z3].Address.Key[:], d.Buffer[:len(obj.Transactions[z1].Out[z3].Address.Key)])
								d.Buffer = d.Buffer[len(obj.Transactions[z1].Out[z3].Address.Key):]
							}

							{
								// obj.Transactions[z1].Out[z3].Coins
								i, err := d.Uint64()
								if err != nil {
									return 0, err
								}
								obj.Transactions[z1].Out[z3].Coins = i
							}

							{
								// obj.Transactions[z1].Out[z3].Hours
								i, err := d.Uint64()
								if err != nil {
									return 0, err
								}
								obj.Transactions[z1].Out[z3].Hours = i
							}

							{
								// obj.Transactions[z1].Out[z3].ProgramState

								ul, err := d.Uint32()
								if err != nil {
									return 0, err
								}

								length := int(ul)
								if length < 0 || length > len(d.Buffer) {
									return 0, encoder.ErrBufferUnderflow
								}

								if length != 0 {
									obj.Transactions[z1].Out[z3].ProgramState = make([]byte, length)

									copy(obj.Transactions[z1].Out[z3].ProgramState[:], d.Buffer[:length])
									d.Buffer = d.Buffer[length:]
								}
							}
						}
					}
				}

				{
					// obj.Transactions[z1].MainExpressions

					ul, err := d.Uint32()
					if err != nil {
						return 0, err
					}

					length := int(ul)
					if length < 0 || length > len(d.Buffer) {
						return 0, encoder.ErrBufferUnderflow
					}

					if length != 0 {
						obj.Transactions[z1].MainExpressions = make([]byte, length)

						copy(obj.Transactions[z1].MainExpressions[:], d.Buffer[:length])
						d.Buffer = d.Buffer[length:]
					}
				}
			}
		}
	}

	return uint64(len(buf) - len(d.Buffer)), nil
}

// decodeBlockBodyExact decodes an object of type BlockBody from a buffer.
// If the buffer not long enough to decode the object, returns encoder.ErrBufferUnderflow.
// If the buffer is longer than required to decode the object, returns encoder.ErrRemainingBytes.
func decodeBlockBodyExact(buf []byte, obj *BlockBody) error {
	if n, err := decodeBlockBody(buf, obj); err != nil {
		return err
	} else if n != uint64(len(buf)) {
		return encoder.ErrRemainingBytes
	}

	return nil
}
