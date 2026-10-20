package hyperflow

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/coder/websocket"
)

func readWSMessage(t *testing.T, conn *websocket.Conn, expected string) wsMessage {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, data, err := conn.Read(ctx)
	if err != nil {
		t.Fatalf("read %s: %v", expected, err)
	}
	var message wsMessage
	if err := json.Unmarshal(data, &message); err != nil || message.Type != expected {
		t.Fatalf("expected %s message, got %s (%v)", expected, data, err)
	}
	return message
}

func writeWSMessage(t *testing.T, conn *websocket.Conn, message wsMessage) {
	t.Helper()
	data, err := json.Marshal(message)
	if err != nil {
		t.Fatal(err)
	}
	if err := conn.Write(context.Background(), websocket.MessageText, data); err != nil {
		t.Fatal(err)
	}
}
