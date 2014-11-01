package babel

// import "crypto/rand"
// import "crypto/rsa"
// import "testing"

// func TestAssymetricKey1 (T *testing.T) {
//     key, _ := rsa.GenerateKey(rand.Reader, 1024)
//     pri    := &PriKey1{key}
//     pub    := pri.Pub()

//     for i := 0; i < 100; i++ {
//         plain  := randBits()
//         cypher := pub.Encrypt(plain)
//         plain2 := pri.Decrypt(cypher)

//         if !plain2.Equal(plain) {
//             T.Log("Error:    ", "decypted message != original")
//             T.Log("Plaintext:", plain.String())
//             T.Log("Encrypted:", cypher.String())
//             T.Log("Decrypted:", plain2.String())
//             T.FailNow()
//         }
//     }
// }