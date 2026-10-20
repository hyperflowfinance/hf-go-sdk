package hyperflow

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"

	"github.com/coder/websocket"
)

const graphqlTransportWS = "graphql-transport-ws"

func (c *Client) Subscribe(ctx context.Context, document Document) (*Subscription, error) {
	if document.operation != "subscription" {
		return nil, fmt.Errorf("expected subscription document")
	}
	query, variables, err := document.render()
	if err != nil {
		return nil, err
	}
	endpoint, err := websocketURL(c.endpoint)
	if err != nil {
		return nil, err
	}
	conn, _, err := websocket.Dial(ctx, endpoint, &websocket.DialOptions{
		HTTPClient: c.http, HTTPHeader: c.headers.Clone(), Subprotocols: []string{graphqlTransportWS},
	})
	if err != nil {
		return nil, fmt.Errorf("connect GraphQL subscription: %w", err)
	}
	if conn.Subprotocol() != graphqlTransportWS {
		conn.CloseNow()
		return nil, fmt.Errorf("server did not accept %q WebSocket subprotocol", graphqlTransportWS)
	}
	conn.SetReadLimit(4 << 20)
	opCtx, cancel := context.WithCancel(ctx)
	sub := &Subscription{conn: conn, ctx: opCtx, cancel: cancel}
	if err := sub.start(ctx, query, normalize(reflect.ValueOf(variables))); err != nil {
		sub.abort()
		return nil, err
	}
	return sub, nil
}

func websocketURL(endpoint string) (string, error) {
	parsed, err := url.Parse(endpoint)
	if err != nil {
		return "", fmt.Errorf("parse GraphQL endpoint: %w", err)
	}
	switch parsed.Scheme {
	case "http":
		parsed.Scheme = "ws"
	case "https":
		parsed.Scheme = "wss"
	default:
		return "", fmt.Errorf("unsupported GraphQL endpoint scheme %q", parsed.Scheme)
	}
	return parsed.String(), nil
}

func (s *Subscription) start(ctx context.Context, query string, variables any) error {
	if err := s.write(ctx, wsMessage{Type: "connection_init"}); err != nil {
		return fmt.Errorf("initialize GraphQL subscription: %w", err)
	}
	if err := s.awaitAck(ctx); err != nil {
		return err
	}
	payload, err := json.Marshal(map[string]any{"query": query, "variables": variables})
	if err != nil {
		return fmt.Errorf("encode GraphQL subscription: %w", err)
	}
	return s.write(ctx, wsMessage{ID: "1", Type: "subscribe", Payload: payload})
}
