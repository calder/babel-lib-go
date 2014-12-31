// A PubKey1 is a 2048 bit RSA public key.
//
// Encryption specification:
//     Symmetric:  AES 256 CFB
//     Padding:    OAEP SHA-1

package babel

import "io"
import "crypto/aes"
import "crypto/cipher"
import "crypto/rand"
import "crypto/rsa"
import "crypto/sha1"
import "crypto/sha256"
import "errors"
import "strconv"

var PUBKEY1_STRING = "D1E8A30F"
var PUBKEY1 = NewTypeFromHex(PUBKEY1_STRING)
func (*PubKey1) Type () *Type { return PUBKEY1 }
func (*PubKey1) StringType () string { return PUBKEY1_STRING }
func init () { AddType(PUBKEY1, decodePubKey1) }

type PubKey1 struct {
    Key *rsa.PublicKey
}

func (key *PubKey1) String () string {
    return "<PubKey1:"+key.Key.N.String()+","+strconv.Itoa(key.Key.E)+">"
}

func (key *PubKey1) Encode (enc Encoding) []byte {
    E := EncodeVarint(uint64(key.Key.E))
    N := NewBigInt(key.Key.N).Encode(RAW)
    return Wrap(enc, PUBKEY1, Join(E, N))
}

func decodePubKey1 (data []byte) (res Value, err error) { return DecodePubKey1(data) }
func DecodePubKey1 (data []byte) (res *PubKey1, err error) {
    E, n := ReadVarint(data); data = data[n:]
    if n == 0 { return nil, errors.New("ran out of bytes while parsing PubKey1.E") }
    N, e := DecodeBigInt(data)
    if e != nil { return nil, e }
    return &PubKey1{&rsa.PublicKey{N:N.Data, E:int(E)}}, nil
}

func ReadPubKey1 (data []byte) (res *PubKey1, n int, err error) {
    l, ll := ReadVarint(data)
    if ll == 0 { return nil, 0, errors.New("ran out of bytes while parsing PUBKEY1 length") }
    end := ll + int(l)
    if end > len(data) { return nil, 0, errors.New("ran out of bytes while parsing PUBKEY1") }
    res, err = DecodePubKey1(data[ll:end])
    return res, end, err
}

func (key *PubKey1) Equal (other *PubKey1) bool {
    if key.Key.N.Cmp(other.Key.N) != 0 { return false }
    if key.Key.E != other.Key.E { return false }
    return true
}

func (key *PubKey1) Id1 () *Id1 {
    hash := sha256.New()
    hash.Write(key.Encode(TYPE))
    return NewId1(hash.Sum([]byte{})[:16])
}

func (key *PubKey1) Encrypt (plainText []byte) []byte {
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
    cipherText := make([]byte, len(plainText))
    stream.XORKeyStream(cipherText, plainText)

    // Prepend the initialization vector
    cipherText = append(iv, cipherText...)

    // Encode the message
    println(len(cipherKey))
    return Join(cipherKey, cipherText)
}