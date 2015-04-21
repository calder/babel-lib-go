// A Hash1 is the SHA-224 hash of a value. Its CBR acts like a 256-bit pointer
// to that value (32 bits of type tag + 224 bits of data).

package babel

import "bytes"
import "crypto/sha256"
import "encoding/hex"
import "errors"
import "strconv"

var ID1 = Type("F3C8AF31")
func (*Hash1) Type () uint64 { return ID1 }
func (*Hash1) TypeName () string { return "Hash1" }
func init () { AddType(ID1, decodeHash1) }

type Hash1 struct {
    data [28]byte
}

func NewHash1 (data []byte) *Hash1 {
    if len(data) != 28 {
        panic(errors.New("invalid length for Hash1: "+strconv.Itoa(len(data))))
    }
    h := &Hash1{}
    copy(h.data[:], data)
    return h
}

func Hash1OfData (data []byte) *Hash1 {
    hash := sha256.Sum224(data)
    return NewHash1(hash[:])
}

func Hash1OfValue (value Value) *Hash1 {
    return Hash1OfData(value.CBR())
}

func (h* Hash1) Data () []byte {
    return h.data[:]
}

func (h *Hash1) String () string {
    return "<Hash1:"+hex.EncodeToString(h.Data())+">"
}

func (h *Hash1) CBR () []byte {
    return h.Encode(TYPE)
}

func (h *Hash1) Encode (enc Encoding) []byte {
    return Wrap(enc, ID1, h.Data())
}

func decodeHash1 (data []byte) (res Value, err error) { return DecodeHash1(data) }
func DecodeHash1 (data []byte) (res *Hash1, err error) {
    if len(data) != 28 {
        return nil, errors.New("invalid length for Hash1: "+strconv.Itoa(len(data)))
    }
    return NewHash1(data), nil
}

func (h *Hash1) Equal (other *Hash1) bool {
    return bytes.Equal(h.Data(), other.Data())
}

func (h *Hash1) EqualValue (other Value) bool {
    switch other := other.(type) {
        case *Hash1: return h.Equal(other)
        default: return false
    }
}
