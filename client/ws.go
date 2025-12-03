package client

import (
	"time"

	"github.com/gorilla/websocket"
)

func ConnectWS(url string) (*websocket.Conn, error) {
	d := websocket.Dialer{
		HandshakeTimeout: 10 * time.Second,
	}

	conn, _, err := d.Dial(url, nil)
	if err != nil {
		return nil, err
	}

	// Configure
	conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	// Start ping goroutine
	go func() {
		t := time.NewTicker(15 * time.Second)
		defer t.Stop()
		for range t.C {
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Println("Ping error:", err)
				return
			}
		}
	}()

	return conn, nil
}
