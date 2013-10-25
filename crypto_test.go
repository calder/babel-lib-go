package babel

import "crypto/rand"
import "crypto/rsa"
import "testing"

func Test1 (T *testing.T) {
    pri, _ := rsa.GenerateKey(rand.Reader, 1024)
    pub    := &PubKey1{&pri.PublicKey}
    println(pub.Encrypt(randBits()))
}