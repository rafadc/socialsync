package socialsync

import (
	"github.com/charmbracelet/log"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"os"
)

func PostTweets(feed Feed, client *twitter.Client) {
	for _, item := range feed.Posts {
		tweet, _, err := client.Statuses.Update(string(item), nil)
		if err != nil {
			log.Errorf("Error: ", err)
			continue
		}

		log.Infof("Posted tweet %s \n", tweet.Text)
	}
}

func GetTwitterClient() *twitter.Client {
	var twitterKey = os.Getenv("TWITTER_API_KEY")
	var twitterSecret = os.Getenv("TWITTER_API_SECRET")

	var twitterAccessToken = os.Getenv("TWITTER_ACCESS_TOKEN")
	var twitterAccessTokenSecret = os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
	config := oauth1.NewConfig(twitterKey, twitterSecret)
	token := oauth1.NewToken(twitterAccessToken, twitterAccessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	return twitter.NewClient(httpClient)
}
