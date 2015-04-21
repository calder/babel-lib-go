package babel

type PubKey interface {
    Id1 () *Hash1
    Encrypt ([]byte) []byte
}

type PriKey interface {
    Id1 () *Hash1
    Decrypt ([]byte) []byte
}

type SimKey interface {
    Id1 () *Hash1
    Encrypt ([]byte) []byte
    Decrypt ([]byte) []byte
}
