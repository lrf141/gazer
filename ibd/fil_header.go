package ibd

import (
	"encoding/binary"
	"io"
	"os"
)

type FilHeader struct {
	Checksum                uint32 `json:"checksum"`
	PageNumber              uint32 `json:"pageNumber"`
	PreviousPage            uint32 `json:"previousPage"`
	NextPage                uint32 `json:"nextPage"`
	LastPageModificationLSN uint64 `json:"lastPageModificationLSN"`
	PageType                uint16 `json:"pageType"`
	FlushLSN                uint64 `json:"flushLSN"`
	SpaceId                 uint32 `json:"spaceId"`
}

const filHeaderStartPosition = 0

func InitFilHeader() *FilHeader {
	return &FilHeader{}
}

func (filHeader *FilHeader) Read(f *os.File) error {
	if _, err := f.Seek(filHeaderStartPosition, io.SeekStart); err != nil {
		return err
	}

	var checksum uint32
	if err := binary.Read(f, binary.BigEndian, &checksum); err != nil {
		return err
	}
	filHeader.Checksum = checksum

	var pageNumber uint32
	if err := binary.Read(f, binary.BigEndian, &pageNumber); err != nil {
		return err
	}
	filHeader.PageNumber = pageNumber

	var previousPage uint32
	if err := binary.Read(f, binary.BigEndian, &previousPage); err != nil {
		return err
	}
	filHeader.PreviousPage = previousPage

	var nextPage uint32
	if err := binary.Read(f, binary.BigEndian, &nextPage); err != nil {
		return err
	}
	filHeader.NextPage = nextPage

	var lastPageModificationLSN uint64
	if err := binary.Read(f, binary.BigEndian, &lastPageModificationLSN); err != nil {
		return err
	}
	filHeader.LastPageModificationLSN = lastPageModificationLSN

	var pageType uint16
	if err := binary.Read(f, binary.BigEndian, &pageType); err != nil {
		return err
	}
	filHeader.PageType = pageType

	var flushLSN uint64
	if err := binary.Read(f, binary.BigEndian, &flushLSN); err != nil {
		return err
	}
	filHeader.FlushLSN = flushLSN

	var spaceId uint32
	if err := binary.Read(f, binary.BigEndian, &spaceId); err != nil {
		return err
	}
	filHeader.SpaceId = spaceId

	return nil
}
