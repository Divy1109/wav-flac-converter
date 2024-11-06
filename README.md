

# WAV-FLAC Converter

This project provides a WebSocket-based service for converting WAV audio files to FLAC format using FFmpeg. The service is implemented in Go, utilizing the Gin framework.

## Overview

The WAV-FLAC Converter service enables real-time audio conversion via a WebSocket connection. Clients can stream WAV data and receive FLAC data in return.

## API Endpoints

### WebSocket Endpoint

- **Endpoint**: `/ws/convert`
- **Method**: `GET`
- **Description**: Upgrades the HTTP connection to a WebSocket connection for audio data streaming. Accepts WAV data and returns FLAC data.

#### Usage

1. Connect to the WebSocket server at `ws://localhost:8080/ws/convert`.
2. Send WAV data as a binary message.
3. Receive FLAC data as a binary message.

## Setup Instructions

### Prerequisites

- Go 1.16 or later
- FFmpeg installed and accessible in your system's PATH

### Running Locally

1. **Clone the repository**:

    ```bash
    git clone <https://github.com/Divy1109/wav-flac-converter>
    cd wav-flac-converter
    ```

2. **Install dependencies**:

    ```bash
     go get -u github.com/go-audio/audio
     go get -u github.com/go-audio/wav
    ```

3. **Run the server**:

    ```bash
    go run main.go
    ```

4. **Test the service**:
   - Use the provided `client.go` script to send a sample WAV file and receive the converted FLAC file.

## Deployment

1. **Build the application**:

    ```bash
    go build -o wav-flac-converter
    ```

2. **Deploy the binary** to your server environment.
3. Ensure FFmpeg is installed on the server.
4. **Run the application**:

    ```bash
    ./wav-flac-converter
    ```

## Testing Strategy

### Unit Tests

- **Location**: `handlers/audio_handler_test.go`, `utils/converter_test.go`
- **Run tests**:

    ```bash
    go test ./...
    ```

### Integration Tests

- **Location**: `integration_test.go`
- **Description**: Tests the complete WebSocket conversion flow.
- **Run tests**:

    ```bash
    go test ./...
    ```

### Example Test Invocation

1. **Start the server**:

    ```bash
    go run main.go
    ```

2. **Run the client**:

    ```bash
    go run client/client.go
    ```

3. **Verify the output**:
   - Check `test/output.flac` for the converted audio file.

## Additional Notes

- Ensure FFmpeg is correctly installed and accessible in your environment.
- You can modify `client.go` to test with different WAV files as needed.
- The service logs conversion details to the console for monitoring purposes.

--- 
