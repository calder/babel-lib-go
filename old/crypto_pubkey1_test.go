package babel

import "crypto/rand"
import "crypto/rsa"
import "testing"

func TestPubKey1 (T *testing.T) {
    pri, _ := rsa.GenerateKey(rand.Reader, 1024)
    pub    := &pri.PublicKey
    en     := EncodePubKey1(&PubKey1{pub})
    pub2   := DecodePubKey1(en.To(64), en.From(64)).(*PubKey1)
    if pub2.Key.N.Cmp(pub.N) != 0 || pub2.Key.E != pub.E {
        T.Log("Key:    ", pub)
        T.Log("Encoded:", en)
        T.Log("Decoded:", pub2)
        T.FailNow()
    }
}