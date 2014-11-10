// A variable length type t. The first bit of each byte determines whether to
// read the next byte.

package babel

import "bytes"
import "crypto/rand"
import "encoding/hex"
import "errors"

type Type struct {
    data []byte
}

func NewType (bytes []byte) *Type {
    t, e := NewTypeWithError(bytes)
    if e != nil { panic(e) }
    return t
}

func NewTypeFromHex (s string) *Type {
    var bytes, e = hex.DecodeString(s)
    if (e != nil) { panic(e) }
    return NewType(bytes)
}

func NewTypeWithError (bytes []byte) (res *Type, err error) {
    for i := 0; i < len(bytes)-1; i++ {
        if bytes[i] & 128 == 0 {
            return nil, errors.New("invalid t: inner byte continuation bit is 0")
        }
    }
    if bytes[len(bytes)-1] & 128 != 0 {
        return nil, errors.New("invalid t: final byte continuation bit is 1")
    }
    return NewTypeUnchecked(bytes), nil
}

func NewTypeUnchecked (bytes []byte) *Type {
    return &Type{bytes}
}

func (t *Type) Bytes () []byte {
    return t.data
}

func (t *Type) String () string {
    return "<Type:"+t.RawString()+">"
}

func (t *Type) RawString () string {
    return t.Hex()
}

func (t *Type) Hex () string {
    return hex.EncodeToString(t.data)
}

func (t *Type) Equal (other *Type) bool {
    return bytes.Equal(t.data, other.data)
}

func (t *Type) Len () int {
    return len(t.data)
}

func FirstTypeLen (bytes []byte) (length int, err error) {
    if len(bytes) == 0 { return -1, errors.New("t must be at least one byte long") }

    tLen := 0
    for tLen < len(bytes) && bytes[tLen] & byte(128) != 0 {
        tLen++
    }
    tLen++
    if tLen > len(bytes) { return -1, errors.New("unexpected end of t") }

    return tLen, nil
}

func FirstType (bytes []byte) (res *Type, err error) {
    tLen, e := FirstTypeLen(bytes)
    if e != nil { return nil, e }
    return NewTypeUnchecked(bytes[:tLen]), nil
}

func IsType (bytes []byte) bool {
    tLen, e := FirstTypeLen(bytes)
    return e == nil && tLen == len(bytes)
}

func RandType (length int) *Type {
    data := make([]byte, length)
    rand.Read(data)
    for i := 0; i < length-1; i++ { data[i] |= 128 }
    data[length-1] &= 127
    return NewTypeUnchecked(data)
}