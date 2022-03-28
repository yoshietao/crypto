package key_test

import (
	"testing"

	"github.com/crypto-ys/key"
)

func BenchmarkPbkdf2(b *testing.B) {
	client := key.NewPbkdf2Client()
	for n := 0; n < b.N; n++ {
		client.GenerateKey([]byte("password"), 32)
	}
}
