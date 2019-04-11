package models

import (
	"container/list"
)

type EventType int

const (
	Join EventType = 1 << iota
	Leave
	Message
)

type Event struct {
	Type EventType
	User string
	TimeStamp int64
	Content string
}

const archiveSize = 20

var archive = list.New()

func NewArchive(event Event) {
	if archive.Len() >= archiveSize {
		archive.Remove(archive.Front())
	}
	archive.PushBack(event)
}
