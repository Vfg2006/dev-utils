package server

import (
	"net/http"

	"github.com/vfg2006/dev-utils/hasher"
)

type HasherRequest struct {
	Text string `json:"text"`
}

func EncryptSha256(sha256 hasher.SHA256) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req HasherRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Error in request", http.StatusBadRequest)
			return
		}

		hash, err := sha256.Encrypt(req.Text)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(hash); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}

func EncryptMd5(md5 hasher.MD5) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req HasherRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Error in request", http.StatusBadRequest)
			return
		}

		hash, err := md5.Encrypt(req.Text)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(hash); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
