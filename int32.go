// An Int32 is a signed, big endian, 32 bit integer.

package babel

import "bytes"
import "encoding/binary"
import "errors"
import "strconv"

var INT32_STRING = "B5D7B812"
var INT32 = Tag(INT32_STRING)
func (*Int32) Type () []byte { return INT32 }
func (*Int32) StringType () string { return INT32_STRING }
func init () { AddType(INT32, DecodeInt32) }

type Int32 struct {
    data int
}

func NewInt32 (value int) *Int32 {
    return &Int32{value}
}

func (x* Int32) Data () int {
    return x.data
}

func (x *Int32) String () string {
    return "<Int32:"+strconv.Itoa(x.data)+">"
}

func (x *Int32) Encode () []byte {
    buf := new(bytes.Buffer)
    binary.Write(buf, binary.BigEndian, x.data)
    return Join(INT32, buf.Bytes())
}

func DecodeInt32 (data []byte) (res Any, err error) {
    if len(data) != 4 {
        return nil, errors.New("invalid length for Int32: "+strconv.Itoa(len(data)))
    }

    x := new(Int32)
    e := binary.Read(bytes.NewReader(data), binary.BigEndian, &x.data)
    if e != nil { return nil, e }
    return x, nil
}

func (x *Int32) Equal (other *Int32) bool {
    return x.data == other.data
}

func (x *Int32) EqualAny (other Any) bool {
    switch other := other.(type) {
        case *Int32: return x.Equal(other)
        default: return false
    }
}