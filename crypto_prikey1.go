// A PriKey1 is a 2048 bit RSA key.

package babel

// import "crypto/aes"
// import "crypto/cipher"
import "crypto/rand"
import "crypto/rsa"
// import "crypto/sha1"
import "errors"
import "math/big"
import "strconv"

var PRIKEY1_STRING = "A9BAEF32"
var PRIKEY1 = NewTypeFromHex(PRIKEY1_STRING)
func (*PriKey1) Type () *Type { return PRIKEY1 }
func (*PriKey1) StringType () string { return PRIKEY1_STRING }
func init () { AddType(PRIKEY1, decodePriKey1) }

type PriKey1 struct {
    Key *rsa.PrivateKey
}

func NewPriKey1 (key *rsa.PrivateKey) *PriKey1 {
    return &PriKey1{key}
}

func RandPriKey1 () *PriKey1 {
    key, e := rsa.GenerateKey(rand.Reader, 4096)
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
        pb := p.Bytes()
        primes[i] = Join(EncodeVarint(uint64(len(pb))), pb)
    }

    d := NewBigInt(key.Key.D).Encode(RAW)
    e := EncodeVarint(uint64(key.Key.PublicKey.E))

    return Wrap(enc, PRIKEY1, Join(numPrimes, Join(primes...), d, e))
}

func decodePriKey1 (data []byte) (res Any, err error) { return DecodePriKey1(data) }
func DecodePriKey1 (data []byte) (res *PriKey1, err error) {
    x, n := ReadVarint(data)
    if n == 0 { return nil, errors.New("error parsing number of primes") }
    numPrimes := int(x)
    data = data[n:]

    primes := make([]*big.Int, numPrimes)
    for i := 0; i < numPrimes; i++ {
        p, n, err := ReadBigInt(data); data = data[n:]
        if err != nil { return nil, err }
        primes[i] = p.Data
    }

    d, n, err := ReadBigInt(data); data = data[n:]
    if err != nil { return nil, errors.New("PLACEHOLDER") }
    e, n := ReadVarint(data); data = data[n:]

    if len(data) != 0 { return nil, errors.New("leftover bytes after parsing PRIKEY1") }

    pubKey := rsa.PublicKey{d.Data,int(e)}
    priKey := &rsa.PrivateKey{PublicKey:pubKey, D:d.Data, Primes:primes}
    return &PriKey1{priKey}, nil
}

// func (key *PriKey1) Id1 () *Id1 {
//     return key.Pub().Id1()
// }

// func (key *PriKey1) Pub () *PubKey1 {
//     return &PubKey1{&key.Key.PublicKey}
// }

// func (key *PriKey1) Decrypt (dat []byte) []byte {
//     // Break up the message chunks
//     c := dat.Chunks(2)
//     cipherKey := c[0].RawBytes()
//     cipherText := c[1].RawBytes()

//     // Decrypt session key
//     plainKey, e := rsa.DecryptOAEP(sha1.New(), rand.Reader, key.Key, cipherKey, nil)
//     if e != nil { panic(e) }

//     // Create the block cipher
//     block, e := aes.NewCipher(plainKey)
//     if e != nil { panic(e) }

//     // Read the 128-bit initialization vector
//     iv := cipherText[:16]
//     cipherText = cipherText[16:]

//     // Create the stream cipher
//     stream := cipher.NewCFBDecrypter(block, iv)

//     // Decrypt message
//     plainText := make([]byte, len(cipherText))
//     stream.XORKeyStream(plainText, cipherText)

//     // Decode the message
//     return fiddle.FromBytes(plainText)
// }