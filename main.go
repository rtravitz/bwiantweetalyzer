package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	port := ":" + os.Getenv("PORT")
	http.HandleFunc("/analyze", AnalyzeSentiment)
	log.Fatal(http.ListenAndServe(port, nil))
}

func AnalyzeSentiment(w http.ResponseWriter, r *http.Request) {
	tweets := GetBrianTweets()

	var combinedTweets string

	for tweet := range tweets {
		combinedTweets += (" " + tweets[tweet].Text)
	}
	sentiment := FindSentiment(combinedTweets)
	json.NewEncoder(w).Encode(sentiment)
}
