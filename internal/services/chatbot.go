package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Estructura para la solicitud a Ollama
type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

// QueryOllama envía una solicitud a Ollama y obtiene una respuesta
func QueryOllama(prompt string) (string, error) {
	url := "http://localhost:11434/api/generate"

	data := OllamaRequest{
		Model:  "deepseek-r1:7b",
		Prompt: prompt,
		Stream: false,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var responseMap map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseMap); err != nil {
		return "", err
	}

	response, ok := responseMap["response"].(string)
	if !ok {
		return "", fmt.Errorf("invalid response format")
	}

	return response, nil
}
