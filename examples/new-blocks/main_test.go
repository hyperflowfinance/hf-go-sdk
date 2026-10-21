package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/coder/websocket"
	hyperflow "github.com/hyperflowfinance/hf-go-sdk"
)

func TestWatchReceivesBlockEvents(t *testing.T) {
	serverErrors := make(chan error, 1)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{Subprotocols: []string{"graphql-transport-ws"}})
		if err != nil {
			serverErrors <- err
			return
		}
		defer conn.Close(websocket.StatusNormalClosure, "")
		if _, err := readExampleMessage(conn, "connection_init"); err != nil {
			serverErrors <- err
			return
		}
		if err := writeExampleMessage(conn, wsMessage{Type: "connection_ack"}); err != nil {
			serverErrors <- err
			return
		}
		subscribe, err := readExampleMessage(conn, "subscribe")
		if err != nil || !strings.Contains(string(subscribe.Payload), "newBlocks") {
			if err == nil {
				err = fmt.Errorf("subscription omitted newBlocks")
			}
			serverErrors <- err
			return
		}
		for _, payload := range []json.RawMessage{
			json.RawMessage(`{"data":{"newBlocks":{"number":"1","hash":"h1","timestamp":"t1","gasUsed":"g1"}}}`),
			json.RawMessage(`{"data":{"newBlocks":{"number":"2","hash":"h2","timestamp":null,"gasUsed":null}}}`),
		} {
			if err := writeExampleMessage(conn, wsMessage{
				ID: "1", Type: "next", Payload: payload,
			}); err != nil {
				serverErrors <- err
				return
			}
		}
		if err := writeExampleMessage(conn, wsMessage{ID: "1", Type: "complete"}); err != nil {
			serverErrors <- err
			return
		}
		serverErrors <- nil
	}))
	defer server.Close()
	client, err := hyperflow.Connect(hyperflow.Config{URL: server.URL, HTTPClient: server.Client()})
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var output bytes.Buffer
	if err := watch(ctx, client, &output); err != nil {
		t.Fatal(err)
	}
	want := "{\"gasUsed\":\"g1\",\"hash\":\"h1\",\"number\":\"1\",\"timestamp\":\"t1\"}\n" +
		"{\"gasUsed\":null,\"hash\":\"h2\",\"number\":\"2\",\"timestamp\":null}\n"
	if got := output.String(); got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
	select {
	case err := <-serverErrors:
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("subscription server did not complete")
	}
}
