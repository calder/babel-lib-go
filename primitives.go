package babel

import "github.com/calder/fiddle"

/*****************
***   NilBin   ***
*****************/

var NIL = fiddle.FromHex("0000000000000000")

type NilBin struct {}

func (bin *NilBin) String () string {
    return "<Nil>"
}

func DecodeNil (bin *RawBin, recursive bool, dec *Decoder) Bin {
    return &NilBin{}
}

func (bin *NilBin) Encode () *fiddle.Bits {
    return NIL
}

/****************
***   IdBin   ***
****************/

var ID = fiddle.FromHex("823f70579c7a29bf")

type IdBin struct {
    Dat *fiddle.Bits
}

func (bin *IdBin) String () string {
    return "<Id:"+bin.Dat.Hex()+">"
}

func DecodeId (bin *RawBin, recursive bool, dec *Decoder) Bin {
    return &IdBin{bin.Dat}
}

func (bin *IdBin) Encode () *fiddle.Bits {
    return ID.Plus(bin.Dat)
}

/*****************
***   MsgBin   ***
*****************/

var MSG = fiddle.FromHex("83b10ff1ecf79c0b")

type MsgBin struct {
    To  Bin
    Dat Bin
}

func (bin *MsgBin) String () string {
    return "<Msg:"+bin.To.String()+","+bin.Dat.String()+">"
}

func DecodeMsg (bin *RawBin, recursive bool, dec *Decoder) Bin {
    c := bin.Dat.Chunks(2)
    to  := dec.decode(c[0], recursive)
    dat := dec.decode(c[1], recursive)
    return &MsgBin{to,dat}
}

func (bin *MsgBin) Encode () *fiddle.Bits {
    return MSG.Plus(fiddle.FromChunks(bin.To.Encode(), bin.Dat.Encode()))
}

/*********************
***   UnicodeBin   ***
*********************/

var UNICODE = fiddle.FromHex("85847aa769e16613")

type UnicodeBin struct {
    Dat string
}

func (bin *UnicodeBin) String () string {
    return "<Unicode:"+bin.Dat+">"
}

func DecodeUnicode (bin *RawBin, recursive bool, dec *Decoder) Bin {
    return &UnicodeBin{bin.Dat.Unicode()}
}

func (bin *UnicodeBin) Encode () *fiddle.Bits {
    return UNICODE.Plus(fiddle.FromUnicode(bin.Dat))
}

/************************
***   UdpAddrStrBin   ***
************************/

var UDPADDRSTR = fiddle.FromHex("8027db830a702671")

type UdpAddrStrBin struct {
    Dat string
}

func (bin *UdpAddrStrBin) String () string {
    return "<UdpAddrStr:"+bin.Dat+">"
}

func DecodeUdpAddrStr (bin *RawBin, recursive bool, dec *Decoder) Bin {
    return &UdpAddrStrBin{bin.Dat.Unicode()}
}

func (bin *UdpAddrStrBin) Encode () *fiddle.Bits {
    return UDPADDRSTR.Plus(fiddle.FromUnicode(bin.Dat))
}

/********************
***   UdpSubBin   ***
********************/

var UDPSUB = fiddle.FromHex("D9EB4EACD263ECFD")

type UdpSubBin struct {
    Id   Bin
    Addr Bin
}

func (bin *UdpSubBin) String () string {
    return "<UdpSub:"+bin.Id.String()+","+bin.Addr.String()+">"
}

func DecodeUdpSub (bin *RawBin, recursive bool, dec *Decoder) Bin {
    c := bin.Dat.Chunks(2)
    id   := dec.decode(c[0], recursive)
    addr := dec.decode(c[1], recursive)
    return &UdpSubBin{id,addr}
}

func (bin *UdpSubBin) Encode () *fiddle.Bits {
    return UDPSUB.Plus(fiddle.FromChunks(bin.Id.Encode(), bin.Addr.Encode()))
}