package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

type connection struct {
	ws         *websocket.Conn
	sc         chan []byte
	name       string
	totalMoney int
	data       *Data
}

type Data struct {
	Ip       string   `json:"ip"`
	To       string   `json:"to"`
	User     string   `json:"user"`
	From     string   `json:"from"`
	Type     string   `json:"type"`
	Content  string   `json:"content"`
	UserList []string `json:"user_list"`
	Money    int      `json:"money"`
}

type player struct {
	Name  string `json:"name"`
	Money int    `json:"money"`
}

var h = hub{
	con:   make(map[*connection]bool),
	unreg: make(chan *connection),
	msg:   make(chan []byte),
	reg:   make(chan *connection),
}

type hub struct {
	con   map[*connection]bool
	reg   chan *connection
	unreg chan *connection
	msg   chan []byte
}

var player_list []player

func (h hub) run() {
	for {
		select {
		//收到连接信息
		case c := <-h.reg:
			h.con[c] = true
			c.data.Ip = c.ws.RemoteAddr().String()
			c.data.Type = "handshake"
			c.data.UserList = user_list
			data_b, _ := json.Marshal(c.data)
			c.sc <- data_b
		case c := <-h.unreg:
			if _, ok := h.con[c]; ok {
				delete(h.con, c)
				close(c.sc)
			}
			//收到消息
		case data := <-h.msg:
			for c := range h.con {
				select {
				case c.sc <- data:

				default:
					delete(h.con, c)
					close(c.sc)
				}
			}
		}
	}
}

var user_list []string

var player_map = make(map[string]int)

func display() {
	for key, value := range player_map {
		fmt.Println(key, "余额:", value)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		user := request.URL.Query().Get("user")
		player_list = append(player_list, player{
			Name:  user,
			Money: 0,
		})

		player_map[user] = 0
		user_list = append(user_list, user)
		upgrader := &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			}}
		conn, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			return
		}
		c := &connection{
			ws:   conn,
			sc:   make(chan []byte, 256),
			data: &Data{},
		}
		h.reg <- c
		go c.writer()
		c.reader()
		defer func() {
			c.data.Type = "logout"
			user_list = del(user_list, c.data.From)
			c.data.UserList = user_list
			c.data.Content = c.data.From
			data_b, _ := json.Marshal(c.data)
			h.msg <- data_b
			h.unreg <- c
		}()
		defer conn.Close()
	})
	go h.run()
	http.ListenAndServe(":8081", router)
}

func (c *connection) writer() {
	for message := range c.sc {
		c.ws.WriteMessage(websocket.TextMessage, message)
	}
	c.ws.Close()
}

func (c *connection) reader() {
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			h.reg <- c
			break
		}
		json.Unmarshal(message, &c.data)

		switch c.data.Type {
		case "login":
			user_list = append(user_list, c.data.Content)
			c.name = c.data.Content
			c.data.From = c.data.Content
			c.data.UserList = user_list
			data_b, _ := json.Marshal(c.data)
			h.msg <- data_b
		case "user":
			c.data.Type = "user"
			data_b, _ := json.Marshal(c.data)
			money, _ := strconv.Atoi(c.data.Content)
			//加锁
			player_map[c.data.From] -= money
			player_map[c.data.User] += money
			//解锁
			display()
			h.msg <- data_b
		case "logout":
			c.data.Type = "logout"
			user_list = del(user_list, c.data.From)
			data_b, _ := json.Marshal(c.data)
			h.msg <- data_b
			h.unreg <- c
		default:
			fmt.Print("========default================")
		}
	}
}

func del(slice []string, user string) []string {
	count := len(slice)
	if count == 0 {
		return slice
	}
	if count == 1 && slice[0] == user {
		return []string{}
	}
	var n_slice = []string{}
	for i := range slice {
		if slice[i] == user && i == count {
			return slice[:count]
		} else if slice[i] == user {
			n_slice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	fmt.Println(n_slice)
	return n_slice
}
