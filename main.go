package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

//go:generate go run data/data.go
func main() {
	app := &cli.App{
		Name:    "tinytexonfly",
		Version: "0.1.0",
		Usage:   "Auto install LaTeX packages for TinyTex",
		Authors: []*cli.Author{
			{
				"batkiz",
				"i@batkiz.com",
			},
		},
		Action: func(c *cli.Context) error {
			file := c.Args().Get(0)
			var (
				s string
			)

			if file == "" {
				s = ReadPath(".")

				execute(s)
				return nil
			}
			fi, err := os.Stat(file)
			if err != nil {
				log.Fatal(err)
			}
			if fi.IsDir() {
				s = ReadPath(file)
				execute(s)
				return nil
			} else {
				s = ReadFile(file)
				execute(s)
				return nil
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func execute(s string) {
	err := ExecTlmgrInstall(getPackagesNeedInstall(s))
	if err != nil {
		log.Fatal(err)
	}
}
