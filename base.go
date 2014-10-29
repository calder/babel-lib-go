package babel

import "encoding/hex"
import "errors"

var tagSize = 4 // TODO: implement dynamic tag sizes

var decoders = make(map[string]DecoderFunc)

type Any interface{}
type DecoderFunc func([]byte)(Any,error)

func AddType (typ []byte, decoder DecoderFunc) {
    decoders[hex.EncodeToString(typ)] = decoder
}

func Tag (s string) []byte {
    res, e := hex.DecodeString(s)
    if (e != nil) { panic(e) }
    return res
}

func Join (a, b []byte) []byte {
    res := make([]byte, len(a)+len(b))
    copy(res[:len(a)], a)
    copy(res[len(a):], b)
    return res
}

func Decode (data []byte) (res Any, err error) {
    if len(data) < tagSize {
        return nil, errors.New("missing type signature")
    }

    typ := data[:tagSize]
    dat := data[tagSize:]
    decoder := decoders[hex.EncodeToString(typ)]
    if decoder == nil {
        return nil, errors.New("unknown type "+hex.EncodeToString(typ))
    }
    return decoder(dat)
}