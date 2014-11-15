package babel

import "crypto/rand"
import "crypto/rsa"
import "encoding/hex"
import "errors"
import "testing"

func TestPriKey1 (T *testing.T) {
    key, _  := rsa.GenerateKey(rand.Reader, 1024)
    pri1    := &PriKey1{key}
    encoded := pri1.Encode(RAW)
    pri2, e := DecodePriKey1(encoded)

    if e == nil && !pri2.Equal(pri1) {
        e = errors.New("decoded != original")
    }

    if e != nil {
        T.Log("Error:  ", e)
        T.Log("Key:    ", pri1)
        T.Log("Encoded:", hex.EncodeToString(encoded))
        T.Log("Decoded:", pri2)
        T.FailNow()
    }
}