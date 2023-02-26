package protocol

type PsshBody struct {
	Type    string `json:"protectionType"`
	Variant string `json:"variant"`
	Data    struct {
		Authority string `json:"authority"`
		ChainId   uint16 `json:"chainId"`
		RpcURL    string `json:"rpc"`
	} `json:"data"`
}

type LicensePayload struct {
	Kids []string `json:"kids"`
	Type string   `json:"type"` // "temporary" or "permanent"
}

type Request struct {
	Pssh    PsshBody
	Payload LicensePayload
}

func InitLicenseRequest(r *Request) <-chan Response {
	sl := NewSellex(r.Pssh.Data.Authority, r.Pssh.Data.RpcURL)
	return sl.AwaitFor(r.Payload)
}
