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

func main() {
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

	type BearerToken struct {
		AccessToken string `json:"access_token"`
	}

	var b BearerToken
	json.Unmarshal(respBody, &b)

	twitterEndPoint := "https://api.twitter.com/1.1/statuses/user_timeline.json?screen_name=bafeltra8&count=5"
	req, err = http.NewRequest("GET", twitterEndPoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", b.AccessToken))

	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err, resp)
	}

	defer resp.Body.Close()
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("response status:", resp.Status)
	fmt.Println("response headers:", resp.Header)
	fmt.Println(string(respBody))
}
