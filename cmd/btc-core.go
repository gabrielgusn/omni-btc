package cmd

import "fmt"

func CallBitcoinCoreRPC(method string, params []interface{}) (interface{}, string) {

	fmt.Println("Calling Bitcoin Core RPC:", method, params)
	return "Bitcoin Core Response", ""
}
