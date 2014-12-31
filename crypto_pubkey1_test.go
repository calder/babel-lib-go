package babel

import "bytes"
import "crypto/rand"
import "encoding/hex"
import "errors"
import "testing"

func TestPubKey1Encoding (T *testing.T) {
    pri     := NewPriKey1(rsa2048PrivateKey)
    pub1    := pri.Pub()
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

func TestPubKey1Encryption (T *testing.T) {
    pri := NewPriKey1(rsa2048PrivateKey)
    pub := pri.Pub()

    message := make([]byte, 1024)
    rand.Read(message)

    encrypted := pub.Encrypt(message)
    decrypted := pri.Decrypt(encrypted)

    if !bytes.Equal(message, decrypted) {
        T.FailNow()
    }
}