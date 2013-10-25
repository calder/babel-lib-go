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
    pri, _ := rsa.GenerateKey(rand.Reader, 4096)
    key    := &pri.PublicKey
    en     := EncodePubKey1(&PubKey1{key})
    key2   := DecodePubKey1(en.To(64), en.From(64)).(*PubKey1).Key
    if key2.N.Cmp(key.N) != 0 || key2.E != key.E {
        T.Log("Key:    ", key)
        T.Log("Encoded:", en)
        T.Log("Decoded:", key2)
        T.FailNow()
    }
}