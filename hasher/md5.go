package hasher

import (
	"crypto/md5"
	"fmt"
)

type MD5 struct{}

func NewMD5() MD5 {
	return MD5{}
}

func (m MD5) Encrypt(text string) (string, error) {
	h := md5.New()
	h.Write([]byte(text))
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs), nil
}
