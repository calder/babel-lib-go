package babel

import "github.com/calder/fiddle"

/******************************
***   Bits (don't encode)   ***
******************************/

func init () { AddType(fiddle.Nil(), EncodeBits, nil) }

func EncodeBits (val Any) *fiddle.Bits {
    return val.(*fiddle.Bits)
}