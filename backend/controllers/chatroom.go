package controllers

import "github.com/gorilla/websocket"

type Subscriber struct {
	Name string
	Conn *websocket.Conn
}

