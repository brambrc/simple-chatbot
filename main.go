package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)


func main() {
	url := "https://api-inference.huggingface.co/models/eleutherai/gpt-neo-2.7B"
	var jsonStr = []byte(`{"inputs":"Tell me best poem for my girlfriend"}`)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	req.Header.Set("Authorization", "Bearer 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")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)	

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(data))
	}
}