// main.go
package main

import (
    "github.com/gin-gonic/gin"
    "wav-flac-converter/handlers"
    "wav-flac-converter/utils"
)

func main() {
    router := gin.Default()

    // WebSocket route for audio conversion
    router.GET("/ws/convert", func(c *gin.Context) {
        handlers.ConvertAudioStream(c, utils.ConvertWAVToFLAC)
    })

    // Start the server
    router.Run(":8080")
}