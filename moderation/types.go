package moderation

type Response struct {
	Passed  bool   `json:"passed"`
	Summary string `json:"summary"`
}

type ChatGPTResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}
