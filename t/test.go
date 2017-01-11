package main

import (
	"os"

	cli "gopkg.in/urfave/cli.v2"
)

func main() {
	app := &cli.App{}
	app.Run(os.Args)
}
