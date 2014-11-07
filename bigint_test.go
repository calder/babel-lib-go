package babel

import "encoding/hex"
import "errors"
import "math/big"
import "math/rand"
import "testing"

func TestBigIntEncoding (T *testing.T) {
    for i := 0; i < 100; i++ {
        original := NewBigInt(big.NewInt(int64(rand.Int63())))
        original.Data.Mul(original.Data, big.NewInt(int64(rand.Int63())))
        encoded := original.Encode()
        decoded, e := Decode(encoded)

        if e == nil && !original.EqualAny(decoded) {
            e = errors.New("decoded != original")
        }

        if e != nil {
            T.Log("Error:   ", e)
            T.Log("Original:", original)
            T.Log("Encoded: ", hex.EncodeToString(encoded))
            T.Log("Decoded: ", decoded)
            T.FailNow()
        }
    }
}