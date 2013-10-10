package babel

import "github.com/calder/fiddle"

/*****************
***   NilBin   ***
*****************/

var NIL = fiddle.FromRawHex("0000000000000000")
func init () { AddType(NIL, DecodeNil) }

type NilBin struct {}

func (bin *NilBin) String () string {
    return "<Nil>"
}

func DecodeNil (bin *RawBin, recursive bool) Bin {
    return &NilBin{}
}

func (bin *NilBin) Encode () *fiddle.Bits {
    return NIL
}

/****************
***   IdBin   ***
****************/

var ID = fiddle.FromRawHex("823f70579c7a29bf")
func init () { AddType(ID, DecodeId) }

type IdBin struct {
    Dat *fiddle.Bits
}

func (bin *IdBin) String () string {
    return "<Id:"+bin.Dat.RawHex()+">"
}

func DecodeId (bin *RawBin, recursive bool) Bin {
    return &IdBin{bin.Dat}
}

func (bin *IdBin) Encode () *fiddle.Bits {
    return ID.Plus(bin.Dat)
}

/*****************
***   MsgBin   ***
*****************/

var MSG = fiddle.FromRawHex("83b10ff1ecf79c0b")
func init () { AddType(MSG, DecodeMsg) }

type MsgBin struct {
    To  Bin
    Dat Bin
}

func (bin *MsgBin) String () string {
    return "<Msg:"+bin.To.String()+","+bin.Dat.String()+">"
}

func DecodeMsg (bin *RawBin, recursive bool) Bin {
    c := bin.Dat.Chunks(2)
    to  := decode(c[0], recursive)
    dat := decode(c[1], recursive)
    return &MsgBin{to,dat}
}

func (bin *MsgBin) Encode () *fiddle.Bits {
    return MSG.Plus(fiddle.FromChunks(bin.To.Encode(), bin.Dat.Encode()))
}

/*********************
***   UnicodeBin   ***
*********************/

var UNICODE = fiddle.FromRawHex("85847aa769e16613")
func init () { AddType(UNICODE, DecodeUnicode) }

type UnicodeBin struct {
    Dat string
}

func (bin *UnicodeBin) String () string {
    return "<Unicode:"+bin.Dat+">"
}

func DecodeUnicode (bin *RawBin, recursive bool) Bin {
    return &UnicodeBin{bin.Dat.Unicode()}
}

func (bin *UnicodeBin) Encode () *fiddle.Bits {
    return UNICODE.Plus(fiddle.FromUnicode(bin.Dat))
}

/************************
***   UdpAddrStrBin   ***
************************/

var UDPADDRSTR = fiddle.FromRawHex("8027db830a702671")
func init () { AddType(UDPADDRSTR, DecodeUdpAddrStr) }

type UdpAddrStrBin struct {
    Dat string
}

func (bin *UdpAddrStrBin) String () string {
    return "<UdpAddrStr:"+bin.Dat+">"
}

func DecodeUdpAddrStr (bin *RawBin, recursive bool) Bin {
    return &UdpAddrStrBin{bin.Dat.Unicode()}
}

func (bin *UdpAddrStrBin) Encode () *fiddle.Bits {
    return UDPADDRSTR.Plus(fiddle.FromUnicode(bin.Dat))
}

/********************
***   UdpSubBin   ***
********************/

var UDPSUB = fiddle.FromRawHex("D9EB4EACD263ECFD")
func init () { AddType(UDPSUB, DecodeUdpSub) }

type UdpSubBin struct {
    Id   Bin
    Addr Bin
}

func (bin *UdpSubBin) String () string {
    return "<UdpSub:"+bin.Id.String()+","+bin.Addr.String()+">"
}

func DecodeUdpSub (bin *RawBin, recursive bool) Bin {
    c := bin.Dat.Chunks(2)
    id   := decode(c[0], recursive)
    addr := decode(c[1], recursive)
    return &UdpSubBin{id,addr}
}

func (bin *UdpSubBin) Encode () *fiddle.Bits {
    return UDPSUB.Plus(fiddle.FromChunks(bin.Id.Encode(), bin.Addr.Encode()))
}