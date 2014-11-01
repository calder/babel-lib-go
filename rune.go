package babel

import "crypto/rand"
import "errors"

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