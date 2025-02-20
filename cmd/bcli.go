package cmd

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

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

// Wrapper function for the Btcd commands
func BtcdCli() {

}
