package babel

import "github.com/calder/fiddle"

/*****************
***   PubKey   ***
*****************/

type PubKey interface {
    Enbox (*fiddle.Bits) *Box
    Id () *Id
}

/*******************
***   KeyCache   ***
*******************/

type KeyCache interface {
    GetKey (id *Id)
}

type MemKeyCache struct {
    keys map[string]*PubKey
}

/******************
***   PubKey1   ***
******************/

func (key *PubKey1) Encrypt (dat *fiddle.Bits) *Box {
    return nil // FIXME
}