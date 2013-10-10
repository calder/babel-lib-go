package babel

import "errors"
import "github.com/calder/fiddle"

/*****************
***   Public   ***
*****************/

type DecoderFunc func(*RawBin,bool)Bin

func AddType (typ *fiddle.Bits, decoder DecoderFunc) {
    decoders[typ.Hex()] = decoder
}

func Decode (bits *fiddle.Bits) Bin {
    return decode(bits, true)
}

func DecodeShallow (bits *fiddle.Bits) Bin {
    return decode(bits, false)
}

func DecodeBytes (bytes []byte) Bin {
    return decode(fiddle.FromBytes(bytes), true)
}

func DecodeShallowBytes (bytes []byte) Bin {
    return decode(fiddle.FromBytes(bytes), false)
}

/******************
***   Private   ***
******************/

var decoders = make(map[string]DecoderFunc)

func decode (bits *fiddle.Bits, recursive bool) Bin {
    if bits.Len() < 64 { panic(errors.New("Decoding error: missing type signature")) }
    typ := bits.To(64).Hex()
    fun := decoders[typ]
    if fun == nil { panic(errors.New("Decoding error: unkown type "+typ)) }
    bin := DecodeRaw(bits)
    return decoders[typ](bin, recursive)
}