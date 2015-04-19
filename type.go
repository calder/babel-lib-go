// Types are varint tags that preceed data and identify its format.

package babel

import "crypto/rand"
import "encoding/hex"
import "strings"
import "code.google.com/p/goprotobuf/proto"

func Type (s string) uint64 {
    bytes, e := hex.DecodeString(s)
    if (e != nil) { panic(e) }
    typ, n := proto.DecodeVarint(bytes)
    if (n != len(bytes)) { panic("malformed type: "+s) }
    return typ
}

func TypeHex (t uint64) string {
    return strings.ToUpper(hex.EncodeToString(proto.EncodeVarint(t)))
}

func RandType (length int) uint64 {
    bytes := make([]byte, length)
    rand.Read(bytes)
    for i := 0; i < length-1; i++ { bytes[i] |= 128 }
    bytes[length-1] &= 127
    typ, n := proto.DecodeVarint(bytes)
    if (n != len(bytes)) { panic("malformed type: "+hex.EncodeToString(bytes)) }
    return typ
}
