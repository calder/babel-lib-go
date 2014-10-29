package babel

import "math/rand"
import "testing"
import "github.com/calder/fiddle"

func randBits () *fiddle.Bits {
    b := fiddle.FromInt(rand.Int())
    return b.To(rand.Intn(b.Len()+1))
}

func randType () *fiddle.Bits {
    return fiddle.FromInt(rand.Int()).PadLeft(64)
}

func TestDecoder (T *testing.T) {
    for i := 0; i < 1000; i++ {
        id  := randId1()
        id2 := decode(encode(id)).(*Id1)
        if !id2.Dat.Equal(id.Dat) { T.FailNow() }
    }
}