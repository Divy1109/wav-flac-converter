// converter_test.go
package utils

import (
    "bytes"
    "os/exec"
    "testing"
)

func TestConvertWAVToFLAC(t *testing.T) {
    // Check if ffmpeg is installed
    if _, err := exec.LookPath("ffmpeg"); err != nil {
        t.Skip("ffmpeg is not installed, skipping test")
    }

    // Sample WAV data (this is just a placeholder; in a real test, use actual WAV data)
    sampleWAVData := []byte{
        // WAV header and some data (this is just an example, not actual audio data)
        0x52, 0x49, 0x46, 0x46, 0x24, 0x08, 0x00, 0x00, 0x57, 0x41, 0x56, 0x45,
        0x66, 0x6D, 0x74, 0x20, 0x10, 0x00, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00,
        0x40, 0x1F, 0x00, 0x00, 0x80, 0x3E, 0x00, 0x00, 0x02, 0x00, 0x10, 0x00,
        0x64, 0x61, 0x74, 0x61, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
    }

    // Call the conversion function
    flacData, err := ConvertWAVToFLAC(sampleWAVData)
    if err != nil {
        t.Fatalf("Conversion failed: %v", err)
    }

    // Check if the output is not empty
    if len(flacData) == 0 {
        t.Error("Expected non-empty FLAC data")
    }

    // Optionally, you can add more checks to verify the FLAC data format
    // For example, check the FLAC header
    if !bytes.HasPrefix(flacData, []byte("fLaC")) {
        t.Error("FLAC data does not start with 'fLaC' header")
    }
}