package babel

import "errors"
import "math/rand"
import "testing"

func TestRandType (T *testing.T) {
    for t := 0; t < 100; t++ {
        t := RandType(1 + rand.Intn(10))
        _, e := FirstTypeLen(t.Bytes())

        if e == nil && !IsType(t.Bytes()) {
            e = errors.New("invalid t")
        }

        if !IsType(t.Bytes()) {
            T.Log("Error: ", e)
            T.Log("Output:", t)
            T.FailNow()
        }
    }
}

func TestFirstType (T *testing.T) {
    for t := 0; t < 100; t++ {
        ts := make([]*Type, 1 + rand.Int31n(10))
        concatenated := make([]byte, 0)
        for i := 0; i < len(ts); i++ {
            ts[i] = RandType(1 + i % 10)
            concatenated = append(concatenated, ts[i].Bytes()...)
        }

        for i := 0; i < len(ts); i++ {
            t, e := FirstType(concatenated)

            if e == nil && !t.Equal(ts[i]) {
                e = errors.New("decoded t != original")
            }

            if e != nil {
                T.Log("Error:       ", e)
                T.Log("Original:    ", ts[i])
                T.Log("Concatenated:", concatenated)
                T.Log("Decoded:     ", t)
                T.FailNow()
            }

            concatenated = concatenated[t.Len():]
        }
    }
}

var t1 = RandType(1)
var t2 = RandType(2)
var t4 = RandType(4)
var t8 = RandType(8)

func BenchmarkNewType1 (B *testing.B) { for t := 0; t < B.N; t++ { NewType(t1.Bytes()) } }
func BenchmarkNewType2 (B *testing.B) { for t := 0; t < B.N; t++ { NewType(t2.Bytes()) } }
func BenchmarkNewType4 (B *testing.B) { for t := 0; t < B.N; t++ { NewType(t4.Bytes()) } }
func BenchmarkNewType8 (B *testing.B) { for t := 0; t < B.N; t++ { NewType(t8.Bytes()) } }

func BenchmarkNewTypeUnchecked1 (B *testing.B) { for t := 0; t < B.N; t++ { NewTypeUnchecked(t1.Bytes()) } }
func BenchmarkNewTypeUnchecked2 (B *testing.B) { for t := 0; t < B.N; t++ { NewTypeUnchecked(t2.Bytes()) } }
func BenchmarkNewTypeUnchecked4 (B *testing.B) { for t := 0; t < B.N; t++ { NewTypeUnchecked(t4.Bytes()) } }
func BenchmarkNewTypeUnchecked8 (B *testing.B) { for t := 0; t < B.N; t++ { NewTypeUnchecked(t8.Bytes()) } }

func BenchmarkRandType1 (B *testing.B) { for t := 0; t < B.N; t++ { RandType(1) } }
func BenchmarkRandType2 (B *testing.B) { for t := 0; t < B.N; t++ { RandType(2) } }
func BenchmarkRandType4 (B *testing.B) { for t := 0; t < B.N; t++ { RandType(4) } }
func BenchmarkRandType8 (B *testing.B) { for t := 0; t < B.N; t++ { RandType(8) } }