package main

import (
	"encoding/json"
	"flag"
	"mime/multipart"
	"net/http"

	"github.com/elacity/crypto-protocol/core"
	"github.com/elacity/crypto-protocol/log"
	"github.com/pkg/errors"
)

var addr = flag.String("addr", ":8090", "http service address")

func main() {
	flag.Parse()

	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/acquireLicense", acquireLicense)

	log.Infof("listening on %s...", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Errorf("failed to start server; %s", err)
	}
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

	req := new(core.Request)

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

	if r.MultipartForm.File["request"] != nil {
		var rawRequest core.RawFrontRequest
		if err := parseInput(*r.MultipartForm.File["request"][0], &rawRequest); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if req.Request, err = core.FromRawToComprehensive(&rawRequest); err != nil {
			log.Warnf("failed to parse request: %s, empty value will be passed", err)
		}
	}

	response := <-core.RequestLicense(req)

	w.WriteHeader(response.Code())
	w.Write(response.Bytes())
}
