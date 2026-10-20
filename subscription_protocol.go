package hyperflow

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/coder/websocket"
)

type wsMessage struct {
	ID      string          `json:"id,omitempty"`
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload,omitempty"`
}

func (s *Subscription) write(ctx context.Context, message wsMessage) error {
	data, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("encode subscription message: %w", err)
	}
	return s.conn.Write(ctx, websocket.MessageText, data)
}

func (s *Subscription) read() (wsMessage, error) {
	messageType, data, err := s.conn.Read(s.ctx)
	if err != nil {
		return wsMessage{}, fmt.Errorf("read GraphQL subscription: %w", err)
	}
	if messageType != websocket.MessageText {
		return wsMessage{}, fmt.Errorf("GraphQL subscription received a non-text message")
	}
	var message wsMessage
	if err := json.Unmarshal(data, &message); err != nil {
		return wsMessage{}, fmt.Errorf("decode subscription message: %w", err)
	}
	return message, nil
}

func (s *Subscription) awaitAck(ctx context.Context) error {
	for {
		messageType, data, err := s.conn.Read(ctx)
		if err != nil {
			return fmt.Errorf("wait for GraphQL subscription acknowledgement: %w", err)
		}
		if messageType != websocket.MessageText {
			return fmt.Errorf("GraphQL subscription received a non-text message")
		}
		var message wsMessage
		if err := json.Unmarshal(data, &message); err != nil {
			return fmt.Errorf("decode subscription acknowledgement: %w", err)
		}
		if message.Type == "ping" {
			if err := s.write(ctx, wsMessage{Type: "pong", Payload: message.Payload}); err != nil {
				return err
			}
			continue
		}
		if message.Type == "error" {
			return decodeSubscriptionErrors(message.Payload)
		}
		if message.Type != "connection_ack" {
			return fmt.Errorf("unexpected GraphQL subscription message %q before acknowledgement", message.Type)
		}
		return nil
	}
}
