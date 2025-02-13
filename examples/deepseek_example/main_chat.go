package main

import (
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"os"
)

func main() {
	token := os.Getenv("ONE_API_KEY")
	config := openai.DefaultConfig(token)
	config.BaseURL = "http://localhost:3000/v1"
	client := openai.NewClientWithConfig(config)

	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: "deepseek-reasoner",
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleUser, Content: "你好"},
		},
		Stream: false,
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	if len(resp.Choices) > 0 {
		fmt.Println("reasoning_content:", resp.Choices[0].Message.ReasoningContent)
		fmt.Println("content:", resp.Choices[0].Message.Content)
	}

}
