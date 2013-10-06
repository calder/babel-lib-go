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
        dec  := NewDecoder()
        obj  := &IdObj{randId()}
        obj2 := dec.Decode(obj.Encode()).(*IdObj)
        if !obj2.Dat.Equal(obj.Dat) { T.FailNow() }
    }
}