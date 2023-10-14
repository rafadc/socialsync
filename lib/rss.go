package socialsync

import (
	"github.com/charmbracelet/log"
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
			Content: item.Description,
			Date:    *item.PublishedParsed,
		}
	})
	return Feed{Posts: convertedFeed}
}
