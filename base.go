package babel

import "encoding/hex"
import "errors"

type Any interface {
    Type () []byte
    StringType () string
}

func AddType (typ []byte, decoder Decoder) {
    decoders[hex.EncodeToString(typ)] = decoder
}

func Tag (s string) []byte {
    var res, e = hex.DecodeString(s)
    if (e != nil) { panic(e) }
    return res
}

func Join (a, b []byte) []byte {
    var res = make([]byte, len(a)+len(b))
    copy(res[:len(a)], a)
    copy(res[len(a):], b)
    return res
}

type Decoder func([]byte)(Any,error)

var decoders = make(map[string]Decoder)

func Decode (data []byte) (res Any, err error) {
    var typ, e = FirstRune(data)

    if e != nil { return nil, errors.New("tag error:" + e.Error()) }

    var dat = data[len(typ):]
    var decoder = decoders[hex.EncodeToString(typ)]
    if decoder == nil {
        return nil, errors.New("unknown type "+hex.EncodeToString(typ))
    }
    return decoder(dat)
}