// A protobuf compatible variable length integer.
// 
// The first bit of each byte determines whether to read the next byte.
// The last seven bits of each byte are the value, least significant byte first.
//     1******* 1******* 1******* 0*******
//      ^^^^^^^  ^^^^^^^  ^^^^^^^  ^^^^^^^
//  bits 22-28    15-21    7-14      0-6

package babel

import "code.google.com/p/goprotobuf/proto"

func EncodeVarint (n uint64) []byte {
    return proto.EncodeVarint(n)
}

func ReadVarint (data []byte) (x uint64, n int) {
    return proto.DecodeVarint(data)
}