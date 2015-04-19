package babel

import "bytes"
import "errors"

type Value interface {
    Type () uint64
    TypeName () string
}

func AddType (t uint64, decoder Decoder) {
    decoders[t] = decoder
}

type Decoder func([]byte)(Value,error)

var decoders = make(map[uint64]Decoder)

// func Decode (data []byte) (res Value, err error) {
//     var t, e = FirstType(data)

//     if e != nil { return nil, errors.New("type error:" + e.Error()) }

//     var dat = data[t.Len():]
//     var decoder = decoders[t.RawString()]
//     if decoder == nil {
//         return nil, errors.New("unknown type "+t.RawString())
//     }
//     return decoder(dat)
// }

// Concatenate byte arrays.
func Join (args ...[]byte) []byte {
    return bytes.Join(args, []byte{})
}

// Optional metadata prepended by Wrap.
type Encoding byte
const RAW  = Encoding(0)
const LEN  = Encoding(1 << 0)
const TYPE = Encoding(1 << 1)

// Prepend a varint length and/or type tag to data.
//
// Examples:
//     Wrap(RAW,      type, data) // Return data
//     Wrap(TYPE,     type, data) // Return type + data
//     Wrap(LEN,      type, data) // Return len + data
//     Wrap(TYPE+LEN, type, data) // Return len + type + data
func Wrap (enc Encoding, typ uint64, data []byte) []byte {
    if enc&TYPE>0 { data = Join(EncodeVarint(typ), data) }
    if enc&LEN>0 { data = Join(EncodeVarint(uint64(len(data))), data) }
    return data
}

// Deconstruct bytes which were encoded by Wrap.
func Unwrap (enc Encoding, bytes []byte) (typ uint64, data []byte, length int, err error) {
    data = bytes
    if enc & LEN > 0 {
        l, ll := ReadVarint(data)
        if ll == 0 { return 0, nil, 0, errors.New("ran out of bytes while parsing length") }
        length += ll
        end := ll + int(l)
        if end > len(data) { return 0, nil, 0, errors.New("length  > available bytes") }
        data = data[ll:end]
    }
    if enc & TYPE > 0 {
        t, tl := ReadVarint(data)
        typ = t
        if tl == 0 { return 0, nil, 0, errors.New("ran out of bytes while parsing type") }
        length += tl
        data = data[tl:]
    }
    length += len(data)
    return typ, data, length, nil
}
