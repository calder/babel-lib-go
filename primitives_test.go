package babel

import "crypto/rand"
import "crypto/rsa"
import "testing"

func TestBox (T *testing.T) {
    for i := 0; i < 1000; i++ {
        key := randId()
        dat := randBits()
        en  := EncodeBox(&Box{key,dat})
        de  := DecodeBox(en.To(64), en.From(64)).(*Box)
        if !de.Key.Dat.Equal(key.Dat) || !de.Dat.Equal(dat) {
            T.Log("Key:    ", key)
            T.Log("Dat:    ", dat)
            T.Log("Encoded:", en)
            T.Log("Decoded:", de)
            T.FailNow()
        }
    }
}

func TestPubKey1 (T *testing.T) {
    pri, _ := rsa.GenerateKey(rand.Reader, 1024)
    pub    := &pri.PublicKey
    en     := EncodePubKey1(&PubKey1{pub})
    pub2   := DecodePubKey1(en.To(64), en.From(64)).(*PubKey1).Key
    if pub2.N.Cmp(pub.N) != 0 || pub2.E != pub.E {
        T.Log("Key:    ", pub)
        T.Log("Encoded:", en)
        T.Log("Decoded:", pub2)
        T.FailNow()
    }
}