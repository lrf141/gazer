package cmd

import (
	"errors"
	"github.com/lrf141/gazer/ibd"
	"github.com/urfave/cli/v2"
	"os"
)

func Execute(context *cli.Context) error {
	if context.Args().Len() < 1 {
		return errors.New("file path required.")
	}

	path := context.Args().Get(0)
	_, err := os.Stat(path)
	if err != nil {
		return err
	}

	f, err := os.Open(path)
	if err != nil {
		return err
	}

	return ibd.Parse(f)
}
