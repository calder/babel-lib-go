// An Int32 is a signed, big endian, 32 bit int32eger.

package babel

import "bytes"
import "encoding/binary"
import "errors"
import "strconv"

var INT32 = Type("B5D7B812")
func (*Int32) Type () uint64 { return INT32 }
func (*Int32) TypeName () string { return "Int32" }
func init () { AddType(INT32, DecodeInt32) }

type Int32 struct {
    Data int32
}

func NewInt32 (value int32) *Int32 {
    return &Int32{value}
}

func (x *Int32) String () string {
    return "<Int32:"+strconv.Itoa(int(x.Data))+">"
}

func (x *Int32) Encode (enc Encoding) []byte {
    buf := new(bytes.Buffer)
    binary.Write(buf, binary.BigEndian, x.Data)
    return Wrap(enc, INT32, buf.Bytes())
}

func DecodeInt32 (data []byte) (res Value, err error) {
    x, n, e := ReadInt32(data)
    if e != nil { return nil, e }
    if n < len(data) { return nil, errors.New("leftover bytes after parsing int32") }
    return x, nil
}

func ReadInt32 (data []byte) (res *Int32, n int, err error) {
    x := new(Int32)
    e := binary.Read(bytes.NewReader(data[:4]), binary.BigEndian, &x.Data)
    if e != nil { return nil, 0, e }
    return x, 4, nil
}

func (x *Int32) Equal (other *Int32) bool {
    return x.Data == other.Data
}

func (x *Int32) EqualValue (other Value) bool {
    switch other := other.(type) {
        case *Int32: return x.Equal(other)
        default: return false
    }
}
