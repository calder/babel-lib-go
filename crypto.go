package babel

import "io"
import "crypto/aes"
import "crypto/cipher"
import "crypto/rand"
import "crypto/rsa"
import "crypto/sha1"
import "github.com/calder/fiddle"

/********************
***   Key Types   ***
********************/

type PubKey interface {
    Encrypt (*fiddle.Bits) *fiddle.Bits
    Id () *Id
}

type PriKey interface {
    Decrypt (*fiddle.Bits) *fiddle.Bits
    Id () *Id
}

type SimKey interface {
    Encrypt (*fiddle.Bits) *fiddle.Bits
    Decrypt (*fiddle.Bits) *fiddle.Bits
    Id () *Id
}

/******************
***   PubKey1   ***
******************/

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

/******************
***   PriKey1   ***
******************/

func (key *PriKey1) Decrypt (dat *fiddle.Bits) *fiddle.Bits {
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