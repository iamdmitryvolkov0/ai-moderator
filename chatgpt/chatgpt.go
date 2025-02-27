package chatgpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"moderator/cfg"
	"net/http"
)

type Request struct {
	Model    string              `json:"model"`
	Store    bool                `json:"store"` // Параметр store, если нужен
	Messages []map[string]string `json:"messages"`
}

type Response struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Error struct {
		Message string `json:"message"`
		Code    string `json:"code"`
	} `json:"error"`
}

func CheckWithOpenAI(text string) (bool, string, error) {
	apiKey := config.AppConfig.OpenAIAPIKey
	url := "https://api.openai.com/v1/chat/completions"

	requestBody := Request{
		Model: "gpt-4o-mini",
		Store: true,
		Messages: []map[string]string{
			{
				"role": "system",
				"content": "Ты модератор комментариев. " +
					"Определи, есть ли мат или токсичность в комментарии. " +
					"Ответь строго в формате {'passed': true/false, 'summary': <description>}. " +
					"Если комментарий чист, верни 'passed': true и описание, что все в порядке. " +
					"Если найден мат или токсичность, верни 'passed': false и описание проблемы.",
			},
			{
				"role":    "user",
				"content": text,
			},
		},
	}

	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return false, "", err
	}
	fmt.Println("Request Body:", string(requestBodyBytes))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return false, "", err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, "", err
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return false, "", err
	}

	if response.Error.Message != "" {
		return false, "Error: " + response.Error.Message, nil
	}

	fmt.Println("Response Body:", response)

	if len(response.Choices) > 0 {
		result := response.Choices[0].Message.Content

		fmt.Println("Model Response Content:", result)

		if result == "" {
			return false, "No response from model", nil
		}

		var parsedResponse map[string]interface{}
		err := json.Unmarshal([]byte(result), &parsedResponse)
		if err != nil {
			return false, "Error parsing response", nil
		}

		passed, ok := parsedResponse["passed"].(bool)
		if !ok {
			return false, "Error: 'passed' not found in response", nil
		}

		summary, ok := parsedResponse["summary"].(string)
		if !ok {
			return false, "Error: 'summary' not found in response", nil
		}

		return passed, summary, nil
	}

	return false, "No response from model", nil
}
