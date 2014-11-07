// An arbitrarily long big-endian integer, prefaced by a varint byte length.

package babel

import "errors"
import "math/big"
import "code.google.com/p/goprotobuf/proto"

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

func (x *BigInt) Encode () []byte {
    d := x.Data.Bytes()
    l := proto.EncodeVarint(uint64(len(d)))
    return BIGINT.Wrap(Join(l, d))
}

func DecodeBigInt (data []byte) (res Any, err error) {
    l, n := proto.DecodeVarint(data)
    if n == 0 { return nil, errors.New("ran out of bytes while parsing bigint length") }
    if int(l) > len(data) - n { return nil, errors.New("length > remaining bytes") }
    x := NewBigInt(big.NewInt(0))
    x.Data.SetBytes(data[n:n+int(l)])
    return x, nil
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