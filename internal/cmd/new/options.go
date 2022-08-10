package new

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/17media/autogen/internal/generator"
)

type GenType int32

type AppOptions struct {
	Name     string
	ProjType generator.ProjectType
	Type     string
}

func parseNewOption(cmd *cobra.Command, args []string) (AppOptions, error) {
	opts := AppOptions{}
	if len(args) == 0 {
		return opts, fmt.Errorf("you must enter a name for your new application")
	}

	opts.Name = strings.TrimSpace(args[0])

	genType, err := cmd.Flags().GetString("type")
	if err != nil {
		return opts, err
	}

	if projectType, ok := generator.ProjectTypeMap()[genType]; !ok {
		return opts, fmt.Errorf("project type is unknown")
	} else {
		opts.ProjType = projectType
	}

	return opts, nil
}
