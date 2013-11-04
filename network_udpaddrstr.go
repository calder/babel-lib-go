package babel

import "github.com/calder/fiddle"

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