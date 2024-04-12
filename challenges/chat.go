package challenges

type logProb struct {
	Content struct {
		Token       string   `json:"token"`
		Logprob     int      `json:"logprob"`
		Bytes       []int8   `json:"bytes"`
		TopLogprobs []string `json:"top_logprobs"`
	} `json:"content"`
}

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type choice struct {
	Index     int      `json:"index"`
	Msg       message  `json:"message"`
	Logprobs  *logProb `json:"logprobs"`
	FinReason string   `json:"finish_reason"`
}

type usage struct {
	PromptTokens int `json:"prompt_tokens"`
	ComplTokens  int `json:"completion_tokens"`
	TotalTokens  int `json:"total_tokens"`
}

type chatResponse struct {
	Id          string   `json:"id"`
	Object      string   `json:"object"`
	Created     int      `json:"created"`
	Model       string   `json:"model"`
	Choices     []choice `json:"choices"`
	Usge        usage    `json:"usage"`
	Fingerprint string   `json:"system_fingerprint"`
}

type prompt struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type promptReq struct {
	Model    string   `json:"model"`
	Messages []prompt `json:"messages"`
}
