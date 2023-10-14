package socialsync

import (
	"context"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/g8rswimmer/go-twitter/v2"
	"net/http"
	"os"
)

type authorize struct {
	Token string
}

func (a authorize) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

func PostTweets(feed Feed, client *twitter.Client) {
	for _, item := range feed.Posts {
		if item.Date.Before(LatestSyncDate()) {
			continue
		}
		req := twitter.CreateTweetRequest{
			Text: item.Content,
		}

		_, err := client.CreateTweet(context.Background(), req)
		if err != nil {
			log.Fatalf("create tweet error: %v", err)
		}

		log.Debugf("Posted tweet %s \n", item.Content)
	}
}

func GetTwitterClient() *twitter.Client {
	//var twitterKey = os.Getenv("TWITTER_API_KEY")
	//var twitterSecret = os.Getenv("TWITTER_API_SECRET")

	//var twitterAccessToken = os.Getenv("TWITTER_ACCESS_TOKEN")
	var twitterAccessTokenSecret = os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
	client := &twitter.Client{
		Authorizer: authorize{
			Token: twitterAccessTokenSecret,
		},
		Client: http.DefaultClient,
		Host:   "https://api.twitter.com",
	}

	return client
}
