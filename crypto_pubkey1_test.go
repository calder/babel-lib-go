package babel

import "crypto/rand"
import "crypto/rsa"
import "testing"

func TestPubKey1 (T *testing.T) {
    pri, _     := rsa.GenerateKey(rand.Reader, 1024)
    pub        := &PubKey1{&pri.PublicKey}
    encoded    := pub.Encode(RAW)
    decoded, e := DecodePubKey1(encoded)

    if e == nil && decoded

    if pub2.Key.N.Cmp(pub.N) != 0 || pub2.Key.E != pub.E {
        T.Log("Key:    ", pub)
        T.Log("Encoded:", en)
        T.Log("Decoded:", pub2)
        T.FailNow()
    }
}