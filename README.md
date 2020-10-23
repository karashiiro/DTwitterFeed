# DTwitterFeed
A Discord bot that posts to a channel when a Twitter account posts.

## Setup
Install the dependencies, build it with `go build`, and set the environment variable `DTWITTERFEED_BOT_TOKEN` to your Discord bot token.
Also be sure to add the role IDs of your admin roles to the configuration file to restrict feed setup.

## Usage
`^addtwitter <Twitter User ID> <Channel ID>` - Creates a feed for a Twitter account on the specified channel.
`^removetwitter <Twitter User ID> <Channel ID>` - Removes a feed for a Twitter account from the specified channel.