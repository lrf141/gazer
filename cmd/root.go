package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func Execute(context *cli.Context) error {
	fmt.Println("call this command")
	return nil
}
