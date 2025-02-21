package cmd

import (
	"encoding/json"
	"net/http"
	"strings"
)

type RpcRequest struct {
	Backend string        `json:"backend"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type RpcResponse struct {
	Result interface{} `json:"result"`
	Error  string      `json:"error,omitempty"`
}

func HandleRpcRequest(w http.ResponseWriter, r *http.Request) {
	var req RpcRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "INvalid request format", http.StatusBadRequest)
		return
	}

	var response RpcResponse

	switch strings.ToLower(req.Backend) {
	case "bitcoind":
		response.Result, response.Error = CallBitcoinCoreRPC(req.Method, req.Params)
	}
}
