package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var rootCommand = &cobra.Command{}
	var height uint32

	var cmd = &cobra.Command{
		Use:   "getblockhash",
		Short: "Get block hash by height",
		Run: func(cmd *cobra.Command, args []string) {
			// if height < 0 {
			// 	panic("Height was not provided")
			// }
			result, err := Bcli("-regtest -rpcuser=alice -rpcpassword=password getblockhash " + strconv.Itoa(int(height)))
			if err != nil {
				fmt.Printf("Error while getting header: %s", err)
			}

			fmt.Println(string(result))
		},
	}

	cmd.Flags().Uint32VarP(&height, "height", "s", 10, "Block Height")

	rootCommand.AddCommand(cmd)
	rootCommand.Execute()
}

func Bcli(cmd string) ([]byte, error) {
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
