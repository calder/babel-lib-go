package babel

import "github.com/calder/fiddle"

/*****************
***   NilObj   ***
*****************/

var NIL = fiddle.FromHex("0000000000000000")

type NilObj struct {}

func (obj *NilObj) String () string {
    return "<Nil>"
}

func DecodeNil (obj *RawObj, recursive bool, dec *Decoder) Obj {
    return &NilObj{}
}

func (obj *NilObj) Encode () *fiddle.Bits {
    return NIL
}

/****************
***   IdObj   ***
****************/

var ID = fiddle.FromHex("823f70579c7a29bf")

type IdObj struct {
    Dat *fiddle.Bits
}

func (obj *IdObj) String () string {
    return "<Id:"+obj.Dat.Hex()+">"
}

func DecodeId (obj *RawObj, recursive bool, dec *Decoder) Obj {
    return &IdObj{obj.Dat}
}

func (obj *IdObj) Encode () *fiddle.Bits {
    return ID.Plus(obj.Dat)
}

/*****************
***   MsgObj   ***
*****************/

var MSG = fiddle.FromHex("83b10ff1ecf79c0b")

type MsgObj struct {
    To  Obj
    Dat Obj
}

func (obj *MsgObj) String () string {
    return "<Msg:"+obj.To.String()+","+obj.Dat.String()+">"
}

func DecodeMsg (obj *RawObj, recursive bool, dec *Decoder) Obj {
    toLen := int(obj.Dat.To(8).Byte())
    to    := dec.decode(obj.Dat.FromTo(8,toLen), recursive)
    dat   := dec.decode(obj.Dat.From(toLen+8), recursive)
    return &MsgObj{to,dat}
}

func (obj *MsgObj) Encode () *fiddle.Bits {
    to    := obj.To.Encode()
    dat   := obj.Dat.Encode()
    toLen := to.Len()
    return MSG.Plus(fiddle.FromByte(byte(toLen))).Plus(to).Plus(dat)
}

/*********************
***   UnicodeObj   ***
*********************/

var UNICODE = fiddle.FromHex("85847aa769e16613")

type UnicodeObj struct {
    Dat string
}

func (obj *UnicodeObj) String () string {
    return "<Unicode:"+obj.Dat+">"
}

func DecodeUnicode (obj *RawObj, recursive bool, dec *Decoder) Obj {
    return &UnicodeObj{obj.Dat.Unicode()}
}

func (obj *UnicodeObj) Encode () *fiddle.Bits {
    return UNICODE.Plus(fiddle.FromUnicode(obj.Dat))
}

/*********************
***   UdpAddrObj   ***
*********************/

var UDPADDR = fiddle.FromHex("8027db830a702671")

type UdpAddrObj struct {
    Dat string
}

func (obj *UdpAddrObj) String () string {
    return "<UdpAddr:"+obj.Dat+">"
}

func DecodeUdpAddr (obj *RawObj, recursive bool, dec *Decoder) Obj {
    return &UdpAddrObj{obj.Dat.Unicode()}
}

func (obj *UdpAddrObj) Encode () *fiddle.Bits {
    return UDPADDR.Plus(fiddle.FromUnicode(obj.Dat))
}