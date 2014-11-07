package babel

import "errors"

type Any interface {
    Type () *Type
    StringType () string
}

func AddType (t *Type, decoder Decoder) {
    decoders[t.Hex()] = decoder
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
    var t, e = FirstType(data)

    if e != nil { return nil, errors.New("type error:" + e.Error()) }

    var dat = data[t.Len():]
    var decoder = decoders[t.RawString()]
    if decoder == nil {
        return nil, errors.New("unknown type "+t.RawString())
    }
    return decoder(dat)
}