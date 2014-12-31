package babel

import "bytes"
import "errors"

type Value interface {
    Type () *Type
    StringType () string
}

func AddType (t *Type, decoder Decoder) {
    decoders[t.Hex()] = decoder
}

type Decoder func([]byte)(Value,error)

var decoders = make(map[string]Decoder)

func Decode (data []byte) (res Value, err error) {
    var t, e = FirstType(data)

    if e != nil { return nil, errors.New("type error:" + e.Error()) }

    var dat = data[t.Len():]
    var decoder = decoders[t.RawString()]
    if decoder == nil {
        return nil, errors.New("unknown type "+t.RawString())
    }
    return decoder(dat)
}

func Join (args ...[]byte) []byte {
    return bytes.Join(args, []byte{})
}

type Encoding byte
const RAW  = Encoding(0)
const LEN  = Encoding(1 << 0)
const TYPE = Encoding(1 << 1)

func Wrap (enc Encoding, typ *Type, data []byte) []byte {
    if enc&TYPE>0 { data = Join(typ.data, data) }
    if enc&LEN>0 { data = Join(EncodeVarint(uint64(len(data))), data) }
    return data
}