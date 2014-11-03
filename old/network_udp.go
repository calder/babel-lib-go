package babel

import "errors"
import "net"

func ReceiveUdp (addrStr string, maxPacketBytes int, handler PacketHandler, errorHandler ErrorHandler) (realAddr string, err error) {
    // Create UDP listener
    addr, e := net.ResolveUDPAddr("udp", addrStr)
    if e != nil { return "", e }
    conn, e := net.ListenUDP("udp", addr)
    if e != nil { return "", e }

    go func () {
        for {
            // Wait for a packet
            buf := make([]byte, maxPacketBytes, maxPacketBytes)
            n, e := conn.Read(buf)
            if e != nil { errorHandler(e); continue }
            if n == maxPacketBytes { errorHandler(errors.New("oversized UDP packet")); continue }

            // Process the packet
            pkt, e := DecodeBytes(buf[:n])
            if e != nil { errorHandler(e); continue }
            handler(pkt)
        }
    }()

    return conn.LocalAddr().String(), nil
}

func SendUdp (addrStr string, pkt Any) error {
    // Create UDP connection
    addr, e := net.ResolveUDPAddr("udp", addrStr)
    if e != nil { return e }
    conn, e := net.DialUDP("udp", nil, addr)
    if e != nil { return e }

    // Send the packet
    bits, e := Encode(pkt)
    if e != nil { return e }
    bytes := bits.Bytes()
    n, e := conn.Write(bytes)
    if e != nil { return e }
    if n < len(bytes) { return errors.New("incomplete send") }
    return nil
}