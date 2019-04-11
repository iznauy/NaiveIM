package controllers

import (
	"NaiveIM/backend/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"net/http"
)

type WebSocketController struct {
	beego.Controller
}

// handle web socket
func (controller *WebSocketController) Get()  {
	name := controller.GetString("name")
	if len(name) == 0 {
		return
	}

	ws, err := (&websocket.Upgrader{
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(controller.Ctx.ResponseWriter, controller.Ctx.Request,
		nil)

	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(controller.Ctx.ResponseWriter, "error!", 400)
		return
	} else if err != nil {
		return
	}
	
	Join(name, ws)
	defer Leave(name)

	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		messages <- newEvent(models.Message, name, string(p))
	}

}

func broadcast(event models.Event) {
	data, err := json.Marshal(event)
	if err != nil {
		return
	}

	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		ws := sub.Value.(Subscriber).Conn
		if ws != nil {
			if ws.WriteMessage(websocket.TextMessage, data) != nil {
				unSubscribers <- sub.Value.(Subscriber).Name
			}
		}
	}
}