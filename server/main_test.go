package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartServer(t *testing.T) {
	server, listener, err := StartServer()
	defer func() {
		if listener != nil {
			listener.Close()
		}
		if server != nil {
			server.Stop()
		}
	}()

	assert.NoError(t, err, "Expected no error starting the server")
	assert.NotNil(t, server, "Expected server to be non-nil")
	assert.NotNil(t, listener, "Expected listener to be non-nil")

	// Check if the listener is listening on the correct port
	addr := listener.Addr().String()
	expectedPort := ":50051"
	assert.Contains(t, addr, expectedPort, "Expected listener to bind on port %v", expectedPort)
}
