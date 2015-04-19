package main

import "os"
import "crypto/rand"
import "crypto/rsa"
import "crypto/x509"
import "strconv"

func NewRsaPrivateKey (bits int) *rsa.PrivateKey {
    key, e := rsa.GenerateKey(rand.Reader, bits)
    if e != nil { panic(e) }
    return key
}

func usage() {
    println("Usage: go run generate.go BITS >> FILENAME")
    os.Exit(1)
}

func main() {
    if len(os.Args) != 2 { usage() }
    bits, e := strconv.Atoi(os.Args[1])
    if e != nil { usage() }

    key := NewRsaPrivateKey(bits)
    bytes := x509.MarshalPKCS1PrivateKey(key)
    _, e = os.Stdout.Write(bytes)
    if e != nil { panic(e) }
}
