package babel

import "github.com/calder/fiddle"

/********************
***   Key Types   ***
********************/

type PubKey interface {
    Encrypt (*fiddle.Bits) *fiddle.Bits
    Id () *Id
}

type PriKey interface {
    Decrypt (*fiddle.Bits) *fiddle.Bits
    Id () *Id
}

type SimKey interface {
    Encrypt (*fiddle.Bits) *fiddle.Bits
    Decrypt (*fiddle.Bits) *fiddle.Bits
    Id () *Id
}

/*******************
***   KeyCache   ***
*******************/

type KeyCache interface {
    PubKey (id *Id) PubKey
}

type MemKeyCache struct {
    keys map[string]PubKey
}

func (cache *MemKeyCache) PubKey (id *Id) PubKey {
    return cache.keys[id.Dat.RawHex()]
}

/*******************
***   KeyVault   ***
*******************/

type KeyVault interface {
    PriKey (id *Id) PriKey
}

type MemKeyVault struct {
    keys map[string]PriKey
}

func (vault *MemKeyVault) PriKey (id *Id) PriKey {
    return vault.keys[id.Dat.RawHex()]
}

/******************
***   PubKey1   ***
******************/

func (key *PubKey1) Encrypt (dat *fiddle.Bits) *Box {
    return nil // FIXME
}