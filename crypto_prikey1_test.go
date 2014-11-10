package babel

import "crypto/rand"
import "crypto/rsa"
import "errors"
import "testing"

func TestPriKey1 (T *testing.T) {
    key1, _ := rsa.GenerateKey(rand.Reader, 1024)
    pri1    := &PriKey1{key1}
    encoded := pri1.Encode(TYPE)
    p2, e   := Decode(encoded)
    pri2    := p2.(*PriKey1)
    key2    := pri2.Key

    if e == nil {
        for i, p := range key2.Primes {
            if p.Cmp(key2.Primes[i]) != 0 {
                e = errors.New("decoded primes != original primes")
                break
            }
        }
    }

    if e == nil && key2.D.Cmp(key2.D) != 0 {
        e = errors.New("decoded D != original D")
    }

    if e == nil && key2.PublicKey.N.Cmp(key2.PublicKey.N) != 0 {
        e = errors.New("decoded N != original N")
    }

    if e == nil && key2.PublicKey.E != key2.PublicKey.E {
        e = errors.New("decoded public key != original public key")
    }

    if e != nil {
        T.Log("Error:  ", e)
        T.Log("Key:    ", pri1)
        T.Log("Encoded:", encoded)
        T.Log("Decoded:", pri2)
        T.FailNow()
    }
}