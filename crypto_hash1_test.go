package babel

import "crypto/rand"
import "encoding/hex"
import "errors"
import "testing"

func randHash1 () *Hash1 {
    data := [28]byte{}
    rand.Read(data[:])
    return NewHash1(data[:])
}

func TestHash1Encoding (T *testing.T) {
    for i := 0; i < 100; i++ {
        id := randHash1()
        encoded := id.Encode(RAW)
        decoded, err := DecodeHash1(encoded)
        if err == nil && !id.EqualValue(decoded) {
            err = errors.New("decoded id != original")
        }
        if err != nil {
            T.Log("Error:   ", err)
            T.Log("Original:", id)
            T.Log("Encoded: ", hex.EncodeToString(encoded))
            T.Log("Decoded: ", decoded)
            T.FailNow()
        }
    }
}
