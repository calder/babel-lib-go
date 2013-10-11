package babel

import "errors"
import "github.com/calder/fiddle"

/**************
***   Nil   ***
**************/

var NIL = fiddle.FromRawHex("0000000000000000")
func init () { AddType(NIL, EncodeNil, DecodeNil) }

func DecodeNil (typ *fiddle.Bits, dat *fiddle.Bits) Any {
    return nil
}

func EncodeNil (val Any) *fiddle.Bits {
    if val != nil { panic(errors.New("EncodeNil() called on non-nil")) }
    return NIL
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

func DecodeId (typ *fiddle.Bits, dat *fiddle.Bits) Any {
    return &Id{dat}
}

func EncodeId (val Any) *fiddle.Bits {
    id := val.(*Id)
    return ID.Plus(id.Dat)
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
    return "<Message:"+msg.To.String()+","+msg.Dat.String()+">"
}

func DecodeMessage (typ *fiddle.Bits, dat *fiddle.Bits) Any {
    c := dat.Chunks(2)
    return &Message{decode(c[0]).(*Id), c[1]}
}

func EncodeMessage (val Any) *fiddle.Bits {
    msg := val.(*Message)
    return MESSAGE.Plus(fiddle.FromChunks(encode(msg.To), encode(msg.Dat)))
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

func DecodeUdpAddrStr (typ *fiddle.Bits, dat *fiddle.Bits) Any {
    return &UdpAddrStr{dat.Unicode()}
}

func EncodeUdpAddrStr (val Any) *fiddle.Bits {
    addr := val.(*UdpAddrStr)
    return UDPADDRSTR.Plus(fiddle.FromUnicode(addr.Dat))
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

func DecodeUdpSub (typ *fiddle.Bits, dat *fiddle.Bits) Any {
    c := dat.Chunks(2)
    return &UdpSub{decode(c[0]).(*Id), decode(c[1]).(*UdpAddrStr)}
}

func EncodeUdpSub (val Any) *fiddle.Bits {
    sub := val.(*UdpSub)
    return UDPSUB.Plus(fiddle.FromChunks(encode(sub.Id), encode(sub.Addr)))
}