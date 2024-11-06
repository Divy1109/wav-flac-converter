package utils

import (
	"bytes"
	"log"
	"os/exec"
)

// ConvertWAVToFLAC uses FFmpeg to convert WAV byte data to FLAC format.
func ConvertWAVToFLAC(wavData []byte) ([]byte, error) {
    // Prepare FFmpeg command
    cmd := exec.Command("ffmpeg", "-i", "pipe:0", "-f", "flac", "pipe:1")

    // Set up pipes to send and receive data
    cmd.Stdin = bytes.NewReader(wavData)
    var flacBuffer bytes.Buffer
    cmd.Stdout = &flacBuffer
	
    // Execute the command
    if err := cmd.Run(); err != nil {
        return nil, err
    }

    // Return the converted FLAC data
	log.Println("FFmpeg conversion successful:", len(flacBuffer.Bytes()), "bytes")
    return flacBuffer.Bytes(), nil
}
