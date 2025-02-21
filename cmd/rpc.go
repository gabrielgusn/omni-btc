package cmd

import (
	"encoding/json"
	"net/http"
	"strings"
)

type RpcRequest struct {
	JsonRpcVersion string        `json:"jsonrpc"`
	Id             string        `json:"id"`
	Backend        string        `json:"backend"`
	Method         string        `json:"method"`
	Params         []interface{} `json:"params"`
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
		/*response, err = */ CallBitcoinCoreRPC(req.Method, req.Params, url, user, pass)
		// if err != nil {
		// fmt.Println("Error:", err)
		// os.Exit(1)
		// }
		// fmt.Println(response)
	}

	if err != nil {
		response = &RpcResponse{Error: err.Error()}
	}

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
