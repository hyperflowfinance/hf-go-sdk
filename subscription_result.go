package hyperflow

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
)

// Next waits for the next event and decodes it into result.
// It returns io.EOF when the server completes the subscription.
func (s *Subscription) Next(result any) error {
	s.readMu.Lock()
	defer s.readMu.Unlock()
	value := reflect.ValueOf(result)
	if !value.IsValid() || value.Kind() != reflect.Pointer || value.IsNil() {
		return fmt.Errorf("subscription result must be a non-nil pointer")
	}
	for {
		message, err := s.read()
		if err != nil {
			return err
		}
		switch message.Type {
		case "ping":
			if err := s.write(s.ctx, wsMessage{Type: "pong", Payload: message.Payload}); err != nil {
				return fmt.Errorf("reply to subscription ping: %w", err)
			}
		case "next":
			if message.ID != "1" {
				return fmt.Errorf("subscription event has unexpected operation ID %q", message.ID)
			}
			return decodeSubscriptionResult(message.Payload, result)
		case "error":
			if message.ID != "1" {
				return fmt.Errorf("subscription error has unexpected operation ID %q", message.ID)
			}
			return decodeSubscriptionErrors(message.Payload)
		case "complete":
			if message.ID != "1" {
				return fmt.Errorf("subscription completion has unexpected operation ID %q", message.ID)
			}
			s.mu.Lock()
			s.completed = true
			s.mu.Unlock()
			return io.EOF
		case "pong":
			continue
		default:
			return fmt.Errorf("unexpected GraphQL subscription message %q", message.Type)
		}
	}
}

func decodeSubscriptionResult(payload json.RawMessage, result any) error {
	var response struct {
		Data   json.RawMessage `json:"data"`
		Errors GraphQLErrors   `json:"errors"`
	}
	if err := json.Unmarshal(payload, &response); err != nil {
		return fmt.Errorf("decode subscription payload: %w", err)
	}
	if len(response.Errors) != 0 {
		return response.Errors
	}
	if len(response.Data) == 0 || string(response.Data) == "null" {
		return fmt.Errorf("subscription payload is missing data")
	}
	if err := json.Unmarshal(response.Data, result); err != nil {
		return fmt.Errorf("decode subscription data: %w", err)
	}
	return nil
}

func decodeSubscriptionErrors(payload json.RawMessage) error {
	var errors GraphQLErrors
	if err := json.Unmarshal(payload, &errors); err != nil {
		return fmt.Errorf("decode subscription errors: %w", err)
	}
	return errors
}
