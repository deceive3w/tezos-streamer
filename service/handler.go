package service

import (
	"encoding/json"
	"net/http"

	"github.com/ecadlabs/tezos-streamer/streamer"
	"github.com/gorilla/websocket"
	"github.com/prometheus/common/log"
)

const (
	maxLimit = 1000000
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

type Handler struct {
	streamer *streamer.Streamer
}

func (h *Handler) Subscribe(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer c.Close()

	sub := h.streamer.NewHeadSubscription()
	defer sub.Close()
	defer log.Debug("Websocket connection exited properly")
	for head := range sub.Stream {
		w, err := c.NextWriter(websocket.TextMessage)
		if err != nil {
			return
		}
		err = json.NewEncoder(w).Encode(head)

		if err != nil {
			log.Error(err.Error())
			return
		}

		err = w.Close()

		if err != nil {
			log.Error(err.Error())
			return
		}
	}
}
