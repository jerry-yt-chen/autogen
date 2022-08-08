package new

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"

	"github.com/17media/autogen/internal/generator"
)

type GenType int32

type AppOptions struct {
	Name     string
	ProjType generator.ProjectType
}

type options struct {
	App  AppOptions
	Type string
}

func parseNewOption(args []string) (options, error) {
	opts := options{}
	if len(args) == 0 {
		return opts, fmt.Errorf("you must enter a name for your new application")
	}
	opts.App = AppOptions{}
	opts.App.Name = strings.TrimSpace(args[0])

	genType := viper.GetString("type")
	if projectType, ok := generator.ProjectTypeMap()[genType]; !ok {
		return opts, fmt.Errorf("project type is unknown")
	} else {
		opts.App.ProjType = projectType
	}

	return opts, nil
}
