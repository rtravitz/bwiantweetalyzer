package main

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

var (
	ConsumerKey    = os.Getenv("TWEETKEY")
	ConsumerSecret = os.Getenv("TWEETSECRET")
)

type Tweet struct {
	Text string `json:"text"`
}

type BearerToken struct {
	AccessToken string `json:"access_token"`
}

func GetBrianTweets() []Tweet {
	client := &http.Client{}
	b := getBearerToken()
	twitterEndPoint := "https://api.twitter.com/1.1/statuses/user_timeline.json?screen_name=bafeltra8&count=5&trim_user=true"
	req, err := http.NewRequest("GET", twitterEndPoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", b.AccessToken))

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err, resp)
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var tweets []Tweet
	json.Unmarshal(respBody, &tweets)

	return tweets
}

func getBearerToken() (b BearerToken) {
	client := &http.Client{}
	encodedKeySecret := b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s",
		url.QueryEscape(ConsumerKey),
		url.QueryEscape(ConsumerSecret))))

	reqBody := bytes.NewBuffer([]byte(`grant_type=client_credentials`))
	req, err := http.NewRequest("POST", "https://api.twitter.com/oauth2/token", reqBody)
	if err != nil {
		log.Fatal(err, client, req)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", encodedKeySecret))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Add("Content-Length", strconv.Itoa(reqBody.Len()))

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err, resp)
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err, respBody)
	}

	json.Unmarshal(respBody, &b)
	return
}
