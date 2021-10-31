package ibd

import (
	"encoding/binary"
	"os"
)

type ListNode struct {
	PrevPageNumber uint32 `json:"prevPageNumber"`
	PrevPageOffset uint16 `json:"prevPageOffset"`
	NextPageNumber uint32 `json:"nextPageNumber"`
	NextPageOffset uint16 `json:"nextPageOffset"`
}

func InitListNode() *ListNode {
	return &ListNode{}
}

func (listNode *ListNode) Read(f *os.File) error {

	var prevPageNumber uint32
	if err := binary.Read(f, binary.BigEndian, &prevPageNumber); err != nil {
		return err
	}
	listNode.PrevPageNumber = prevPageNumber

	var prevPageOffset uint16
	if err := binary.Read(f, binary.BigEndian, &prevPageOffset); err != nil {
		return err
	}
	listNode.PrevPageOffset = prevPageOffset

	var nextPageNumber uint32
	if err := binary.Read(f, binary.BigEndian, &nextPageNumber); err != nil {
		return err
	}
	listNode.NextPageNumber = nextPageNumber

	var nextPageOffset uint16
	if err := binary.Read(f, binary.BigEndian, &nextPageOffset); err != nil {
		return err
	}
	listNode.NextPageOffset = nextPageOffset
	return nil
}
