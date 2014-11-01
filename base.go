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

// func FirstRune (bytes []byte) (rune []byte, err error) {
//     if len(bytes) == 0 { return nil, errors.New("expected non-empty slice") }

//     var byteLen = 0
//     for byteLen < len(bytes) && bytes[byteLen] & byte(128) != 0 {
//         byteLen++
//     }
//     byteLen++
//     if byteLen > len(bytes) { return nil, errors.New("rune length > slice length") }

//     var runeLen = (byteLen * 7 + 7) / 8
//     var offset = byteLen%8
//     rune = make([]byte, runeLen)

//     for i := 0; i < byteLen; i++ {
//         var start = offset + i * 7
//         var end = start + 6

//         // Copy most significant bits
//         // println("Copying byte", i, "to bits", start, "through", end)
//         rune[start/8] |= (bytes[i] & 127) >> uint(start % 8 - 1)
//         // println(hex.EncodeToString(rune))

//         // Copy least significant bits
//         if end/8 > start/8 {
//             // println("Copying second part", uint(7-end % 8))
//             rune[end/8] |= (bytes[i] & 127) << uint(7 - end % 8)
//             // println(hex.EncodeToString(rune))
//         }

//         // println()
//     }

//     return rune, nil
// }

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