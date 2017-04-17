package main

import (
	"encoding/json"
	"github.com/rs/cors"
	"net/http"
	"os"
)

func main() {
	port := ":" + os.Getenv("PORT")
	mux := http.NewServeMux()
	mux.HandleFunc("/analyze", AnalyzeSentiment)
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(port, handler)
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(fullResponse)
}
