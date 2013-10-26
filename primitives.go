package babel

import "crypto/rsa"
import "errors"
import "math/big"
import "strconv"
import "github.com/calder/fiddle"

/******************************
***   Bits (don't encode)   ***
******************************/

func init () { AddType(fiddle.Nil(), EncodeBits, nil) }

func EncodeBits (val Any) *fiddle.Bits {
    return val.(*fiddle.Bits)
}

/**************
***   Nil   ***
**************/

var NIL = fiddle.FromRawHex("0000000000000000")
func init () { AddType(NIL, EncodeNil, DecodeNil) }

func EncodeNil (val Any) *fiddle.Bits {
    if val != nil { panic(errors.New("EncodeNil() called on non-nil")) }
    return NIL
}

func DecodeNil (typ *fiddle.Bits, dat *fiddle.Bits) Any {
    return nil
}

/*************
***   Id   ***
*************/

var ID = fiddle.FromRawHex("823f70579c7a29bf")
func init () { AddType(ID, EncodeId, DecodeId) }

type Id struct {
    Dat *fiddle.Bits
}

func (id *Id) String () string {
    return "<Id:"+id.Dat.RawHex()+">"
}

func (id *Id) Equal (id2 *Id) bool {
    return id.Dat.Equal(id2.Dat)
}

func EncodeId (val Any) *fiddle.Bits {
    id := val.(*Id)
    return ID.Plus(id.Dat)
}

func DecodeId (typ *fiddle.Bits, dat *fiddle.Bits) Any {
    return &Id{dat}
}

/******************
***   Message   ***
******************/

var MESSAGE = fiddle.FromRawHex("83b10ff1ecf79c0b")
func init () { AddType(MESSAGE, EncodeMessage, DecodeMessage) }

type Message struct {
    To  *Id
    Dat *fiddle.Bits
}

func (msg *Message) String () string {
    return "<Message:"+msg.To.String()+","+msg.Dat.RawHex()+">"
}

func EncodeMessage (val Any) *fiddle.Bits {
    msg := val.(*Message)
    return MESSAGE.Plus(fiddle.FromChunks(encode(msg.To), msg.Dat))
}

func DecodeMessage (typ *fiddle.Bits, dat *fiddle.Bits) Any {
    c := dat.Chunks(2)
    return &Message{decode(c[0]).(*Id), c[1]}
}

/******************
***   Unicode   ***
******************/

var UNICODE = fiddle.FromRawHex("85847aa769e16613")
func init () { AddType(UNICODE, EncodeUnicode, DecodeUnicode) }

func DecodeUnicode (typ *fiddle.Bits, dat *fiddle.Bits) Any {
    return dat.Unicode()
}

func EncodeUnicode (val Any) *fiddle.Bits {
    str := val.(string)
    return UNICODE.Plus(fiddle.FromUnicode(str))
}

/*********************
***   UdpAddrStr   ***
*********************/

var UDPADDRSTR = fiddle.FromRawHex("8027db830a702671")
func init () { AddType(UDPADDRSTR, EncodeUdpAddrStr, DecodeUdpAddrStr) }

type UdpAddrStr struct {
    Dat string
}

func (addr *UdpAddrStr) String () string {
    return "<UdpAddrStr:"+addr.Dat+">"
}

func EncodeUdpAddrStr (val Any) *fiddle.Bits {
    addr := val.(*UdpAddrStr)
    return UDPADDRSTR.Plus(fiddle.FromUnicode(addr.Dat))
}

func DecodeUdpAddrStr (typ *fiddle.Bits, dat *fiddle.Bits) Any {
    return &UdpAddrStr{dat.Unicode()}
}

/*****************
***   UdpSub   ***
*****************/

var UDPSUB = fiddle.FromRawHex("D9EB4EACD263ECFD")
func init () { AddType(UDPSUB, EncodeUdpSub, DecodeUdpSub) }

type UdpSub struct {
    Id   *Id
    Addr *UdpAddrStr
}

func (sub *UdpSub) String () string {
    return "<UdpSub:"+sub.Id.String()+","+sub.Addr.String()+">"
}

func EncodeUdpSub (val Any) *fiddle.Bits {
    sub := val.(*UdpSub)
    return UDPSUB.Plus(fiddle.FromChunks(encode(sub.Id), encode(sub.Addr)))
}

func DecodeUdpSub (typ *fiddle.Bits, dat *fiddle.Bits) Any {
    c := dat.Chunks(2)
    return &UdpSub{decode(c[0]).(*Id), decode(c[1]).(*UdpAddrStr)}
}

/**************
***   Box   ***
**************/

var BOX = fiddle.FromRawHex("5946F91D56354917")
func init () { AddType(BOX, EncodeBox, DecodeBox) }

type Box struct {
    Key *Id
    Dat *fiddle.Bits
}

func (dat *Box) String () string {
    return "<Box:"+dat.Key.String()+","+dat.Dat.String()+">"
}

func EncodeBox (val Any) *fiddle.Bits {
    box := val.(*Box)
    return BOX.Plus(fiddle.FromChunks(encode(box.Key), box.Dat))
}

func DecodeBox (typ *fiddle.Bits, dat *fiddle.Bits) Any {
    c := dat.Chunks(2)
    return &Box{decode(c[0]).(*Id), c[1]}
}

/******************
***   PubKey1   ***
******************/

// Asymmetric: RSA (any key size)
// Symmetric:  AES 256 CFB
// Padding:    OAEP SHA-1

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

/******************
***   PriKey1   ***
******************/

var PRIKEY1 = fiddle.FromRawHex("D4B1E1B24361AFAF")
func init () { AddType(PRIKEY1, EncodePriKey1, DecodePriKey1) }

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

func EncodePriKey1 (val Any) *fiddle.Bits {
    key := val.(*PriKey1)
    primes := make([]*fiddle.Bits, len(key.Key.Primes))
    for i, p := range key.Key.Primes { primes[i] = fiddle.FromBigInt(p) }
    ps := fiddle.FromList(primes)
    d  := fiddle.FromBigInt(key.Key.D)
    e  := fiddle.FromInt(key.Key.PublicKey.E)
    return PRIKEY1.Plus(fiddle.FromChunks(ps, d, e))
}

func DecodePriKey1 (typ *fiddle.Bits, dat *fiddle.Bits) Any {
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