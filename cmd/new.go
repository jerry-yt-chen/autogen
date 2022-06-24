package cmd

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/jerry-yt-chen/autogen/generator"
)

type NewCmd struct {
	Commend *cobra.Command
	Fs      embed.FS
}

func (c *NewCmd) Command() *cobra.Command {
	return c.Commend
}

func ProvideNewCmd(files embed.FS) Cmd {
	c := &NewCmd{
		Fs: files,
	}
	c.Commend = &cobra.Command{
		Use:  "new",
		RunE: c.runE,
	}
	return c
}

func (c *NewCmd) runE(cmd *cobra.Command, args []string) error {
	currPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}
	fmt.Printf("currPath:%+v\n", currPath)
	err = generator.NewProjectGenerator(c.Fs).Generate()
	return err
}
