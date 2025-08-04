package gemma

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type GemmaClient struct {
	BaseURL    string
	HttpClient *http.Client
}

// Request structures for Ollama API
type OllamaGenerateRequest struct {
	Model   string                 `json:"model"`
	Prompt  string                 `json:"prompt"`
	Stream  bool                   `json:"stream"`
	Options map[string]interface{} `json:"options,omitempty"`
}

type OllamaGenerateResponse struct {
	Model     string    `json:"model"`
	CreatedAt time.Time `json:"created_at"`
	Response  string    `json:"response"`
	Done      bool      `json:"done"`
	Context   []int     `json:"context,omitempty"`
}

func (c *GemmaClient) CallGemma(prompt string) error {
	req := &OllamaGenerateRequest{
		Model:   "gemma2:2b",
		Prompt:  prompt,
		Stream:  false,
		Options: nil,
	}
	reqBody, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("unable to marshall request body: %w", err)
	}
	log.Println(string(reqBody))
	resp, err := c.HttpClient.Post(c.BaseURL+"/api/generate", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return fmt.Errorf("error while calling gemma API: %w", err)
	}

	var gemmaResp OllamaGenerateResponse
	if err := json.NewDecoder(resp.Body).Decode(&gemmaResp); err != nil {
		return fmt.Errorf("unable to decode response body: %w", err)
	}

	log.Println("This is the gemma response: ", gemmaResp.Response)
	return nil
}
