package socialsync

import (
	"github.com/charmbracelet/log"
	"github.com/k3a/html2text"
	"github.com/mmcdole/gofeed"
	"github.com/samber/lo"
	"time"
)

type Feed struct {
	Posts []Post
}

type Post struct {
	Content string
	Date    time.Time
}

func ParseFeed(rssFeed string) Feed {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(rssFeed)

	if err != nil {
		log.Fatal(err)
	}

	log.Debugf("Feed read: ", feed)

	convertedFeed := lo.Map[*gofeed.Item, Post](feed.Items, func(item *gofeed.Item, index int) Post {
		return Post{
			Content: html2text.HTML2Text(item.Description),
			Date:    *item.PublishedParsed,
		}
	})
	convertedFeed = lo.Reverse(convertedFeed)
	return Feed{Posts: convertedFeed}
}
