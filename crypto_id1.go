package babel

import "github.com/calder/fiddle"

/**************
***   Id1   ***
**************/

var ID = fiddle.FromRawHex("823f70579c7a29bf")
func init () { AddType(ID, EncodeId1, DecodeId1) }

type Id1 struct {
    Dat *fiddle.Bits
}

func (id *Id1) String () string {
    return "<Id1:"+id.Dat.RawHex()+">"
}

func (id *Id1) Equal (id2 *Id1) bool {
    return id.Dat.Equal(id2.Dat)
}

func EncodeId1 (val Any) *fiddle.Bits {
    id := val.(*Id1)
    return ID.Plus(id.Dat)
}

func DecodeId1 (typ *fiddle.Bits, dat *fiddle.Bits) Any {
    return &Id1{dat}
}