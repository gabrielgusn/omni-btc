package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"
)

func CallBitcoinCoreRPC(method string, params []interface{}, rpcUrl string, rpcUser string, rpcPass string) /*(*RpcResponse, error)*/ {

	fmt.Println("Calling Bitcoin Core RPC:", method, params)

	reqBody, err := json.Marshal(RpcRequest{
		JsonRpcVersion: "1.0",
		Id:             "curltest",
		Method:         "getblockhash",
		Params:         []int{10},
	})
	if err != nil {
		// return nil, err
		fmt.Println("Error:", err)
		return

	}

	req, err := http.NewRequest("POST", "http://127.0.0.1:18443", bytes.NewBuffer((reqBody)))
	if err != nil {
		// return nil, err
		fmt.Println("Error on newreq:", err)
		return

	}

	fmt.Println("Will send the request")
	fmt.Println(req)
	req.SetBasicAuth(rpcUser, rpcPass)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		// return nil, err
		fmt.Println("Error on do req:", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// return nil, err
		fmt.Println("Error:", err)
		return
	}

	var rpcResp RpcResponse
	if err := json.Unmarshal(body, &rpcResp); err != nil {
		// return nil, err
		fmt.Println("Error Unmarshal:", err)
	}

	fmt.Println("Response:", rpcResp.Result)
	// return &rpcResp, nil
}

// Wrapper function for spawning bitcoin core commands
func BtcCoreCli(cmd string) ([]byte, error) {
	args := strings.Split(cmd, " ")

	command := exec.Command("bitcoin-cli", args...)

	out, err := command.CombinedOutput()

	if err != nil {
		fmt.Println(string(out))
		return nil, fmt.Errorf("error executing command: %s", cmd)
	}

	if command.ProcessState.Success() {
		return out, nil
	} else {
		return bytes.TrimSpace(out), nil
	}
}
