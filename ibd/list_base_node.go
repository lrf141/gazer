package ibd

import (
	"encoding/binary"
	"os"
)

type ListBaseNode struct {
	ListLength      uint32 `json:"listLength"`
	FirstPageNumber uint32 `json:"firstPageNumber"`
	FirstPageOffset uint16 `json:"firstPageOffset"`
	LastPageNumber  uint32 `json:"lastPageNumber"`
	LastPageOffset  uint16 `json:"lastPageOffset"`
}

func InitListBaseNode() *ListBaseNode {
	return &ListBaseNode{}
}

func (listBaseNode *ListBaseNode) Read(f *os.File) error {
	var listLength uint32
	if err := binary.Read(f, binary.BigEndian, &listLength); err != nil {
		return err
	}
	listBaseNode.ListLength = listLength

	var firstPageNumber uint32
	if err := binary.Read(f, binary.BigEndian, &firstPageNumber); err != nil {
		return err
	}
	listBaseNode.FirstPageNumber = firstPageNumber

	var firstPageOffset uint16
	if err := binary.Read(f, binary.BigEndian, &firstPageOffset); err != nil {
		return err
	}
	listBaseNode.FirstPageOffset = firstPageOffset

	var lastPageNumber uint32
	if err := binary.Read(f, binary.BigEndian, &lastPageNumber); err != nil {
		return err
	}
	listBaseNode.LastPageNumber = lastPageNumber

	var lastPageOffset uint16
	if err := binary.Read(f, binary.BigEndian, &lastPageOffset); err != nil {
		return err
	}
	listBaseNode.LastPageOffset = lastPageOffset
	return nil
}
