package babel

import "crypto/rand"
import "encoding/hex"
import "errors"
import "testing"

func randId1 () *Id1 {
    data := [16]byte{}
    rand.Read(data[:])
    return &Id1{data}
}

func TestId1Encoding (T *testing.T) {
    for i := 0; i < 100; i++ {
        id := randId1()
        encoded := id.Encode(RAW)
        decoded, err := DecodeId1(encoded)
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
