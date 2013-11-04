package babel

import "crypto/rand"
import "crypto/rsa"
import "testing"

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