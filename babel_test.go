package babel

import "math/rand"
import "testing"
import "github.com/calder/fiddle"

func randId () *fiddle.Bits {
    return fiddle.FromInt(rand.Int()).Plus(fiddle.FromInt(rand.Int())).Plus(fiddle.FromInt(rand.Int())).Plus(fiddle.FromInt(rand.Int())).PadLeft(256)
}

func randType () *fiddle.Bits {
    return fiddle.FromInt(rand.Int()).PadLeft(64)
}

func TestDecoder (T *testing.T) {
    for i := 0; i < 1000; i++ {
        id  := &Id{randId()}
        id2 := decode(encode(id)).(*Id)
        if !id2.Dat.Equal(id.Dat) { T.FailNow() }
    }
}