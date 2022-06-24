package cmd

import "github.com/spf13/cobra"

type Cmd interface {
	Command() *cobra.Command
}
