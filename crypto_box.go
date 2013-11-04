package babel

import "github.com/calder/fiddle"

/**************
***   Box   ***
**************/

var BOX = fiddle.FromRawHex("5946F91D56354917")
func init () { AddType(BOX, EncodeBox, DecodeBox) }

type Box struct {
    Key *Id1
    Dat *fiddle.Bits
}

func (dat *Box) String () string {
    return "<Box:"+dat.Key.String()+","+dat.Dat.String()+">"
}

func EncodeBox (val Any) *fiddle.Bits {
    box := val.(*Box)
    return BOX.Plus(fiddle.FromChunks(encode(box.Key), box.Dat))
}

func DecodeBox (typ *fiddle.Bits, dat *fiddle.Bits) Any {
    c := dat.Chunks(2)
    return &Box{decode(c[0]).(*Id1), c[1]}
}