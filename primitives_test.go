package babel

import "crypto/rand"
import "crypto/rsa"
import "testing"

func TestBox (T *testing.T) {
    for i := 0; i < 1000; i++ {
        key := randId1()
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
    pub2   := DecodePubKey1(en.To(64), en.From(64)).(*PubKey1)
    if pub2.Key.N.Cmp(pub.N) != 0 || pub2.Key.E != pub.E {
        T.Log("Key:    ", pub)
        T.Log("Encoded:", en)
        T.Log("Decoded:", pub2)
        T.FailNow()
    }
}

func TestPriKey1 (T *testing.T) {
    pri, _ := rsa.GenerateKey(rand.Reader, 1024)
    en     := EncodePriKey1(&PriKey1{pri})
    pri2   := DecodePriKey1(en.To(64), en.From(64)).(*PriKey1).Key
    mismatch := false
    for i, p := range pri2.Primes {
        if p.Cmp(pri.Primes[i]) != 0 { mismatch = true }
    }
    if mismatch || pri2.D.Cmp(pri.D) != 0 || pri2.PublicKey.N.Cmp(pri.PublicKey.N) != 0 || pri2.PublicKey.E != pri.PublicKey.E {
        T.Log("Key:    ", &PriKey1{pri})
        T.Log("Encoded:", en)
        T.Log("Decoded:", &PriKey1{pri2})
        T.FailNow()
    }
}