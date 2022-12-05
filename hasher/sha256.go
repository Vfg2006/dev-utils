package hasher

import (
	"crypto/sha256"
	"fmt"
)

type SHA256 struct{}

func NewSHA256() SHA256 {
	return SHA256{}
}

func (s SHA256) Encrypt(text string) (string, error) {
	h := sha256.New()
	h.Write([]byte(text))
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs), nil
}
