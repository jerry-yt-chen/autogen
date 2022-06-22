package main

import (
	"os"

	"autogen/cmd"
)

var Gopath string

func init() {
	Gopath = os.Getenv("GOPATH")
	if Gopath == "" {
		panic("cannot find $GOPATH environment variable")
	}
}

//func main() {
//	app := cli.NewApp()
//	app.Version = "0.0.0"
//	app.Commands = []*cli.Command{
//		{
//			Name:    "init",
//			Aliases: []string{"i"},
//			Usage:   " Generate scaffold project layout",
//			Action: func(c *cli.Context) error {
//				currPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
//				if err != nil {
//					return err
//				}
//				err = generator.New().Generate(currPath)
//				fmt.Printf("currPath:%+v\n", currPath)
//				if err == nil {
//					fmt.Println("Success Created. Please excute `make up` to start service.", err)
//				}
//
//				return err
//			},
//		},
//	}
//	err := app.Run(os.Args)
//	if err != nil {
//		log.Fatal(err)
//	}
//}

func main() {
	cmd.Execute()
}
