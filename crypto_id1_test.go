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
        key := randId1()
        encoded := key.Encode()
        T.Log(hex.EncodeToString(encoded))
        decoded, err := Decode(encoded)
        if err == nil && !key.EqualAny(decoded) {
            err = errors.New("decoded key != original")
        }
        if err != nil {
            T.Log("Error:  ", err)
            T.Log("Key:    ", key)
            T.Log("Encoded:", hex.EncodeToString(encoded))
            T.Log("Decoded:", decoded)
            T.FailNow()
        }
    }
}