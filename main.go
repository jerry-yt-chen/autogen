package main

import (
	"embed"

	"github.com/17media/autogen/cmd"
)

//go:embed templates
var files embed.FS

func main() {
	cmd.Execute(files)
}
