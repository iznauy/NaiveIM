package models

import "container/list"

type EventType int

const (
	join EventType = 1 << iota
	leave
	message
)

type Event struct {
	Type EventType
	User string
	TimeStamp int
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

func getNewEvents(lastReceived int) [] Event {
	events := make([]Event, 0, archive.Len())
	for event := archive.Front(); event != nil; event = event.Next() {
		e := event.Value.(Event)
		if e.TimeStamp > lastReceived {
			events = append(events, e)
		}
	}
	return events
}