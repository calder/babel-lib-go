package babel

import "io"
import "crypto/aes"
import "crypto/cipher"
import "crypto/rand"
import "crypto/rsa"
import "crypto/sha1"
import "crypto/sha256"
import "strconv"
import "github.com/calder/fiddle"

/******************
***   PubKey1   ***
******************/

// Asymmetric: RSA (any key size)
// Symmetric:  AES 256 CFB
// Padding:    OAEP SHA-1
// Id1 Hash:   SHA 256

var PUBKEY1 = fiddle.FromRawHex("A7F3D2EE90717395")
func init () { AddType(PUBKEY1, EncodePubKey1, DecodePubKey1) }

type PubKey1 struct {
    Key *rsa.PublicKey
}

func (key *PubKey1) String () string {
    return "<PubKey1:"+key.Key.N.String()+","+strconv.Itoa(key.Key.E)+">"
}

func EncodePubKey1 (val Any) *fiddle.Bits {
    key := val.(*PubKey1)
    n := fiddle.FromBigInt(key.Key.N)
    e := fiddle.FromInt(key.Key.E)
    return PUBKEY1.Plus(fiddle.FromChunks(n, e))
}

func DecodePubKey1 (typ *fiddle.Bits, dat *fiddle.Bits) Any {
    c := dat.Chunks(2)
    return &PubKey1{&rsa.PublicKey{c[0].BigInt(), c[1].Int()}}
}

func (key *PubKey1) Id1 () *Id1 {
    hash := sha256.New()
    hash.Write(encode(key).Bytes())
    return &Id1{fiddle.FromBytes(hash.Sum([]byte{}))}
}

func (key *PubKey1) Encrypt (dat *fiddle.Bits) *fiddle.Bits {
    // Generate 256-bit session key
    plainKey := make([]byte, 256/8)
    _, e := io.ReadFull(rand.Reader, plainKey)
    if e != nil { panic(e) }

    // Encrypt the session key
    cipherKey, e := rsa.EncryptOAEP(sha1.New(), rand.Reader, key.Key, plainKey, nil)
    if e != nil { panic(e) }

    // Create the block cipher
    block, e := aes.NewCipher(plainKey)
    if e != nil { panic(e) }

    // Generate 128-bit initialization vector
    iv := make([]byte, aes.BlockSize)
    _, e = io.ReadFull(rand.Reader, iv)
    if e != nil { panic(e) }

    // Create the stream cipher
    stream := cipher.NewCFBEncrypter(block, iv)

    // Encrypt the message
    plainText := dat.Bytes()
    cipherText := make([]byte, len(plainText))
    stream.XORKeyStream(cipherText, plainText)

    // Prepend the initialization vector
    cipherText = append(iv, cipherText...)

    // Encode the message
    return fiddle.FromChunks(fiddle.FromRawBytes(cipherKey), fiddle.FromRawBytes(cipherText))
}