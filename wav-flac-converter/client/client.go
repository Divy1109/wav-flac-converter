package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	// Connect to the WebSocket server
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws/convert", nil)
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer conn.Close()

	// Read the WAV file
	file, err := os.Open("test/sample.wav")
	if err != nil {
		log.Fatal("File open error:", err)
	}
	defer file.Close()

	wavData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("File read error:", err)
	}

	// Send the WAV data
	err = conn.WriteMessage(websocket.BinaryMessage, wavData)
	if err != nil {
		log.Fatal("Write error:", err)
	}

	// Receive the FLAC data
	_, flacData, err := conn.ReadMessage()
	if err != nil {
		log.Fatal("Read error:", err)
	}

	// Save the FLAC data to a file
	err = ioutil.WriteFile("test/output.flac", flacData, 0644)
	if err != nil {
		log.Fatal("File write error:", err)
	}

	log.Println("FLAC file saved successfully.")
}