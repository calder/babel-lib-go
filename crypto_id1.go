// An Id1 is the first 128 bits of the SHA-256 hash of a public key.

package babel

import "bytes"
import "encoding/hex"
import "errors"
import "strconv"

var ID1_STRING = "F7A98013"
var ID1 = NewTypeFromHex(ID1_STRING)
func (*Id1) Type () *Type { return ID1 }
func (*Id1) StringType () string { return ID1_STRING }
func init () { AddType(ID1, decodeId1) }

type Id1 struct {
    data [16]byte
}

func (id* Id1) Data () []byte {
    return id.data[:]
}

func (id *Id1) String () string {
    return "<Id1:"+hex.EncodeToString(id.Data())+">"
}

func (id *Id1) Encode (enc Encoding) []byte {
    return Wrap(enc, ID1, id.Data())
}

func decodeId1 (data []byte) (res Any, err error) { return DecodeId1(data) }
func DecodeId1 (data []byte) (res *Id1, err error) {
    if len(data) != 16 {
        return nil, errors.New("invalid length for Id1: "+strconv.Itoa(len(data)))
    }

    id := &Id1{}
    copy(id.data[:], data)
    return id, nil
}

func (id *Id1) Equal (other *Id1) bool {
    return bytes.Equal(id.Data(), other.Data())
}

func (id *Id1) EqualAny (other Any) bool {
    switch other := other.(type) {
        case *Id1: return id.Equal(other)
        default: return false
    }
}