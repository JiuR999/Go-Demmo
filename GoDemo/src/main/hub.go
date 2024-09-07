package main

import "encoding/json"

var h = hub{
	connection: make(map[*connection]bool),
	unregister: make(chan *connection),
	b:          make(chan []byte),
	register:   make(chan *connection),
}

type hub struct {
	connection map[*connection]bool
	register   chan *connection
	unregister chan *connection
	b          chan []byte
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.register:
			h.connection[c] = true
			c.data.Ip = c.ws.RemoteAddr().String()
			c.data.Type = "handshake"
			c.data.UserList = user_list
			data_b, _ := json.Marshal(c.data)
			c.sc <- data_b
		case c := <-h.unregister:
			if _, ok := h.connection[c]; ok {
				delete(h.connection, c)
				close(c.sc)
			}
		case data := <-h.b:
			for c := range h.connection {
				select {
				case c.sc <- data:
				default:
					delete(h.connection, c)
					close(c.sc)
				}
			}
		}
	}
}
