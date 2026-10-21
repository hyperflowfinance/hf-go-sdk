package hyperflow

import (
	"context"
	"sync"
	"time"

	"github.com/coder/websocket"
)

type Subscription struct {
	conn      *websocket.Conn
	ctx       context.Context
	cancel    context.CancelFunc
	readMu    sync.Mutex
	mu        sync.Mutex
	completed bool
	closeOnce sync.Once
	closeErr  error
}

// Close stops the stream and releases its WebSocket connection.
func (s *Subscription) Close() error {
	s.closeOnce.Do(func() {
		s.mu.Lock()
		completed := s.completed
		s.mu.Unlock()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		if !completed {
			s.closeErr = s.write(ctx, wsMessage{ID: "1", Type: "complete"})
		}
		if err := s.conn.Close(websocket.StatusNormalClosure, ""); s.closeErr == nil {
			s.closeErr = err
		}
		s.cancel()
	})
	return s.closeErr
}

func (s *Subscription) abort() {
	s.cancel()
	s.conn.CloseNow()
}
