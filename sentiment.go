package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Sentiment struct {
	Probability struct {
		Neg      float64 `json:"neg"`
		Neutral  float64 `json:"neutral"`
		Positive float64 `json:"pos"`
	} `json:"probability"`
	Label string `json:"label"`
}

func FindSentiment(text string) (sentiment Sentiment) {
	client := &http.Client{}
	reqText := "text=" + text
	reqBody := bytes.NewBuffer([]byte(reqText))
	req, err := http.NewRequest("POST", "http://text-processing.com/api/sentiment/", reqBody)
	if err != nil {
		log.Fatal(err, client, req)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err, resp)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err, respBody)
	}
	json.Unmarshal(respBody, &sentiment)
	return
}
