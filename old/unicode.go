package babel

import "github.com/calder/fiddle"

/******************
***   Unicode   ***
******************/

var UNICODE = fiddle.FromRawHex("85847aa769e16613")
func init () { AddType(UNICODE, EncodeUnicode, DecodeUnicode) }

func DecodeUnicode (typ *fiddle.Bits, dat *fiddle.Bits) Any {
    return dat.Unicode()
}

func EncodeUnicode (val Any) *fiddle.Bits {
    str := val.(string)
    return UNICODE.Plus(fiddle.FromUnicode(str))
}