package key

type KDF interface {
	GenerateKey(keyLen int) ([]byte, error)
}
