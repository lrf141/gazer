package ibd

import (
	"encoding/binary"
	"fmt"
	"os"
)

type Xdes struct {
	FileSegmentId   uint64       `json:"fileSegmentId"`
	XdesList        *ListNode    `json:"XdesList"`
	State           uint32       `json:"state"`
	PageStateBitmap []*PageState `json:"pageStateBitmap"`
}

type PageState struct {
	Free  uint8 `json:"free"`
	Clean uint8 `json:"clean"`
}

const pageStateBitmapSizeLimit = 16

func InitXdes() *Xdes {
	return &Xdes{
		XdesList: InitListNode(),
	}
}

func (xdes *Xdes) Read(f *os.File) error {
	var fileSegmentId uint64
	if err := binary.Read(f, binary.BigEndian, &fileSegmentId); err != nil {
		return err
	}
	xdes.FileSegmentId = fileSegmentId

	err := xdes.XdesList.Read(f)
	if err != nil {
		return err
	}

	var state uint32
	if err := binary.Read(f, binary.BigEndian, &state); err != nil {
		return err
	}
	xdes.State = state

	for i := 0; i < pageStateBitmapSizeLimit; i++ {
		var bitmap uint8
		if err := binary.Read(f, binary.BigEndian, &bitmap); err != nil {
			return err
		}
		fmt.Println(bitmap)
		for j := 0; j < 8; j += 2 {
			pageState := InitPageState()
			pageState.Free = refbit(bitmap, uint8(j))
			pageState.Clean = refbit(bitmap, uint8(j+1))
			xdes.PageStateBitmap = append(xdes.PageStateBitmap, pageState)
		}
	}
	return nil
}

func InitPageState() *PageState {
	return &PageState{}
}

func refbit(i uint8, b uint8) uint8 {
	return (i >> b) & 1
}
