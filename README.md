# Socialsync

A tiny Go program to synchronize your Mastodon feed with X. 

Mastodon offers a small RSS feed for every user's timeline so we can take advantage of that.

You just have to set a couple of env variables in order to make it work

```shell
export TWITTER_API_KEY=
export TWITTER_API_SECRET=

export TWITTER_ACCESS_TOKEN=
export TWITTER_ACCESS_TOKEN_SECRET=

export MASTODON_RSS_URL=https://evilmeow.com/@rafadc.rss
```

Your profile needs to be public for this to work.

X limits the number of tweets posted per API key to 1500 so that is our hard cap.