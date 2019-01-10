package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

func main() {
	// CLI app info
	app := cli.NewApp()
	app.Name = "slatt"
	app.Usage = "a Go tool that helps you easily transfer files from one computer to another"

	// CLI commands
	app.Commands = []cli.Command{
		{
			Name:    "send",
			Aliases: []string{"s"},
			Usage:   "send file",
			Action: func(c *cli.Context) error {
				dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
				handleErr(err, "")
				filename := c.Args().Get(0)
				sendFile(dir, filename)
				return nil
			},
		},
		{
			Name:    "receive",
			Aliases: []string{"r"},
			Usage:   "receive a file",
			Action: func(c *cli.Context) error {
				receiveFile()
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
