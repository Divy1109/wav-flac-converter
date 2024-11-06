// integration_test.go
package main

import (
	"bytes"
	"net/http/httptest"
	"os/exec"
	"testing"
	"wav-flac-converter/handlers"
	"wav-flac-converter/utils"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestWebSocketAudioConversion(t *testing.T) {
    // Check if ffmpeg is installed
    if _, err := exec.LookPath("ffmpeg"); err != nil {
        t.Skip("ffmpeg is not installed, skipping test")
    }

    // Set up a Gin router with the WebSocket route
    router := gin.Default()
    router.GET("/ws/convert", func(c *gin.Context) {
        handlers.ConvertAudioStream(c, utils.ConvertWAVToFLAC)
    })

    // Create a test server
    server := httptest.NewServer(router)
    defer server.Close()

    // Convert the server URL to a WebSocket URL
    wsURL := "ws" + server.URL[len("http"):]

    // Connect to the WebSocket server
    ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
    if err != nil {
        t.Fatalf("Failed to connect to WebSocket server: %v", err)
    }
    defer ws.Close()

    // Sample WAV data (replace this with actual WAV byte data)
    sampleWAVData := []byte{
        // Byte slice from your WAV file
    }

    // Send the WAV data to the server
    err = ws.WriteMessage(websocket.BinaryMessage, sampleWAVData)
    assert.NoError(t, err, "Error writing message to WebSocket")

    // Read the response from the server
    _, flacData, err := ws.ReadMessage()
    assert.NoError(t, err, "Error reading message from WebSocket")

    // Check if the output is not empty
    assert.NotEmpty(t, flacData, "Expected non-empty FLAC data")

    // Optionally, check the FLAC header
    assert.True(t, bytes.HasPrefix(flacData, []byte("fLaC")), "FLAC data does not start with 'fLaC' header")
}