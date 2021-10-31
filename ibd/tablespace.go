package ibd

import (
	"encoding/json"
	"os"
)

type Tablespace struct {
	FilHeader *FilHeader `json:"filHeader"`
	FspHeader *FspHeader `json:"fspHeader"`
	Xdes      []*Xdes      `json:"xdes"`
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

	for i := 0; i < 256; i++ {
		xdes := InitXdes()
		err = xdes.Read(f)
		if err != nil {
			return err
		}
		tablespace.Xdes = append(tablespace.Xdes, xdes)
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
