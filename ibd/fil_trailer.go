package ibd

import (
	"encoding/binary"
	"os"
)

type FilTrailer struct {
	OldStyleChecksum uint32 `json:"oldStyleChecksum"`
	Low32bitLsn      uint32 `json:"low32BitLsn"`
}

func InitFilTrailer() *FilTrailer {
	return &FilTrailer{}
}

func (filTrailer *FilTrailer) Read(f *os.File) error {
	var oldStyleChecksum uint32
	if err := binary.Read(f, binary.BigEndian, &oldStyleChecksum); err != nil {
		return nil
	}
	filTrailer.OldStyleChecksum = oldStyleChecksum

	var low32bitLsn uint32
	if err := binary.Read(f, binary.BigEndian, &low32bitLsn); err != nil {
		return nil
	}
	filTrailer.Low32bitLsn = low32bitLsn
	return nil
}
