package new

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/jerry-yt-chen/autogen/cmd"
	"github.com/jerry-yt-chen/autogen/generator"
)

type impl struct {
	*cobra.Command
	Fs embed.FS
}

func ProvideNewCmd(files embed.FS) cmd.Cmd {
	cmd := &impl{}

	cmd.Use = "new"
	cmd.RunE = runE
	cmd.Fs = files
	return cmd
}

func runE(cmd *cobra.Command, args []string) error {
	currPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}
	err = generator.New().Generate(currPath)
	fmt.Printf("currPath:%+v\n", currPath)
	if err == nil {
		fmt.Println("Success Created. Please excute `make up` to start service.")
	}
	return err
}
