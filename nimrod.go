package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Nimrod struct {
	Key string
}

type payload struct {
	APIKey  string `json:"api_key"`
	Message string `json:"message"`
}

const (
	url = "https://www.nimrod-messenger.io/api/v1/message"
)

var httpClient = &http.Client{
	Timeout: time.Second * 10,
}

func (n *Nimrod) sendMessage(message string) error {

	data := payload{
		APIKey:  n.Key,
		Message: message,
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("Response status code: %d - %s", resp.StatusCode, result)
	}
	return nil
}
