package babel

import "errors"
import "math/rand"
import "testing"

func TestRandRune (T *testing.T) {
    for t := 0; t < 100; t++ {
        rune := RandRune(1 + rand.Intn(10))
        _, e := FirstRuneLen(rune.Bytes())

        if e == nil && !IsRune(rune.Bytes()) {
            e = errors.New("invalid rune")
        }

        if !IsRune(rune.Bytes()) {
            T.Log("Error: ", e)
            T.Log("Output:", rune)
            T.FailNow()
        }
    }
}

func TestFirstRune (T *testing.T) {
    for t := 0; t < 100; t++ {
        runes := make([]*Rune, 1 + rand.Int31n(10))
        concatenated := make([]byte, 0)
        for i := 0; i < len(runes); i++ {
            runes[i] = RandRune(1 + i % 10)
            concatenated = append(concatenated, runes[i].Bytes()...)
        }

        for i := 0; i < len(runes); i++ {
            rune, e := FirstRune(concatenated)

            if e == nil && !rune.Equal(runes[i]) {
                e = errors.New("decoded rune != original")
            }

            if e != nil {
                T.Log("Error:       ", e)
                T.Log("Original:    ", runes[i])
                T.Log("Concatenated:", concatenated)
                T.Log("Decoded:     ", rune)
                T.FailNow()
            }

            concatenated = concatenated[rune.Len():]
        }
    }
}