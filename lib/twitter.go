package socialsync

import (
	"context"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/dghubble/oauth1"
	"github.com/g8rswimmer/go-twitter/v2"
	"net/http"
	"os"
)

type authorizer struct {
	Token string
}

func (a authorizer) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

func PostTweets(feed Feed) {
	client := getTwitterClient()

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

func getTwitterClient() *twitter.Client {
	var twitterKey = os.Getenv("TWITTER_API_KEY")
	var twitterSecret = os.Getenv("TWITTER_API_SECRET")

	var twitterAccessToken = os.Getenv("TWITTER_ACCESS_TOKEN")
	var twitterAccessTokenSecret = os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")

	config := oauth1.NewConfig(twitterKey, twitterSecret)
	httpClient := config.Client(oauth1.NoContext, &oauth1.Token{
		Token:       twitterAccessToken,
		TokenSecret: twitterAccessTokenSecret,
	})

	client := &twitter.Client{
		Authorizer: &authorizer{},
		Client:     httpClient,
		Host:       "https://api.twitter.com",
	}

	return client
}
