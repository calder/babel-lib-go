// An arbitrarily long big-endian integer, prefaced by a varint byte length.

package babel

import "math/big"

var BIGINT = Type("8DA78674")
func (*BigInt) Type () uint64 { return BIGINT }
func (*BigInt) TypeName () string { return "BigInt" }
func init () { AddType(BIGINT, decodeBigInt) }

type BigInt struct {
    Data *big.Int
}

func NewBigInt (value *big.Int) *BigInt {
    return &BigInt{value}
}

func (x *BigInt) String () string {
    return "<BigInt:"+x.RawString()+">"
}

func (x *BigInt) RawString () string {
    return x.Data.String()
}

func (x *BigInt) CBR () []byte {
    return x.Encode(TYPE)
}

func (x *BigInt) Encode (enc Encoding) []byte {
    return Wrap(enc, BIGINT, x.Data.Bytes())
}

func decodeBigInt (data []byte) (res Value, err error) { return DecodeBigInt(data) }
func DecodeBigInt (data []byte) (res *BigInt, err error) {
    x := NewBigInt(big.NewInt(0))
    x.Data.SetBytes(data)
    return x, nil
}

func ReadBigInt (bytes []byte) (res *BigInt, n int, err error) {
    _, data, n, err := Unwrap(LEN, bytes)
    if err != nil { return nil, 0, err }
    res, err = DecodeBigInt(data)
    return res, n, err
}

func (x *BigInt) Equal (other *BigInt) bool {
    return x.Data.Cmp(other.Data) == 0
}

func (x *BigInt) EqualValue (other Value) bool {
    switch other := other.(type) {
        case *BigInt: return x.Equal(other)
        default: return false
    }
}
