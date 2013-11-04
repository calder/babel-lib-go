package babel

import "errors"
import "github.com/calder/fiddle"

/**************
***   Nil   ***
**************/

var NIL = fiddle.FromRawHex("0000000000000000")
func init () { AddType(NIL, EncodeNil, DecodeNil) }

func EncodeNil (val Any) *fiddle.Bits {
    if val != nil { panic(errors.New("EncodeNil() called on non-nil")) }
    return NIL
}

func DecodeNil (typ *fiddle.Bits, dat *fiddle.Bits) Any {
    return nil
}