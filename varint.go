// A protobuf compatible variable length integer.
// 
// The first bit of each byte determines whether to read the next byte.
// The last seven bits of each byte are the value, least significant byte first.
//     1******* 1******* 1******* 0*******
//      ^^^^^^^  ^^^^^^^  ^^^^^^^  ^^^^^^^
//   bits 22-28   15-21    7-14      0-6

package babel

import "errors"
import "strconv"
import "code.google.com/p/goprotobuf/proto"

var VARINT_STRING = "B5D7B812"
var VARINT = NewTypeFromHex(VARINT_STRING)
func (*VarInt) Type () *Type { return VARINT }
func (*VarInt) StringType () string { return VARINT_STRING }
func init () { AddType(VARINT, DecodeVarInt) }

type VarInt struct {
    Data uint64
}

func NewVarInt (value uint64) *VarInt {
    return &VarInt{value}
}

func (x *VarInt) String () string {
    return "<VarInt:"+x.RawString()+">"
}

func (x *VarInt) RawString () string {
    return strconv.FormatUint(x.Data, 10)
}

func (x *VarInt) Encode () []byte {
    return VARINT.Wrap(x.RawEncode())
}

func (x *VarInt) RawEncode () []byte {
    return proto.EncodeVarint(x.Data)
}

func DecodeVarInt (data []byte) (res Any, err error) {
    x, n := proto.DecodeVarint(data)
    if n == 0 { return nil, errors.New("ran out of bytes while parsing varint") }
    return NewVarInt(x), nil
}

func (x *VarInt) Equal (other *VarInt) bool {
    return x.Data == other.Data
}

func (x *VarInt) EqualAny (other Any) bool {
    switch other := other.(type) {
        case *VarInt: return x.Equal(other)
        default: return false
    }
}