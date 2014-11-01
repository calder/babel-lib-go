package babel

import "crypto/rand"
import "encoding/hex"
import "errors"

var tagSize = 4 // TODO: implement dynamic tag sizes

var decoders = make(map[string]DecoderFunc)

type Any interface{}
type DecoderFunc func([]byte)(Any,error)

func AddType (typ []byte, decoder DecoderFunc) {
    decoders[hex.EncodeToString(typ)] = decoder
}

func Tag (s string) []byte {
    var res, e = hex.DecodeString(s)
    if (e != nil) { panic(e) }
    return res
}

func Join (a, b []byte) []byte {
    var res = make([]byte, len(a)+len(b))
    copy(res[:len(a)], a)
    copy(res[len(a):], b)
    return res
}

func Decode (data []byte) (res Any, err error) {
    if len(data) < tagSize { return nil, errors.New("missing type signature") }

    var typ = data[:tagSize]
    var dat = data[tagSize:]
    var decoder = decoders[hex.EncodeToString(typ)]
    if decoder == nil {
        return nil, errors.New("unknown type "+hex.EncodeToString(typ))
    }
    return decoder(dat)
}

func FirstRuneLen (bytes []byte) (length int, err error) {
    if len(bytes) == 0 { return -1, errors.New("rune must be at least one byte long") }

    var runeLen = 0
    for runeLen < len(bytes) && bytes[runeLen] & byte(128) != 0 {
        runeLen++
    }
    runeLen++
    if runeLen > len(bytes) { return -1, errors.New("unexpected end of rune") }

    return runeLen, nil
}

func FirstRune (bytes []byte) (rune []byte, err error) {
    var runeLen, e = FirstRuneLen(bytes)
    if e != nil { return nil, e }
    return bytes[:runeLen], nil
}

func IsRune (bytes []byte) bool {
    var runeLen, e = FirstRuneLen(bytes)
    return e == nil && runeLen == len(bytes)
}

func RandRune (length int) []byte {
    var rune = make([]byte, length)
    rand.Read(rune)
    for i := 0; i < length-1; i++ { rune[i] |= 128 }
    rune[length-1] &= 127
    return rune
}