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
        bin  := &IdBin{randId()}
        bin2 := dec.Decode(bin.Encode()).(*IdBin)
        if !bin2.Dat.Equal(bin.Dat) { T.FailNow() }
    }
}