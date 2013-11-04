package babel

import "log"

/********************
***   Constants   ***
********************/

const UdpDispatchAddr  = ":8124"
const UdpMaxPacketSize = 1048576

/*******************
***   Handlers   ***
*******************/

type PacketHandler func(Any)
type ErrorHandler func(error)

func ErrorLogger (err error) {
    log.Println("Warning: ", err)
}