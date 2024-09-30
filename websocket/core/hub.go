package core

type Hub struct {
	//上线注册
	Register chan *Client
	//下线注销
	UnRegister chan *Client
	//所有在线客户端的内存地址
	Clients map[*Client]bool
	// Inbound messages from the clients.
	broadcast chan []byte
}

func CreateHubFactory() *Hub {
	return &Hub{
		Register:   make(chan *Client),
		UnRegister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.UnRegister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
		case message := <-h.broadcast:
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}
