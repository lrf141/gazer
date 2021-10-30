package ibd

import (
	"os"
)

func Parse(f *os.File) error {
	fspHeader := InitFspHeader()
	err := fspHeader.ReadHeader(f)
	if err != nil {
		return err
	}
	err = fspHeader.EncodeToJson()
	if err != nil {
		return err
	}
	return nil
}
