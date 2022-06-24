package main

import (
	"embed"

	"github.com/jerry-yt-chen/autogen/cmd"
)

//go:embed templates
var files embed.FS

func main() {
	cmd.Execute(files)
}
