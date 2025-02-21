package main

import (
	"github.com/gabrielgusn/omni-btc/cmd"
	"github.com/spf13/cobra"
)

func main() {
	var rootCommand = &cobra.Command{}
	var height uint32

	rootCommand.AddCommand(cmd.GetBlockHashRpc(height))
	rootCommand.Execute()
}
