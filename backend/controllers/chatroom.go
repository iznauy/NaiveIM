package controllers

import (
	"NaiveIM/backend/models"
	"container/list"
	"github.com/gorilla/websocket"
	"time"
)

type Subscriber struct {
	Name string
	Conn *websocket.Conn
}

var (
	// Channel for new users
	newSubscribers = make(chan Subscriber, 10)
	// Channel for exiting users
	unSubscribers = make(chan string, 10)
	// Messages to be broadcasting
	messages = make(chan models.Event, 10)
	// Online users
	subscribers = list.New()
)

func newEvent(eventType models.EventType, name, message string) models.Event {
	return models.Event{Type: eventType, User: name, TimeStamp: time.Now().Unix(), Content: message}
}

func Join(name string, conn *websocket.Conn) {
	newSubscribers <- Subscriber{Name: name, Conn: conn}
}

func Leave(name string) {
	unSubscribers <- name
}

func isNewUser(name string) bool {
	for ele := subscribers.Front(); ele != nil; ele = ele.Next() {
		if ele.Value.(Subscriber).Name == name {
			return true
		}
	}
	return false
}

func chatroom() {
	for {
		select {
			case sub := <- newSubscribers: // new user
				if isNewUser(sub.Name) {
					subscribers.PushBack(sub)
					messages <- newEvent(models.Join, sub.Name, "")
				}
			case message := <- messages:
				models.NewArchive(message)
				broadcast(message)
			case unSub := <- unSubscribers:
				for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
					if sub.Value.(Subscriber).Name == unSub {
						subscribers.Remove(sub)
						messages <- newEvent(models.Leave, unSub, "")
						break
					}
				}
		}
	}
}

func init() {
	go chatroom()
}