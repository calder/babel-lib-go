package babel

import "testing"

func TestBox (T *testing.T) {
    for i := 0; i < 1000; i++ {
        key := randId1()
        dat := randBits()
        en  := EncodeBox(&Box{key,dat})
        de  := DecodeBox(en.To(64), en.From(64)).(*Box)
        if !de.Key.Dat.Equal(key.Dat) || !de.Dat.Equal(dat) {
            T.Log("Key:    ", key)
            T.Log("Dat:    ", dat)
            T.Log("Encoded:", en)
            T.Log("Decoded:", de)
            T.FailNow()
        }
    }
}