package main

import (
	"github.com/charmbracelet/log"
	socialsync "github.com/rafadc/socialsync/lib"
	"os"
	"strconv"
	"time"
)

func waitBetweenSyncs() time.Duration {
	secondsBetweenSyncs, ok := os.LookupEnv("SECONDS_BETWEEN_SYNCS")
	if ok {
		secondsBetweenSyncsAsInt, err := strconv.Atoi(secondsBetweenSyncs)
		if err != nil {
			log.Fatal(err)
		}
		return time.Duration(secondsBetweenSyncsAsInt) * time.Second
	} else {
		return 30 * time.Minute
	}
}

func main() {
	logLevel, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		logLevel = "info"
	}
	log.ParseLevel(logLevel)
	log.Info("Starting socialsync")

	var rssFeed = os.Getenv("MASTODON_RSS_URL")
	client := socialsync.GetTwitterClient()

	for {
		log.Info("Fetching feed")
		feed := socialsync.ParseFeed(rssFeed)
		socialsync.PostTweets(feed, client)
		socialsync.UpdateSyncDate()

		time.Sleep(waitBetweenSyncs())
	}
}
