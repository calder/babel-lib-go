package babel

import "errors"
import "github.com/calder/fiddle"

/*****************
***   Public   ***
*****************/

type Any interface{}
type Bits *fiddle.Bits
type EncoderFunc func(Any)*fiddle.Bits
type DecoderFunc func(*fiddle.Bits,*fiddle.Bits)Any

func AddType (typ *fiddle.Bits, encoder EncoderFunc, decoder DecoderFunc) {
    encoders = append(encoders, encoder)
    decoders[typ.RawHex()] = decoder
}

func Encode (val Any) (bits *fiddle.Bits, err error) {
    defer func () {
        e := recover()
        if err != nil { err = e.(error) }
    }()
    return encode(val), err
}

func Decode (bits *fiddle.Bits) (res Any, err error) {
    defer func () {
        e := recover()
        if err != nil { err = e.(error) }
    }()
    return decode(bits), err
}

func DecodeBytes (bytes []byte) (res Any, err error) {
    return Decode(fiddle.FromBytes(bytes))
}

func EncodeUnsafe (val Any) *fiddle.Bits {
    return encode(val)
}

func DecodeUnsafe (bits *fiddle.Bits) Any {
    return decode(bits)
}

/******************
***   Private   ***
******************/

var encoders = make([]EncoderFunc, 0)
var decoders = make(map[string]DecoderFunc)

func encode (val Any) *fiddle.Bits {
    var res *fiddle.Bits
    err := errors.New("unable to decode")
    for _, encoder := range encoders {
        func () (succ bool) {
            defer func () {
                e := recover()
                if e != nil { succ = false }
            }()
            res = encoder(val)
            err = nil
            return succ
        }()
    }
    if err != nil { panic(err) }
    return res
}

func decode (bits *fiddle.Bits) Any {
    if bits.Len() < 64 { panic(errors.New("decoding error: missing type signature")) }
    typStr := bits.To(64).RawHex()
    decoder := decoders[typStr]
    if decoder == nil { panic(errors.New("decoding error: unknown type "+typStr)) }
    typ := bits.To(64)
    dat := bits.From(64)
    return decoder(typ, dat)
}