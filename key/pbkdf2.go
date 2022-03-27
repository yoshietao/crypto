package key

import (
	"hash"

	"golang.org/x/crypto/pbkdf2"
)

type pbkdf2Client struct {
	iteration int
	salt      []byte
	algorithm func() hash.Hash
}

func NewPbkdf2Client() pbkdf2Client {
	return pbkdf2Client{
		iteration: 5,
		salt:      []byte{},
		algorithm: nil,
	}
}

func (pb *pbkdf2Client) GenerateKey(password []byte, keyLen int) ([]byte, error) {
	key := pbkdf2.Key(password, pb.salt, pb.iteration, keyLen, pb.algorithm)
	return key, nil
}
