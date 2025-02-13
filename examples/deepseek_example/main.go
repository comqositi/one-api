package main

import (
	"context"
	"errors"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"io"
	"os"
)

func main() {
	token := os.Getenv("ONE_API_KEY")
	config := openai.DefaultConfig(token)
	config.BaseURL = "http://localhost:3000/v1"
	client := openai.NewClientWithConfig(config)

	stream, err := client.CreateChatCompletionStream(context.Background(), openai.ChatCompletionRequest{
		Model: "deepseek-reasoner",
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleUser, Content: "你好"},
		},
		Stream: true,
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	defer stream.Close()

	for {
		resp, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("finished")
		}
		if err != nil {
			fmt.Println(err)
			break
		}

		//b, err := json.Marshal(resp)
		//fmt.Println(string(b))

		if len(resp.Choices) > 0 {
			fmt.Println("content:", resp.Choices[0].Delta.ReasoningContent)
			fmt.Println("content:", resp.Choices[0].Delta.Content)
		}
	}
}
