package babel

import "github.com/calder/fiddle"

/**************
***   Obj   ***
**************/

type Obj interface {
    Encode () *fiddle.Bits
    String () string
}

/*****************
***   RawObj   ***
*****************/

type RawObj struct {
    Typ *fiddle.Bits
    Dat *fiddle.Bits
}

func (bin *RawObj) String () string {
    return "<"+bin.Typ.Hex()+":"+bin.Dat.Hex()+">"
}

func DecodeRaw (bits *fiddle.Bits) *RawObj {
    return &RawObj{bits.To(64), bits.From(64)}
}

func (bin *RawObj) Encode () *fiddle.Bits {
    return bin.Typ.Plus(bin.Dat)
}