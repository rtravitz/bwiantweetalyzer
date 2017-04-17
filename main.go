package main

import (
	"encoding/json"
	"net/http"
	"os"
)

func main() {
	port := ":" + os.Getenv("PORT")
	http.HandleFunc("/analyze", AnalyzeSentiment)
	http.ListenAndServe(port, nil)
}

type FullResponse struct {
	Tweets    []Tweet   `json:"tweets"`
	Sentiment Sentiment `json:"sentiment"`
}

func AnalyzeSentiment(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	tweets := GetBrianTweets(params)

	var combinedTweets string

	for tweet := range tweets {
		combinedTweets += (" " + tweets[tweet].Text)
	}
	sentiment := FindSentiment(combinedTweets)
	fullResponse := FullResponse{Tweets: tweets, Sentiment: sentiment}
	w.Header().Set("Access-Control-Allow-Origin", "https://fathomless-mesa-29859.herokuapp.com/")
	json.NewEncoder(w).Encode(fullResponse)
}
