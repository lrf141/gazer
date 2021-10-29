package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := cli.App{
		Name: "gazer",
		Action: func(context *cli.Context) error {
			fmt.Println("call this command")
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err.Error())
	}
}
