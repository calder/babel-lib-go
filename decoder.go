package babel

import "errors"
import "github.com/calder/fiddle"

/****************
***   Types   ***
****************/

type DecoderFunc func(*RawBin,bool,*Decoder)Bin

type Decoder struct {
    fns map[string]DecoderFunc
}

/***********************
***   Constructors   ***
***********************/

func NewDecoder () *Decoder {
    dec := NewEmptyDecoder()

    dec.AddType(NIL,        DecodeNil)
    dec.AddType(ID,         DecodeId)
    dec.AddType(MSG,        DecodeMsg)
    dec.AddType(UNICODE,    DecodeUnicode)
    dec.AddType(UDPADDRSTR, DecodeUdpAddrStr)
    dec.AddType(UDPSUB,     DecodeUdpSub)

    return dec
}

func NewEmptyDecoder () *Decoder {
    return &Decoder{make(map[string]DecoderFunc)}
}

/******************
***   Methods   ***
******************/

func (dec *Decoder) AddType (typ *fiddle.Bits, decoder DecoderFunc) {
    dec.fns[typ.Hex()] = decoder
}

func (dec *Decoder) decode (bits *fiddle.Bits, recursive bool) Bin {
    if bits.Len() < 64 { panic(errors.New("Decoding error: missing type signature")) }
    typ := bits.To(64).Hex()
    bin := DecodeRaw(bits)
    return dec.fns[typ](bin, recursive, dec)
}

func (dec *Decoder) Decode (bits *fiddle.Bits) Bin {
    return dec.decode(bits, true)
}

func (dec *Decoder) DecodePartial (bits *fiddle.Bits) Bin {
    return dec.decode(bits, false)
}

func (dec *Decoder) DecodeBytes (bytes []byte) Bin {
    return dec.decode(fiddle.FromBytes(bytes), true)
}

func (dec *Decoder) DecodePartialBytes (bytes []byte) Bin {
    return dec.decode(fiddle.FromBytes(bytes), false)
}