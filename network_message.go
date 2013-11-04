package babel

import "github.com/calder/fiddle"

/******************
***   Message   ***
******************/

var MESSAGE = fiddle.FromRawHex("83b10ff1ecf79c0b")
func init () { AddType(MESSAGE, EncodeMessage, DecodeMessage) }

type Message struct {
    To  *Id1
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
    return &Message{decode(c[0]).(*Id1), c[1]}
}