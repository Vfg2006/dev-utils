package qrcode

import "github.com/skip2/go-qrcode"

type Service interface {
	Generate(text string) ([]byte, error)
}

type service struct{}

func New() Service {
	return service{}
}

func (qr service) Generate(text string) ([]byte, error) {
	var png []byte
	png, _ = qrcode.Encode(text, qrcode.Medium, 256)

	return png, nil
}
