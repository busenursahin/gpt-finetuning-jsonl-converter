package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type PromptCompletion struct {
	Prompt     string `json:"prompt"`
	Completion string `json:"completion"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatFormat struct {
	Messages []Message `json:"messages"`
}

func main() {

	inputFile, err := os.Open("input.jsonl")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	outputFile, err := os.Create("output.jsonl")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)

	for scanner.Scan() {
		line := scanner.Text()

		var promptCompletion PromptCompletion
		err := json.Unmarshal([]byte(line), &promptCompletion)
		if err != nil {
			fmt.Println("Error unmarshaling line:", err)
			return
		}

		chatFormat := ChatFormat{
			Messages: []Message{
				{
					Role:    "user",
					Content: promptCompletion.Prompt,
				},
				{
					Role:    "assistant",
					Content: promptCompletion.Completion,
				},
			},
		}

		chatJSON, err := json.Marshal(chatFormat)
		if err != nil {
			fmt.Println("Error marshaling to chat format:", err)
			return
		}

		_, err = writer.WriteString(string(chatJSON) + "\n")
		if err != nil {
			fmt.Println("Error writing to output file:", err)
			return
		}
	}

	writer.Flush()

	fmt.Println("Conversion complete! Check output.jsonl")
}
