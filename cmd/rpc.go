package cmd

import (
	"encoding/json"
	"net/http"
	"strings"
)

type RpcRequest struct {
	JsonRpcVersion string `json:"jsonrpc"`
	Id             string `json:"id"`
	Backend        string `json:"backend"`
	Method         string `json:"method"`
	Params         []int  `json:"params"`
}

type RpcResponse struct {
	Result interface{} `json:"result"`
	Error  string      `json:"error"`
}

func HandleRpcRequest(w http.ResponseWriter, r *http.Request, url string, user string, pass string) {
	var req RpcRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	var response *RpcResponse
	var err error

	switch strings.ToLower(req.Backend) {
	case "bitcoind":
		CallBitcoinCoreRPC(req.Method, nil, url, user, pass)
		// case "btcd":
		// 	CallBtcdRPC(req.Method, nil, url, user, pass)
	}

	if err != nil {
		response = &RpcResponse{Error: err.Error()}
	}

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
