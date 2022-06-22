package new

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"autogen/generator"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "new",
		RunE: runE,
	}

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
