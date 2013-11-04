package babel

import "github.com/calder/fiddle"

/***********************************
***   Vita Over Local Dispatch   ***
***********************************/

func Receive (id *Id1, handler PacketHandler, errorHandler ErrorHandler) {
    addr, e := ReceiveUdp(":0", UdpMaxPacketSize, handler, errorHandler)
    if e != nil { panic(e) }
    e = SendUdp(UdpDispatchAddr, &UdpSub{id, &UdpAddrStr{addr}})
    if e != nil { panic(e) }
}

func Send (to *Id1, dat *fiddle.Bits) error {
    return SendUdp(UdpDispatchAddr, &Message{to, dat})
}