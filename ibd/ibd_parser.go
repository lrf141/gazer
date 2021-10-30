package ibd

import (
	"os"
)

func Parse(f *os.File) error {
	tablespace := InitTablespace()
	err := tablespace.Read(f)
	if err != nil {
		return err
	}

	err = tablespace.EncodeToJson()
	if err != nil {
		return err
	}
	return nil
}
