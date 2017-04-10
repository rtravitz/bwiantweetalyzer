package main

import (
	"fmt"
)

func main() {
	tweets := GetBrianTweets()

	var combinedTweets string

	for tweet := range tweets {
		combinedTweets += (" " + tweets[tweet].Text)
	}
	fmt.Println(combinedTweets)
	sentiment := FindSentiment(combinedTweets)
	fmt.Println(sentiment)
}
