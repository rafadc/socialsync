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

## Usage

The easiest way is running from Docker

You need to provide a volume to store the last sync date.

```shell
docker run ghcr.io/rafadc/socialsync:0.1
  -v /path/to/your/data:/data
  -e TWITTER_API_KEY=<your key>
  -e TWITTER_API_SECRET=<your secret>
  -e TWITTER_ACCESS_TOKEN=<your token>
  -e TWITTER_ACCESS_TOKEN_SECRET=<your token secret>
  -e MASTODON_RSS_URL=<your mastodon rss url>
```

### Options

- `BASE_FOLDER` env variable that defaults to `/data` and can be used to change the location of the data
folder. It is useful for development.
- `LOG_LEVEL` env variable to control log level.
- `WAIT_BETWEEN_SYNC` env variable to control the time between syncs. Defaults to 30 minutes.