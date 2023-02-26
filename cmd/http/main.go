package main

import (
	"encoding/json"
	"flag"
	"mime/multipart"
	"net/http"

	"github.com/elacity/crypto-protocol/protocol"
	"github.com/pkg/errors"
)

var addr = flag.String("addr", ":8090", "http service address")

func main() {
	flag.Parse()

	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/acquireLicense", acquireLicense)
	http.ListenAndServe(*addr, nil)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
	w.WriteHeader(http.StatusOK)
}

func parseInput(raw multipart.FileHeader, dst interface{}) error {
	m, err := raw.Open()
	if err != nil {
		return errors.Wrap(err, "failed to open file")
	}

	mj := json.NewDecoder(m)
	mj.Decode(&dst)

	return nil
}

func acquireLicense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	const mem = (1 << 20) * 512
	err := r.ParseMultipartForm(mem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req := new(protocol.Request)

	if r.MultipartForm.File["pssh"] != nil {
		if err := parseInput(*r.MultipartForm.File["pssh"][0], &req.Pssh); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if r.MultipartForm.File["payload"] != nil {
		if err := parseInput(*r.MultipartForm.File["payload"][0], &req.Payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	response := <-protocol.InitLicenseRequest(req)

	w.WriteHeader(response.Code())
	w.Write(response.Bytes())
}
