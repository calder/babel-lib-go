package babel

import "testing"

func TestRsaDat (T *testing.T) {
    for i := 0; i < 1000; i++ {
        key := randId()
        dat := randBits()
        en  := EncodeRsaDat(&RsaDat{key,dat})
        de  := DecodeRsaDat(en.To(64), en.From(64)).(*RsaDat)
        if !de.Key.Dat.Equal(key.Dat) || !de.Dat.Equal(dat) {
            T.Log("Key:    ", key)
            T.Log("Dat:    ", dat)
            T.Log("Encoded:", en)
            T.Log("Decoded:", de)
            T.FailNow()
        }
    }
}