package main

import (
	_ "embed"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

//go:embed VERSION
var VERSION string

//go:generate go run data/data.go
func main() {
	app := &cli.App{
		Name:    "tinytexonfly",
		Version: VERSION,
		Usage:   "Auto install LaTeX packages for TinyTex",
		Authors: []*cli.Author{
			{
				Name:  "batkiz",
				Email: "i@batkiz.com",
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
		},
		Commands: []*cli.Command{
			{
				Name:    "search",
				Aliases: []string{"s"},
				Usage:   "search file globally through the database",
				Action: func(c *cli.Context) error {
					file := c.Args().Get(0)
					Search(file)
					return nil
				},
			},
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
