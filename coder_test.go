package babel

import "math/rand"
import "testing"
import "github.com/calder/fiddle"

func randBits () *fiddle.Bits {
    b := fiddle.FromInt(rand.Int())
    return b.To(rand.Intn(b.Len()+1))
}

func randId () *Id {
    return &Id{fiddle.FromInt(rand.Int()).Plus(fiddle.FromInt(rand.Int())).Plus(fiddle.FromInt(rand.Int())).Plus(fiddle.FromInt(rand.Int())).PadLeft(256)}
}

func randType () *fiddle.Bits {
    return fiddle.FromInt(rand.Int()).PadLeft(64)
}

func TestDecoder (T *testing.T) {
    for i := 0; i < 1000; i++ {
        id  := randId()
        id2 := decode(encode(id)).(*Id)
        if !id2.Dat.Equal(id.Dat) { T.FailNow() }
    }
}