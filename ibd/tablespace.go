package ibd

import (
	"encoding/json"
	"os"
)

type Tablespace struct {
	FspHeader *FspHeader `json:"FspHeader"`
}

func InitTablespace() *Tablespace {
	return &Tablespace{
		FspHeader: InitFspHeader(),
	}
}

func (tablespace *Tablespace) Read(f *os.File) error {
	err := tablespace.FspHeader.ReadHeader(f)
	if err != nil {
		return err
	}
	return nil
}

func (tablespace *Tablespace) EncodeToJson() error {
	encoder := json.NewEncoder(os.Stdout)
	if err := encoder.Encode(tablespace); err != nil {
		return err
	}
	return nil
}
