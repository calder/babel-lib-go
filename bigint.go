// An arbitrarily long big-endian integer, prefaced by a varint byte length.

package babel

import "errors"
import "math/big"

var BIGINT_STRING = "8DA78674"
var BIGINT = NewTypeFromHex(BIGINT_STRING)
func (*BigInt) Type () *Type { return BIGINT }
func (*BigInt) StringType () string { return BIGINT_STRING }
func init () { AddType(BIGINT, DecodeBigInt) }

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

func DecodeBigInt (data []byte) (res Any, err error) {
    r, n, e := ReadBigInt(data)
    if e != nil { return nil, e }
    if n < len(data) { return nil, errors.New("leftover bytes after parsing bigint") }
    return r, e
}

func ReadBigInt (data []byte) (res *BigInt, n int, err error) {
    x := NewBigInt(big.NewInt(0))
    x.Data.SetBytes(data)
    return x, len(data), nil
}

func (x *BigInt) Equal (other *BigInt) bool {
    return x.Data.Cmp(other.Data) == 0
}

func (x *BigInt) EqualAny (other Any) bool {
    switch other := other.(type) {
        case *BigInt: return x.Equal(other)
        default: return false
    }
}