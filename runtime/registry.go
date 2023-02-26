package runtime

import (
	"encoding/json"
	"syscall/js"

	"github.com/elacity/crypto-protocol/protocol"
	"github.com/pkg/errors"
)

func Register() {
	global := js.Global()

	global.Get("console").Call("log", "Binding utils...")
	global.Set("acquireLicense", AsyncFunc(acquireLicense))
}

func acquireLicense(value js.Value, args []js.Value) (interface{}, error) {
	req := new(protocol.Request)

	// in javascript, the call signature is below:
	// acquireLicense(payload, pssh);
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

		if err := json.Unmarshal(psshBuf, req); err != nil {
			return nil, errors.Wrapf(err, "failed to decode pssh body")
		}
	}

	response := <-protocol.InitLicenseRequest(req)

	if response.Code() != 200 {
		return nil, errors.New(string(response.Bytes()))
	}

	return response.Bytes(), nil
}
