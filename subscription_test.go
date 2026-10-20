package hyperflow

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/coder/websocket"
)

func TestSubscriptionUsesGraphQLTransportWS(t *testing.T) {
	done := make(chan struct{})
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Test-Auth") != "test-token" {
			t.Errorf("subscription omitted configured auth header")
		}
		conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{Subprotocols: []string{graphqlTransportWS}})
		if err != nil {
			t.Errorf("accept WebSocket: %v", err)
			return
		}
		defer conn.Close(websocket.StatusNormalClosure, "")
		readWSMessage(t, conn, "connection_init")
		writeWSMessage(t, conn, wsMessage{Type: "ping", Payload: json.RawMessage(`{"n":1}`)})
		if pong := readWSMessage(t, conn, "pong"); string(pong.Payload) != `{"n":1}` {
			t.Errorf("unexpected ping response payload %s", pong.Payload)
		}
		writeWSMessage(t, conn, wsMessage{Type: "connection_ack"})
		subscribe := readWSMessage(t, conn, "subscribe")
		var payload struct {
			Query string `json:"query"`
		}
		if err := json.Unmarshal(subscribe.Payload, &payload); err != nil || payload.Query != "subscription { updates { id } }" {
			t.Errorf("unexpected subscription payload: %s (%v)", subscribe.Payload, err)
		}
		writeWSMessage(t, conn, wsMessage{ID: "1", Type: "next", Payload: json.RawMessage(`{"data":{"updates":{"id":"evt-1"}}}`)})
		readWSMessage(t, conn, "complete")
		close(done)
	}))
	defer server.Close()
	client, err := Connect(Config{
		URL: server.URL, HTTPClient: server.Client(),
		Headers: http.Header{"X-Test-Auth": []string{"test-token"}},
	})
	if err != nil {
		t.Fatal(err)
	}
	root := &selectionNode{}
	root.object("updates").field("id")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sub, err := client.Subscribe(ctx, Document{operation: "subscription", root: root})
	if err != nil {
		t.Fatal(err)
	}
	defer sub.Close()
	var result struct {
		Updates struct {
			ID string `json:"id"`
		} `json:"updates"`
	}
	if err := sub.Next(&result); err != nil || result.Updates.ID != "evt-1" {
		t.Fatalf("unexpected result %#v, error %v", result, err)
	}
	if err := sub.Close(); err != nil {
		t.Fatal(err)
	}
	select {
	case <-done:
	case <-time.After(3 * time.Second):
		t.Fatal("server did not receive subscription completion")
	}
}

func TestWebsocketURL(t *testing.T) {
	endpoint, err := websocketURL("https://example.com/graphql?x=1")
	if err != nil || endpoint != "wss://example.com/graphql?x=1" {
		t.Fatalf("unexpected endpoint %q, error %v", endpoint, err)
	}
	if _, err := websocketURL("ftp://example.com/graphql"); err == nil {
		t.Fatal("expected unsupported endpoint scheme to fail")
	}
}
