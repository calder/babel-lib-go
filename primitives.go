package babel

import "crypto/rsa"
import "errors"
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

// Asymmetric: RSA 4096
// Symmetric:  AES 256
// Padding:    OAEP SHA-1

var PUBKEY1 = fiddle.FromRawHex("A7F3D2EE90717395")
func init () { AddType(PUBKEY1, EncodePubKey1, DecodePubKey1) }

type PubKey1 struct {
    Key *rsa.PublicKey
}

func (dat *PubKey1) String () string {
    return "<PubKey1:"+dat.Key.N.String()+","+string(dat.Key.E)+">"
}

func EncodePubKey1 (val Any) *fiddle.Bits {
    key := val.(*PubKey1)
    n := fiddle.FromBigInt(key.Key.N).PadLeft(4096)
    e := fiddle.FromInt(key.Key.E).PadLeft(32)
    return PUBKEY1.Plus(fiddle.FromChunks(n, e))
}

func DecodePubKey1 (typ *fiddle.Bits, dat *fiddle.Bits) Any {
    c := dat.Chunks(2)
    return &PubKey1{&rsa.PublicKey{c[0].BigInt(), c[1].Int()}}
}