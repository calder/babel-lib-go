package babel

import "bytes"
import "encoding/hex"
import "errors"
import "math/rand"
import "testing"

func TestRandRune (T *testing.T) {
    for t := 0; t < 100; t++ {
        var rune = RandRune(1 + rand.Intn(10))
        var _, e = FirstRuneLen(rune)

        if e == nil && !IsRune(rune) {
            e = errors.New("invalid rune")
        }

        if !IsRune(rune) {
            T.Log("Error: ", e)
            T.Log("Output:", hex.EncodeToString(rune))
            T.FailNow()
        }
    }
}

func TestFirstRune (T *testing.T) {
    for t := 0; t < 100; t++ {
        var runes = make([][]byte, 1 + rand.Int31n(10))
        var concatenated = make([]byte, 0)
        for i := 0; i < len(runes); i++ {
            runes[i] = RandRune(1 + i % 10)
            concatenated = append(concatenated, runes[i]...)
        }

        for i := 0; i < len(runes); i++ {
            var rune, e = FirstRune(concatenated)

            if e == nil && !bytes.Equal(rune, runes[i]) {
                e = errors.New("decoded rune != original")
            }

            if e != nil {
                T.Log("Error:       ", e)
                T.Log("Original:    ", hex.EncodeToString(runes[i]))
                T.Log("Concatenated:", hex.EncodeToString(concatenated))
                T.Log("Decoded:     ", hex.EncodeToString(rune))
                T.FailNow()
            }

            concatenated = concatenated[len(rune):]
        }
    }
}