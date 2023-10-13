package socialsync

import (
	"github.com/mmcdole/gofeed"
	"github.com/samber/lo"
)

type Feed struct {
	Posts []Post
}

type Post string

func ParseFeed(rssFeed string) Feed {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(rssFeed)
	convertedFeed := lo.Map[*gofeed.Item, Post](feed.Items, func(item *gofeed.Item, index int) Post {
		return Post(item.Description)
	})
	return Feed{Posts: convertedFeed}
}
