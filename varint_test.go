package babel

import "errors"
import "math/rand"
import "testing"

func TestVarIntEncoding (T *testing.T) {
    for i := 0; i < 100; i++ {
        original := NewVarInt(uint64(rand.Int63()))
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