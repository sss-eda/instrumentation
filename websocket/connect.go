package websocket

import (
	ws "github.com/gorilla/websocket"
)

// Connection TODO
type Connection struct {
	conn *ws.Conn
}

// Register TODO
func (connection *Connection) Register() error {
	return nil
}

// Connect TODO
func Connect(url string) (*Connection, error) {
	conn, _, err := ws.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}

	return &Connection{
		conn: conn,
	}, nil
}
