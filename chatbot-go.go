package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

)



type Response struct {
	GeneratedText string `json:"generated_text"`
}


func main() {
	apiUrl := "https://api-inference.huggingface.co/models/EleutherAI/gpt-neo-2.7B"
	apiToken := "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"

	reader := bufio.NewReader(os.Stdin)
	client := &http.Client{}

	for {
		fmt.Print("You: ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		if userInput == "exit" {
			break
		}

		jsonStr := []byte(fmt.Sprintf(`{"inputs":"%s"}`, userInput))

		req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonStr))

		if err != nil {
			fmt.Println("Error reading request")
			continue
		}

		req.Header.Set("Authorization", "Bearer "+apiToken)
		req.Header.Set("Content-Type", "application/json")


		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error reading response")
			continue
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading body")
			continue
		}

		var response []Response
		if err := json.Unmarshal(body, &response); err != nil {
			fmt.Println("Error unmarshalling response")
			continue
		}

		if len(response) > 0 {
			fmt.Printf("Bot: %s\n", response[0].GeneratedText)
		} else {
			fmt.Println("Bot: I'm sorry, I don't understand that.")
		}
	}
}