package babel

import "crypto/rand"
import "crypto/rsa"
import "encoding/hex"
import "errors"
import "testing"

func TestPubKey1 (T *testing.T) {
    pri, _  := rsa.GenerateKey(rand.Reader, 1024)
    pub1    := &PubKey1{&pri.PublicKey}
    encoded := pub1.Encode(RAW)
    pub2, e := DecodePubKey1(encoded)

    if e == nil && !pub2.Equal(pub1) {
        e = errors.New("decoded != original")
    }

    if e != nil {
        T.Log("Error:  ", e)
        T.Log("Key:    ", pub1)
        T.Log("Encoded:", hex.EncodeToString(encoded))
        T.Log("Decoded:", pub2)
        T.FailNow()
    }
}