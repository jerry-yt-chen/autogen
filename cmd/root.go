package cmd

import (
	"embed"
	"os"

	"github.com/spf13/cobra"
)

func Execute(f embed.FS) {
	rootCmd := cmd(f)
	err := rootCmd.Execute()
	if err == nil {
		return
	}
	os.Exit(-1)
}

func cmd(f embed.FS) *cobra.Command {
	rootCmd := &cobra.Command{
		SilenceErrors: true,
		Use:           "autogen",
		Short:         "generate project",
	}

	newCmd := ProvideNewCmd(f)
	rootCmd.AddCommand(newCmd.Command())
	return rootCmd
}
