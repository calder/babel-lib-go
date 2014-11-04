package babel

import "bytes"
import "crypto/rand"
import "encoding/hex"
import "errors"

type Rune struct {
    data []byte
}

func NewRune (bytes []byte) *Rune {
    rune, e := NewRuneWithError(bytes)
    if e != nil { panic(e) }
    return rune
}

func NewRuneWithError (bytes []byte) (res *Rune, err error) {
    for i := 0; i < len(bytes)-1; i++ {
        if bytes[i] & 128 == 0 {
            return nil, errors.New("invalid rune: inner byte continuation bit is 0")
        }
    }
    if bytes[len(bytes)-1] & 128 != 0 {
        return nil, errors.New("invalid rune: final byte continuation bit is 1")
    }
    return NewRuneUnchecked(bytes), nil
}

func NewRuneUnchecked (bytes []byte) *Rune {
    return &Rune{bytes}
}

func (rune *Rune) Bytes () []byte {
    return rune.data
}

func (rune *Rune) String () string {
    return "<Rune:"+rune.RawString()+">"
}

func (rune *Rune) RawString () string {
    return hex.EncodeToString(rune.data)
}

func (rune *Rune) Equal (other *Rune) bool {
    return bytes.Equal(rune.data, other.data)
}

func (rune *Rune) Len () int {
    return len(rune.data)
}

func FirstRuneLen (bytes []byte) (length int, err error) {
    if len(bytes) == 0 { return -1, errors.New("rune must be at least one byte long") }

    runeLen := 0
    for runeLen < len(bytes) && bytes[runeLen] & byte(128) != 0 {
        runeLen++
    }
    runeLen++
    if runeLen > len(bytes) { return -1, errors.New("unexpected end of rune") }

    return runeLen, nil
}

func FirstRune (bytes []byte) (res *Rune, err error) {
    runeLen, e := FirstRuneLen(bytes)
    if e != nil { return nil, e }
    return NewRuneUnchecked(bytes[:runeLen]), nil
}

func IsRune (bytes []byte) bool {
    runeLen, e := FirstRuneLen(bytes)
    return e == nil && runeLen == len(bytes)
}

func RandRune (length int) *Rune {
    data := make([]byte, length)
    rand.Read(data)
    for i := 0; i < length-1; i++ { data[i] |= 128 }
    data[length-1] &= 127
    return NewRuneUnchecked(data)
}