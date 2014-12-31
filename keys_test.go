package babel

import "crypto/rsa"
import "crypto/x509"
import "io/ioutil"
import "os"

func parseRsaPrivateKey(filename string) *rsa.PrivateKey {
    file, e := os.Open(filename)
    if e != nil { panic(e) }
    bytes, e := ioutil.ReadAll(file)
    if e != nil { panic(e) }
    key, e := x509.ParsePKCS1PrivateKey(bytes)
    if e != nil { panic(e) }
    return key
}

var rsa2048PrivateKey *rsa.PrivateKey

func init() {
    rsa2048PrivateKey = parseRsaPrivateKey("test-keys/2048.rsaprivatekey")
}