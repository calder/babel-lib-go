// A variable length integer.

package babel

import "math/big"

var VARINT_STRING = "B5D7B812"
var VARINT = Tag(VARINT_STRING)
func (*VarInt) Type () []byte { return VARINT }
func (*VarInt) StringType () string { return VARINT_STRING }
func init () { AddType(VARINT, DecodeVarInt) }

type VarInt struct {
    Data *big.Int
}

func NewVarInt (value int) *VarInt {
    return &VarInt{big.NewInt(int64(value))}
}

func (x *VarInt) String () string {
    return "<VarInt:"+x.Data.String()+">"
}

func (x *VarInt) Encode () []byte {
    return Join(VARINT, x.Data.Bytes())
}

func DecodeVarInt (data []byte) (res Any, err error) {
    x := NewVarInt(0)
    x.Data.SetBytes(data)
    return x, nil
}

func (x *VarInt) Equal (other *VarInt) bool {
    return x.Data.Cmp(other.Data) == 0
}

func (x *VarInt) EqualAny (other Any) bool {
    switch other := other.(type) {
        case *VarInt: return x.Equal(other)
        default: return false
    }
}