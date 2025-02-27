package moderation

import (
	"fmt"
	"moderator/chatgpt"
)

type Result struct {
	Passed  bool   `json:"passed"`
	Summary string `json:"summary"`
}

func ModerateComment(text string) (*Result, error) {
	passed, summary, err := chatgpt.CheckWithOpenAI(text)
	if err != nil {
		return nil, fmt.Errorf("error with ChatGPT API: %v", err)
	}

	return &Result{
		Passed:  passed,
		Summary: summary,
	}, nil
}
