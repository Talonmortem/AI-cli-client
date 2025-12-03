package client

import (
	"encoding/json"
	"strings"

	"github.com/Talonmortem/AI-cli-client/AI-cli-client/logger"

	"github.com/Talonmortem/AI-cli-client/AI-cli-client/utils"
	"github.com/gorilla/websocket"
)

var log = logger.For("client")

type WSClient struct {
	Conn                *websocket.Conn
	TokenHandler        func(string)
	FullResponseHandler func(string)
	DoneHandler         func()
}

type ResultMessage struct {
	Response string `json:"response,omitempty"`
	Error    string `json:"error,omitempty"`
	Done     bool   `json:"done,omitempty"`
}

// WSClient wraps a websocket connection with helper methods

// WSClient handles token-streaming WebSocket
// TokenHandler is called for each token chunk
// DoneHandler is fired when stream ends

// NewWSClient connects to WS and creates client
func NewWSClient(url string) (*WSClient, error) {
	url = utils.ProcessWSURL(url)
	conn, err := ConnectWS(url)
	if err != nil {
		return nil, err
	}
	return &WSClient{Conn: conn}, nil
}

// Send sends a text message
func (c *WSClient) Send(msg string) error {
	return c.Conn.WriteMessage(websocket.TextMessage, []byte(msg))
}

// Recv blocks until a message is received
func (c *WSClient) Recv() (string, error) {
	_, data, err := c.Conn.ReadMessage()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// // StartReader parses JSON tokens from WS
func (c *WSClient) StartReader(isGlamour bool) {
	var fullResponse strings.Builder
	go func() {
		for {
			_, data, err := c.Conn.ReadMessage()
			if err != nil {
				log.Println("WS read error:", err)
				if c.DoneHandler != nil {
					c.DoneHandler()
				}
				return
			}

			var msg ResultMessage

			if err := json.Unmarshal(data, &msg); err != nil {
				log.Printf("Error unmarshalling JSON: %v \n\nData: %v\n\n", err, data)
				continue
			}

			if isGlamour {
				fullResponse.WriteString(msg.Response)
			}

			if c.TokenHandler != nil && !isGlamour {
				c.TokenHandler(msg.Response)
			}

			if msg.Done && c.FullResponseHandler != nil {
				c.FullResponseHandler(fullResponse.String())
			}

			if msg.Done && c.DoneHandler != nil {
				c.DoneHandler()
			}

		}

	}()
}
