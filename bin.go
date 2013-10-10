package babel

import "github.com/calder/fiddle"

/**************
***   Bin   ***
**************/

type Bin interface {
    Encode () *fiddle.Bits
    String () string
}

/*****************
***   RawBin   ***
*****************/

type RawBin struct {
    Typ *fiddle.Bits
    Dat *fiddle.Bits
}

func (bin *RawBin) String () string {
    return "<"+bin.Typ.RawHex()+":"+bin.Dat.RawHex()+">"
}

func DecodeRaw (bits *fiddle.Bits) *RawBin {
    return &RawBin{bits.To(64), bits.From(64)}
}

func (bin *RawBin) Encode () *fiddle.Bits {
    return bin.Typ.Plus(bin.Dat)
}