package ibd

import (
	"encoding/json"
	"os"
)

type Tablespace struct {
	FilHeader *FilHeader `json:"filHeader"`
	FspHeader *FspHeader `json:"fspHeader"`
}

func InitTablespace() *Tablespace {
	return &Tablespace{
		FilHeader: InitFilHeader(),
		FspHeader: InitFspHeader(),
	}
}

func (tablespace *Tablespace) Read(f *os.File) error {
	err := tablespace.FilHeader.Read(f)
	if err != nil {
		return err
	}

	err = tablespace.FspHeader.ReadHeader(f)
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
