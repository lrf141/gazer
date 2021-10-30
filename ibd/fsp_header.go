package ibd

import (
	"encoding/binary"
	"io"
	"os"
)

type FspHeader struct {
	SpaceId                uint32 `json:"spaceId"`
	_Unused                uint32
	HighestPageNumber      uint32        `json:"highestPageNumber"`
	HighestPageNumberInit  uint32        `json:"highestPageNumberInit"`
	Flags                  uint32        `json:"flags"`
	FreeFlagUsedListNumber uint32        `json:"freeFlagUsedListNumber"`
	Free                   *ListBaseNode `json:"free"`
	FreeFlag               *ListBaseNode `json:"freeFlag"`
	FullFlag               *ListBaseNode `json:"fullFlag"`
	NextUnusedSegmentId    uint64        `json:"nextUnusedSegmentId"`
	FullInodes             *ListBaseNode `json:"fullInodes"`
	FreeInodes             *ListBaseNode `json:"freeInodes"`
}

const startPosition = 38

func InitFspHeader() *FspHeader {
	return &FspHeader{
		Free:       InitListBaseNode(),
		FreeFlag:   InitListBaseNode(),
		FullFlag:   InitListBaseNode(),
		FullInodes: InitListBaseNode(),
		FreeInodes: InitListBaseNode(),
	}
}

func (fspHeader *FspHeader) ReadHeader(f *os.File) error {
	if _, err := f.Seek(startPosition, io.SeekStart); err != nil {
		return err
	}

	var spaceId uint32
	if err := binary.Read(f, binary.BigEndian, &spaceId); err != nil {
		return nil
	}
	fspHeader.SpaceId = spaceId

	var unused uint32
	if err := binary.Read(f, binary.BigEndian, &unused); err != nil {
		return err
	}
	fspHeader._Unused = unused

	var highestPageNumber uint32
	if err := binary.Read(f, binary.BigEndian, &highestPageNumber); err != nil {
		return err
	}
	fspHeader.HighestPageNumber = highestPageNumber

	var highestPageNumberInit uint32
	if err := binary.Read(f, binary.BigEndian, &highestPageNumberInit); err != nil {
		return err
	}
	fspHeader.HighestPageNumberInit = highestPageNumberInit

	var flags uint32
	if err := binary.Read(f, binary.BigEndian, &flags); err != nil {
		return err
	}
	fspHeader.Flags = flags

	var freeFlagUsedListNumber uint32
	if err := binary.Read(f, binary.BigEndian, &freeFlagUsedListNumber); err != nil {
		return err
	}
	fspHeader.FreeFlagUsedListNumber = freeFlagUsedListNumber

	if err := fspHeader.Free.Read(f); err != nil {
		return err
	}

	if err := fspHeader.FreeFlag.Read(f); err != nil {
		return err
	}

	if err := fspHeader.FullFlag.Read(f); err != nil {
		return err
	}

	var nextUnusedSegmentId uint64
	if err := binary.Read(f, binary.BigEndian, &nextUnusedSegmentId); err != nil {
		return err
	}
	fspHeader.NextUnusedSegmentId = nextUnusedSegmentId

	if err := fspHeader.FullInodes.Read(f); err != nil {
		return err
	}

	if err := fspHeader.FreeInodes.Read(f); err != nil {
		return err
	}

	return nil
}
