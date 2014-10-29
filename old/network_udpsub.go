package babel

import "github.com/calder/fiddle"

/*****************
***   UdpSub   ***
*****************/

var UDPSUB = fiddle.FromRawHex("D9EB4EACD263ECFD")
func init () { AddType(UDPSUB, EncodeUdpSub, DecodeUdpSub) }

type UdpSub struct {
    Id1   *Id1
    Addr *UdpAddrStr
}

func (sub *UdpSub) String () string {
    return "<UdpSub:"+sub.Id1.String()+","+sub.Addr.String()+">"
}

func EncodeUdpSub (val Any) *fiddle.Bits {
    sub := val.(*UdpSub)
    return UDPSUB.Plus(fiddle.FromChunks(encode(sub.Id1), encode(sub.Addr)))
}

func DecodeUdpSub (typ *fiddle.Bits, dat *fiddle.Bits) Any {
    c := dat.Chunks(2)
    return &UdpSub{decode(c[0]).(*Id1), decode(c[1]).(*UdpAddrStr)}
}