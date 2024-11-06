   // audio_handler_test.go
   package handlers

   import (
       "net/http/httptest"
       "testing"
       "github.com/gin-gonic/gin"
       "github.com/gorilla/websocket"
       "github.com/stretchr/testify/assert"
   )

   func TestConvertAudioStream(t *testing.T) {
       // Mock conversion function
       mockConvertFunc := func(wavData []byte) ([]byte, error) {
           // Mock conversion logic: just return the input data for testing
           return wavData, nil
       }

       // Set up a Gin router
       router := gin.Default()
       router.GET("/ws", func(c *gin.Context) {
           ConvertAudioStream(c, mockConvertFunc)
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

       // Send a test WAV message
       testWAVData := []byte("test WAV data")
       err = ws.WriteMessage(websocket.BinaryMessage, testWAVData)
       assert.NoError(t, err, "Error writing message to WebSocket")

       // Read the response
       _, flacData, err := ws.ReadMessage()
       assert.NoError(t, err, "Error reading message from WebSocket")
       assert.Equal(t, testWAVData, flacData, "FLAC data should match the test WAV data")
   }