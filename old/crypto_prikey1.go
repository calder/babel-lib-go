// A PriKey1 is a 2048 bit RSA key.

package babel

import "crypto/aes"
import "crypto/cipher"
import "crypto/rand"
import "crypto/rsa"
import "crypto/sha1"
import "math/big"
import "strconv"

var PRIKEY1_STRING = "CBB6A416"
var PRIKEY1 = Tag(PRIKEY1_STRING)
func (*PriKey1) Type () []byte { return PRIKEY1 }
func (*PriKey1) StringType () string { return PRIKEY1_STRING }
func init () { AddType(PRIKEY1, DecodePriKey1) }

type PriKey1 struct {
    Key *rsa.PrivateKey
}

func (key *PriKey1) String () string {
    s := "<PriKey1:"
    for _, p := range key.Key.Primes { s += p.String()+"," }
    s += key.Key.D.String() + ","
    s += strconv.Itoa(key.Key.PublicKey.E) + ">"
    return s
}

func (key *PriKey1) Encode () []byte {
    primes := make([][]byte, len(key.Key.Primes))
    for i, p := range key.Key.Primes { primes[i] = fiddle.FromBigInt(p) }
    ps := fiddle.FromList(primes)
    d  := fiddle.FromBigInt(key.Key.D)
    e  := fiddle.FromInt(key.Key.PublicKey.E)
    return PRIKEY1.Plus(fiddle.FromChunks(ps, d, e))
}

func DecodePriKey1 (dat []byte) Any {
    c := dat.Chunks(3)
    ps := c[0].List()
    primes := make([]*big.Int, len(ps))
    for i, p := range ps { primes[i] = p.BigInt() }
    n := big.NewInt(1)
    for _, p := range primes { n.Mul(n,p) }
    d := c[1].BigInt()
    e := c[2].Int()
    return &PriKey1{&rsa.PrivateKey{PublicKey:rsa.PublicKey{n,e}, D:d, Primes:primes}}
}

func NewPriKey1 () *PriKey1 {
    key, e := rsa.GenerateKey(rand.Reader, 4096)
    if e != nil { panic(e) }
    return &PriKey1{key}
}

func (key *PriKey1) Id1 () *Id1 {
    return key.Pub().Id1()
}

func (key *PriKey1) Pub () *PubKey1 {
    return &PubKey1{&key.Key.PublicKey}
}

func (key *PriKey1) Decrypt (dat []byte) []byte {
    // Break up the message chunks
    c := dat.Chunks(2)
    cipherKey := c[0].RawBytes()
    cipherText := c[1].RawBytes()

    // Decrypt session key
    plainKey, e := rsa.DecryptOAEP(sha1.New(), rand.Reader, key.Key, cipherKey, nil)
    if e != nil { panic(e) }

    // Create the block cipher
    block, e := aes.NewCipher(plainKey)
    if e != nil { panic(e) }

    // Read the 128-bit initialization vector
    iv := cipherText[:16]
    cipherText = cipherText[16:]

    // Create the stream cipher
    stream := cipher.NewCFBDecrypter(block, iv)

    // Decrypt message
    plainText := make([]byte, len(cipherText))
    stream.XORKeyStream(plainText, cipherText)

    // Decode the message
    return fiddle.FromBytes(plainText)
}