package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/coder/websocket"
)

type wsMessage struct {
	ID      string          `json:"id,omitempty"`
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload,omitempty"`
}

func readExampleMessage(conn *websocket.Conn, expected string) (wsMessage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, data, err := conn.Read(ctx)
	if err != nil {
		return wsMessage{}, err
	}
	var message wsMessage
	err = json.Unmarshal(data, &message)
	if err != nil || message.Type != expected {
		return wsMessage{}, fmt.Errorf("expected %s, got %s (%v)", expected, data, err)
	}
	return message, nil
}

func writeExampleMessage(conn *websocket.Conn, message wsMessage) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return conn.Write(context.Background(), websocket.MessageText, data)
}
