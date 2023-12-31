// Code generated by goctl. DO NOT EDIT.
package types

type ReqChatCompletion struct {
	Model       string    `json:"model,optional"`
	Messages    []Message `json:"messages"`
	Stream      bool      `json:"stream,optional"`
	MaxTokens   int       `json:"max_tokens,optional"`
	Temperature float32   `json:"temperature,optional"`
	TopP        float32   `json:"top_p,optional"`
	User        string    `json:"user,optional"`
	Seed        int       `json:"seed,optional"`
	Prompt      string    `json:"prompt,optional"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ResChatCompletion struct {
	ID      string   `json:"id"`
	Choices []Choice `json:"choices"`
	Created int      `json:"created"`
	Model   string   `json:"model"`
	Usage   Usage    `json:"usage"`
}

type Usage struct {
	CompletionToken int `json:"completion_tokens"`
	PromptTokens    int `json:"prompt_tokens"`
	TotalTokens     int `json:"total_tokens tokens"`
}

type Choice struct {
	Index   int     `json:"index"`
	Message Message `json:"message,optional"`
	Delta   Message `json:"delta,optional"`
}

type ReqEmbeddings struct {
	Input string `json:"input"`
	User  string `json:"user,optional"`
}

type ResEmbeddings struct {
	Data  []Embedding `json:"data"`
	Model string      `json:"model"`
}

type Embedding struct {
	Embedding []float32 `json:"embedding"`
	Index     int       `json:"index"`
}
