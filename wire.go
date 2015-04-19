package babel

import "errors"
import "net"
import "strconv"
import "code.google.com/p/goprotobuf/proto"

// Write a message to a connection.
//
// Wire format:
//   LENGTH - a varint which specifies the number of bytes in TYPE + DATA
//   TYPE   - a varint which identifies the semantics of the DATA part
//   DATA   - the data (empty for singleton types like TRUE, FALSE or NIL)
func WriteMessage (conn net.Conn, typ uint64, data []byte) error {
    // Format message. NIL is encoded as LENGTH = 0 (no TYPE or DATA)
    msg := proto.EncodeVarint(0)
    if typ != NIL {
        msg = append(proto.EncodeVarint(typ), data...)
        msg = append(proto.EncodeVarint(uint64(len(msg))), msg...)
    }

    // Write message
    n, err := conn.Write(msg)
    if err == nil && n < len(msg) {
        err = errors.New("incomplete write (wrote "+
                         strconv.Itoa(n)+" of "+
                         strconv.Itoa(len(msg))+" bytes)")
    }
    return err
}

func WriteSingleton (conn net.Conn, typ uint64) error {
    return WriteMessage(conn, typ, []byte{})
}

func WriteNil (conn net.Conn) error {
    return WriteSingleton(conn, NIL)
}

func Min (x, y int) int {
    if x > y { return y } else { return x }
}

// Attempt to read a message from a connection.
func ReadMessage (conn net.Conn) (typ uint64, data []byte, err error) {
    // Read message length
    length := make([]byte, 8)
    n, err := conn.Read(length)
    if err != nil { return 0, nil, errors.New("message length read error: "+err.Error()) }
    l, n := proto.DecodeVarint(length)
    if n == 0 { return 0, nil, errors.New("ran out of bytes while parsing message length") }
    if l > 1<<20 { return 0, nil, errors.New("message length > 1MB: "+strconv.Itoa(int(l))) }

    // Read message
    typ, data = NIL, make([]byte, 0)
    if l > 0 {
        resp := make([]byte, int(l))
        copy(resp[:Min(int(l), 8-n)], length[n:])
        for N := 8-n; err == nil && N < int(l); N += n {
            n, err = conn.Read(resp[N:])
        }
        if err != nil { return 0, nil, errors.New("read error: "+err.Error()) }

        // Read type
        typ, n = proto.DecodeVarint(resp)
        if n == 0 { return 0, nil, errors.New("ran out of bytes while parsing message type") }

        // Read data
        data = resp[n:]
    }

    return typ, data, nil
}
