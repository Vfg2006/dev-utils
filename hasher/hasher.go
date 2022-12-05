package hasher

type Hasher interface {
	Encrypt(text string) (string, error)
}

type Service struct {
	MD5
	SHA256
}

func NewHasher(md5 MD5, sha256 SHA256) Service {
	return Service{
		MD5:    md5,
		SHA256: sha256,
	}
}
