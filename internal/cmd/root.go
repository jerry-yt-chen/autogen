package cmd

import (
	"embed"
	"os"

	"github.com/spf13/cobra"

	"github.com/17media/autogen/internal/cmd/new"
)

//go:embed templates
var files embed.FS

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

	newCmd := new.Cmd(files)
	rootCmd.AddCommand(newCmd.Commend)
	return rootCmd
}
