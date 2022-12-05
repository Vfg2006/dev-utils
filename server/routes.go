package server

import (
	"net/http"

	"github.com/vfg2006/dev-utils/hasher"
	"github.com/vfg2006/dev-utils/qrcode"
	"github.com/vfg2006/dev-utils/server/router"
)

func Healthcheck() []router.Route {
	return []router.Route{
		{
			Path:    "/healthcheck",
			Method:  http.MethodGet,
			Handler: HealthcheckHandler(),
		},
	}
}

func EncryptRoutes(hasherService hasher.Service) []router.Route {
	return []router.Route{
		{
			Path:    "/v1/encrypt/sha256",
			Method:  http.MethodPost,
			Handler: EncryptSha256(hasherService.SHA256),
		},
		{
			Path:    "/v1/encrypt/md5",
			Method:  http.MethodPost,
			Handler: EncryptMd5(hasherService.MD5),
		},
	}
}

func QRCodeRoutes(qrcode qrcode.Service) []router.Route {
	return []router.Route{
		{
			Path:    "/v1/qrcode/generate",
			Method:  http.MethodPost,
			Handler: GenerateQRCode(qrcode),
		},
	}
}
