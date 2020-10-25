package http

type message struct {
	data []byte
	room string
}

type subscription struct {
	conn *connection
	room string
}

// hub maintains the set of active connections and broadcasts messages to the
// connections.
type Hub struct {
	// Registered connections.
	rooms map[string]map[*connection]bool

	// Inbound messages from the connections.
	broadcast chan message

	// Register requests from the connections.
	register chan subscription

	// Unregister requests from connections.
	unregister chan subscription
}

// HubConn ...
var HubConn = Hub{
	broadcast:  make(chan message),
	register:   make(chan subscription),
	unregister: make(chan subscription),
	rooms:      make(map[string]map[*connection]bool),
}

// ServeHub ...
func (HubConn *Hub) run() {
	for {
		select {
			case s := <-HubConn.register:
				connections := HubConn.rooms[s.room]
				if connections == nil {
					connections = make(map[*connection]bool)
					HubConn.rooms[s.room] = connections
				}
				HubConn.rooms[s.room][s.conn] = true
			case s := <-HubConn.unregister:
				connections := HubConn.rooms[s.room]
				if connections != nil {
					if _, ok := connections[s.conn]; ok {
						delete(connections, s.conn)
						close(s.conn.send)
						if len(connections) == 0 {
							delete(HubConn.rooms, s.room)
						}
					}
				}
			case m := <-HubConn.broadcast:
				connections := HubConn.rooms[m.room]
				for c := range connections {
					select {
					case c.send <- m.data:
					default:
						close(c.send)
						delete(connections, c)
						if len(connections) == 0 {
							delete(HubConn.rooms, m.room)
						}
					}
			}
		}
	}
}