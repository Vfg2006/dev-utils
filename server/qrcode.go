package server

import (
	"net/http"

	"github.com/vfg2006/dev-utils/qrcode"
)

type QRCodeRequest struct {
	Text string `json:"text"`
}

func GenerateQRCode(qrcode qrcode.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req QRCodeRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Error in request", http.StatusBadRequest)
			return
		}

		data, err := qrcode.Generate(req.Text)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// w.Header().Set("Content-Type", "application/octet-stream")
		// w.WriteHeader(http.StatusOK)
		// w.Write(data)

		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
