package requests

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/Talonmortem/AI-cli-client/AI-cli-client/logger"

	"net/http"
)

var log = logger.For("client")

func SendRequest(req Request) (string, error) {

	log.Printf("Sending request to ollama %v\n", req)

	jsonData, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %v", err)
	}

	resp, err := http.Post("http://localhost:8082/ask", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()
	log.Println("Response received:", resp.Status)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status: %s", resp.Status)
	}
	return "Success", nil
}
