package babel

import "github.com/calder/fiddle"

/********************
***   Key Types   ***
********************/

type PubKey interface {
    Id1 () *Id1
    Encrypt (*fiddle.Bits) *fiddle.Bits
}

type PriKey interface {
    Id1 () *Id1
    Decrypt (*fiddle.Bits) *fiddle.Bits
}

type SimKey interface {
    Id1 () *Id1
    Encrypt (*fiddle.Bits) *fiddle.Bits
    Decrypt (*fiddle.Bits) *fiddle.Bits
}