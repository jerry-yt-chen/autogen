package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"autogen/cmd/new"
)

func Execute() {
	rootCmd := cmd()
	err := rootCmd.Execute()
	if err == nil {
		return
	}
	os.Exit(-1)
}

func cmd() *cobra.Command {
	rootCmd := &cobra.Command{
		SilenceErrors: true,
		Use:           "autogen",
		Short:         "generate project",
	}

	newCmd := new.Cmd()
	rootCmd.AddCommand(newCmd)
	return rootCmd
}
