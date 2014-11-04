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

var rune1 = RandRune(1)
var rune2 = RandRune(2)
var rune4 = RandRune(4)
var rune8 = RandRune(8)

func BenchmarkNewRune1 (B *testing.B) { for t := 0; t < B.N; t++ { NewRune(rune1.Bytes()) } }
func BenchmarkNewRune2 (B *testing.B) { for t := 0; t < B.N; t++ { NewRune(rune2.Bytes()) } }
func BenchmarkNewRune4 (B *testing.B) { for t := 0; t < B.N; t++ { NewRune(rune4.Bytes()) } }
func BenchmarkNewRune8 (B *testing.B) { for t := 0; t < B.N; t++ { NewRune(rune8.Bytes()) } }

func BenchmarkNewRuneUnchecked1 (B *testing.B) { for t := 0; t < B.N; t++ { NewRuneUnchecked(rune1.Bytes()) } }
func BenchmarkNewRuneUnchecked2 (B *testing.B) { for t := 0; t < B.N; t++ { NewRuneUnchecked(rune2.Bytes()) } }
func BenchmarkNewRuneUnchecked4 (B *testing.B) { for t := 0; t < B.N; t++ { NewRuneUnchecked(rune4.Bytes()) } }
func BenchmarkNewRuneUnchecked8 (B *testing.B) { for t := 0; t < B.N; t++ { NewRuneUnchecked(rune8.Bytes()) } }

func BenchmarkRandRune1 (B *testing.B) { for t := 0; t < B.N; t++ { RandRune(1) } }
func BenchmarkRandRune2 (B *testing.B) { for t := 0; t < B.N; t++ { RandRune(2) } }
func BenchmarkRandRune4 (B *testing.B) { for t := 0; t < B.N; t++ { RandRune(4) } }
func BenchmarkRandRune8 (B *testing.B) { for t := 0; t < B.N; t++ { RandRune(8) } }