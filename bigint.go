// An arbitrarily long big-endian integer, prefaced by a varint byte length.

package babel

import "errors"
import "math/big"

var BIGINT_STRING = "8DA78674"
var BIGINT = NewTypeFromHex(BIGINT_STRING)
func (*BigInt) Type () *Type { return BIGINT }
func (*BigInt) StringType () string { return BIGINT_STRING }
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

func (x *BigInt) Encode (enc Encoding) []byte {
    return Wrap(enc, BIGINT, x.Data.Bytes())
}

func decodeBigInt (data []byte) (res Value, err error) { return DecodeBigInt(data) }
func DecodeBigInt (data []byte) (res *BigInt, err error) {
    x := NewBigInt(big.NewInt(0))
    x.Data.SetBytes(data)
    return x, nil
}

func ReadBigInt (data []byte) (res *BigInt, n int, err error) {
    l, ll := ReadVarint(data)
    if ll == 0 { return nil, 0, errors.New("ran out of bytes while parsing length") }
    end := ll + int(l)
    if end > len(data) { return nil, 0, errors.New("ran out of bytes while parsing BIGINT") }
    res, err = DecodeBigInt(data[ll:end])
    return res, end, err
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