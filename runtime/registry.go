package runtime

import (
	"encoding/json"
	"syscall/js"

	"github.com/elacity/crypto-protocol/core"
	"github.com/elacity/crypto-protocol/log"
	"github.com/pkg/errors"
)

func Register() {
	global := js.Global()

	global.Set("acquireLicense", AsyncFunc(acquireLicense))
}

func acquireLicense(value js.Value, args []js.Value) (interface{}, error) {
	req := new(core.Request)

	// in javascript, the call signature is below:
	// acquireLicense(payload, pssh, licreq);
	{
		payloadBuf := make([]byte, args[0].Get("length").Int())
		_ = js.CopyBytesToGo(payloadBuf, args[0])

		if err := json.Unmarshal(payloadBuf, &req.Payload); err != nil {
			return nil, errors.Wrapf(err, "failed to decode payload")
		}
	}

	{
		psshBuf := make([]byte, args[1].Get("length").Int())
		_ = js.CopyBytesToGo(psshBuf, args[1])

		if err := json.Unmarshal(psshBuf, &req.Pssh); err != nil {
			return nil, errors.Wrapf(err, "failed to decode pssh body")
		}
	}

	{
		var rawRequest core.RawFrontRequest
		lrequest := make([]byte, args[2].Get("length").Int())
		_ = js.CopyBytesToGo(lrequest, args[2])

		if err := json.Unmarshal(lrequest, &rawRequest); err != nil {
			return nil, errors.Wrapf(err, "failed to decode request")
		}
		req.Request, _ = core.FromRawToComprehensive(&rawRequest)
	}

	// validation
	log.Tracef("payload: %+v, pssh: %+v, req: %+v", req.Payload, req.Pssh, req.Request)
	if err := validateRequest(req); err != nil {
		return nil, errors.Wrapf(err, "invalid request")
	}

	response := <-core.RequestLicense(req)

	if response.Code() != 200 {
		return nil, errors.New(string(response.Bytes()))
	}

	return response.Bytes(), nil
}

func validateRequest(r *core.Request) error {
	if len(r.Payload.Kids) == 0 {
		return errors.New("no kids in payload")
	}

	if r.Request == nil {
		return errors.New("request is empty")
	}

	if r.Pssh.Data.Authority == "" {
		return errors.New("no authority address")
	}

	if r.Pssh.Data.RpcURL == "" {
		return errors.New("no rpc provided")
	}

	return nil
}
