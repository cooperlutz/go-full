package hteeteepee

import (
	"fmt"
	"net/http"
)

// SSEBroker manages Server-Sent Events (SSE) connections and broadcasts events to connected clients utilizing channels.
type SSEBroker struct {
	// Notifier channel for broadcasting events to all connected clients
	Notifier chan []byte
	// New clients will be sent here
	newClients chan chan []byte
	// Closing clients will be sent here
	closingClients chan chan []byte
	// Client connections registry
	clients map[chan []byte]bool
}

// NewSSEBroker creates a new instance of SSEBroker.
func NewSSEBroker() *SSEBroker {
	return &SSEBroker{
		Notifier:       make(chan []byte, 1),
		newClients:     make(chan chan []byte),
		closingClients: make(chan chan []byte),
		clients:        make(map[chan []byte]bool),
	}
}

// Notify sends a byte slice event to all connected clients.
func (b *SSEBroker) Notify(event []byte) {
	b.Notifier <- event
}

// NotifyString sends a string event to all connected clients.
func (b *SSEBroker) NotifyString(event string) {
	b.Notify([]byte(event))
}

// Start begins the SSE broker's main loop, handling new clients, closing clients, and broadcasting events.
func (b *SSEBroker) Start() {
	go func() {
		for {
			select {
			case s := <-b.newClients:
				b.clients[s] = true
			case s := <-b.closingClients:
				delete(b.clients, s)
				close(s)
			case event := <-b.Notifier:
				for clientMessageChan := range b.clients {
					clientMessageChan <- event
				}
			}
		}
	}()
}

// ServeHTTP handles incoming HTTP requests and upgrades them to Server-Sent Events (SSE) connections.
func (b *SSEBroker) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Make sure that the writer supports flushing.
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)

		return
	}

	// Each connection registers its own message channel with the Broker's connections registry
	messageChan := make(chan []byte)
	b.newClients <- messageChan

	// Remove this client from the map of connected clients when this handler exits
	defer func() {
		b.closingClients <- messageChan
	}()

	// Set the headers related to event streaming.
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	for {
		select {
		case <-r.Context().Done():
			return
		case msg := <-messageChan:
			// Write to the ResponseWriter. Server Sent Events requires a specific format, we add "data: " before each message.
			fmt.Fprintf(w, "data: %s\n\n", msg)
			// Flush the data immediately instead of buffering it for later.
			flusher.Flush()
		}
	}
}
