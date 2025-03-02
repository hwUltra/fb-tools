package http_client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// NewSSEClient creates a new SSE client
func NewSSEClient(timeout time.Duration) *SSEClient {
	return &SSEClient{
		client: NewClient(timeout),
	}
}

// Listen listens for SSE events from the specified URL
func (s *SSEClient) Listen(url string) error {
	resp, err := s.client.Get(url, map[string]string{"Accept": "text/event-stream"})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to connect to SSE: %v", resp.Status)
	}

	// Process the SSE events
	decoder := json.NewDecoder(resp.Body)
	for {
		var event map[string]interface{}
		if err := decoder.Decode(&event); err == io.EOF {
			break // Connection closed
		} else if err != nil {
			return err
		}

		// Handle the event
		s.handleEvent(event)
	}

	return nil
}

// handleEvent processes individual SSE events
func (s *SSEClient) handleEvent(event map[string]interface{}) {
	// Implement your (or user-defined) event handling logic here
	fmt.Println("Received event:", event)
}
