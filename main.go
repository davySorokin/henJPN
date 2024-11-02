package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	payload, err := readJSONFile("submission.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	email := payload["contact_email"]
	totp, err := GenerateTOTP(email)
	if err != nil {
		fmt.Println("Error generating TOTP:", err)
		return
	}

	auth := base64.StdEncoding.EncodeToString([]byte(email + ":" + totp))

	jsonData, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", "https://api.challenge.hennge.com/challenges/003", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+auth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}

func readJSONFile(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)
	var result map[string]string
	json.Unmarshal(byteValue, &result)
	return result, nil
}
