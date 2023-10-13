package main

import (
	"github.com/charmbracelet/log"
	socialsync "github.com/rafadc/socialsync/lib"
	"os"
)

func main() {
	log.Info("Starting socialsync")

	var rssFeed = os.Getenv("RSS_FEED")

	client := socialsync.GetTwitterClient()
	feed := socialsync.ParseFeed(rssFeed)
	socialsync.PostTweets(feed, client)
}
