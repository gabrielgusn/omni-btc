package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// Wrapper function for the getblockhash cli command
func GetBlockHashCmd(height uint32) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "getblockhash",
		Short: "Get block hash by height",
		Run: func(cmd *cobra.Command, args []string) {
			result, err := BtcCoreCli("-rpcport=8334 getblockhash " + strconv.Itoa(int(height)))
			if err != nil {
				fmt.Printf("Error while getting header: %s", err)
			}

			fmt.Println(string(result))
		},
	}

	cmd.Flags().Uint32VarP(&height, "height", "s", 10, "Block Height")

	return cmd
}

func GetBlockHashRpc(height uint32) *cobra.Command {
	// var rpcUser string
	// var rpcPass string
	var cmd = &cobra.Command{
		Use:   "getblockhash",
		Short: "Get block hash by height",
		Run: func(cmd *cobra.Command, args []string) {
			var paramsInterface []interface{}
			paramsInterface = append(paramsInterface, height)
			/*response, err := */ CallBitcoinCoreRPC("getblockhash", paramsInterface, "localhost:18333", "alice", "password")
			// if err != nil {
			// 	fmt.Println("Error while calling rpc:", err)
			// 	return
			// }

			// fmt.Println("Response:", string(response.Result))
		},
	}

	cmd.Flags().Uint32Var(&height, "height", 0, "Block Height")
	// cmd.Flags().StringVarP(&rpcUser, "rpcuser", "u", "", "Block Height")
	// cmd.Flags().StringVarP(&rpcPass, "rpcpass", "p", "", "Block Height")

	return cmd
}

// Wrapper function for the Btcd commands
func BtcdCli() {

}
