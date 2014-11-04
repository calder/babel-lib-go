package babel

import "errors"
import "math"
import "math/big"
import "testing"

func TestVarIntEncoding (T *testing.T) {
    for i := 0; i < 100; i++ {
        x := big.NewInt(math.MaxInt64)
        y := big.NewInt(math.MaxInt64)
        original := NewVarInt(0)
        original.Data.Mul(x,y)

        encoded := original.Encode()
        decoded, e := Decode(encoded)

        if e == nil && !original.EqualAny(decoded) {
            e = errors.New("decoded != original")
        }

        if e != nil {
            T.Log("Error:   ", e)
            T.Log("Original:", original)
            T.Log("Decoded: ", decoded)
            T.FailNow()
        }
    }
}