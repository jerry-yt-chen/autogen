package new

import (
	"embed"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/17media/autogen/internal/generator"
)

const (
	flagType = "type"
)

type Command struct {
	Commend *cobra.Command
	Fs      embed.FS
}

func Cmd(fs embed.FS) Command {
	c := Command{
		Fs: fs,
	}

	cmd := &cobra.Command{
		Use:   "new [name]",
		Short: "Creates a project template",
		RunE:  c.runE,
	}

	cmd.Flags().String(flagType, "category", "Generate an API server (category-svc style)")
	viper.BindPFlags(cmd.Flags())

	c.Commend = cmd
	return c
}

func (c *Command) runE(_ *cobra.Command, args []string) error {
	opts, err := parseNewOption(args)
	if err != nil {
		return err
	}
	
	gen := generator.NewProjectGenerator(c.Fs, opts.App.Name, opts.App.ProjType)
	if err = gen.Generate(); err != nil {
		return err
	}
	return nil
}
