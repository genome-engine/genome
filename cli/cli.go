package cli

import (
	"github.com/spf13/cobra"
)

func GetCli() *cobra.Command {
	rootCmd.AddCommand(runCmd)
	return rootCmd
}
