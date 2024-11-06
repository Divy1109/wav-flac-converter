package handlers

import (
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
    "wav-flac-converter/utils"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

// ConvertAudioStream handles incoming WebSocket connections for audio conversion
func ConvertAudioStream(c *gin.Context, convertFunc func([]byte) ([]byte, error)) {
    // Upgrade HTTP connection to WebSocket
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        log.Println("Failed to upgrade to WebSocket:", err)
        return
    }
    defer conn.Close()

    for {
        // Read WAV data from client
        _, wavData, err := conn.ReadMessage()
        if err != nil {
            log.Println("Error reading message:", err)
            break
        }
		log.Println("Received WAV data:", len(wavData), "bytes")
        // Convert WAV to FLAC
        flacData, err := utils.ConvertWAVToFLAC(wavData)
        if err != nil {
            log.Println("Error converting audio:", err)
            continue
        }
		log.Println("Converted FLAC data:", len(flacData), "bytes")
        // Send FLAC data back to the client
        err = conn.WriteMessage(websocket.BinaryMessage, flacData)
        if err != nil {
            log.Println("Error sending FLAC data:", err)
            break
        }
		log.Println("FLAC data sent successfully")
    }
}
