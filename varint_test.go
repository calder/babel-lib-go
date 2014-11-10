package babel

import "encoding/hex"
import "errors"
import "math/rand"
import "testing"

func TestVarIntEncoding (T *testing.T) {
    for i := 0; i < 100; i++ {
        original := uint64(rand.Int63())
        encoded := EncodeVarint(original)
        decoded, n := ReadVarint(encoded)
        var e error

        if e == nil && n < len(encoded) {
            e = errors.New("bytes read < encoded bytes")
        }

        if e == nil && original != decoded {
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