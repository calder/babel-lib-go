// A PriKey1 is a 2048 bit RSA private key.

package babel

import "crypto/aes"
import "crypto/cipher"
import "crypto/rand"
import "crypto/rsa"
import "crypto/sha1"
import "errors"
import "math/big"
import "strconv"

var PRIKEY1 = Type("A9BAEF32")
func (*PriKey1) Type () uint64 { return PRIKEY1 }
func (*PriKey1) TypeName () string { return "PriKey1" }
func init () { AddType(PRIKEY1, decodePriKey1) }

type PriKey1 struct {
    Key *rsa.PrivateKey
}

func NewPriKey1 (key *rsa.PrivateKey) *PriKey1 {
    return &PriKey1{key}
}

func RandPriKey1 () *PriKey1 {
    key, e := rsa.GenerateKey(rand.Reader, 2048)
    if e != nil { panic(e) }
    return &PriKey1{key}
}

func (key *PriKey1) String () string {
    s := "<PriKey1:"
    for _, p := range key.Key.Primes { s += p.String()+"," }
    s += key.Key.D.String() + ","
    s += strconv.Itoa(key.Key.PublicKey.E) + ">"
    return s
}

func (key *PriKey1) Encode (enc Encoding) []byte {
    numPrimes := EncodeVarint(uint64(len(key.Key.Primes)))

    primes := make([][]byte, len(key.Key.Primes))
    for i, p := range key.Key.Primes {
        primes[i] = NewBigInt(p).Encode(LEN)
    }

    d := NewBigInt(key.Key.D).Encode(LEN)
    e := EncodeVarint(uint64(key.Key.PublicKey.E))

    return Wrap(enc, PRIKEY1, Join(numPrimes, Join(primes...), d, e))
}

func decodePriKey1 (data []byte) (res Value, err error) { return DecodePriKey1(data) }
func DecodePriKey1 (data []byte) (res *PriKey1, err error) {
    x, n := ReadVarint(data); data = data[n:]
    if n == 0 { return nil, errors.New("error parsing number of primes") }
    numPrimes := int(x)

    primes := make([]*big.Int, numPrimes)
    N := big.NewInt(1)
    for i := 0; i < numPrimes; i++ {
        P, n, e := ReadBigInt(data); data = data[n:]
        if e != nil { return nil, e }
        primes[i] = P.Data
        N.Mul(N, P.Data)
    }

    D, n, e := ReadBigInt(data); data = data[n:]
    if e != nil { return nil, e }
    E, n := ReadVarint(data); data = data[n:]

    if len(data) != 0 { return nil, errors.New("leftover bytes after parsing PRIKEY1") }

    pubKey := rsa.PublicKey{N:N,E:int(E)}
    priKey := &rsa.PrivateKey{PublicKey:pubKey, D:D.Data, Primes:primes}
    return &PriKey1{priKey}, nil
}

func ReadPriKey1 (data []byte) (res *PriKey1, n int, err error) {
    l, ll := ReadVarint(data)
    if ll == 0 { return nil, 0, errors.New("ran out of bytes while parsing PRIKEY1 length") }
    end := ll + int(l)
    if end > len(data) { return nil, 0, errors.New("ran out of bytes while parsing PRIKEY1") }
    res, err = DecodePriKey1(data[ll:end])
    return res, end, err
}

func (key *PriKey1) Equal (other *PriKey1) bool {
    for i, p := range key.Key.Primes {
        if p.Cmp(other.Key.Primes[i]) != 0 { return false }
    }
    if key.Key.D.Cmp(other.Key.D) != 0 { return false }
    if key.Key.PublicKey.N.Cmp(other.Key.PublicKey.N) != 0 { return false }
    if key.Key.PublicKey.E != other.Key.PublicKey.E { return false }
    return true
}

func (key *PriKey1) Id1 () *Id1 {
    return key.Pub().Id1()
}

func (key *PriKey1) Pub () *PubKey1 {
    return &PubKey1{&key.Key.PublicKey}
}

func (key *PriKey1) Decrypt (data []byte) []byte {
    // Break up the message chunks
    cipherKey := data[:2048/8]
    cipherText := data[2048/8:]

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
    return plainText
}
