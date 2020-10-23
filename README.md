# DTwitterFeed
A Discord bot that reposts tweets into text channels, additionally translating them if necessary.

## Setup
Install the dependencies, build it with `go build`, and set the environment variable `DTWITTERFEED_BOT_TOKEN` to your Discord bot token.
Also be sure to add the role IDs of your admin roles, as strings, to the configuration file that will be generated on launch to restrict feed setup.

## Usage
`^addtwitter <Twitter User ID> <Channel ID>` - Creates a feed for a Twitter account on the specified channel.
`^removetwitter <Twitter User ID> <Channel ID>` - Removes a feed for a Twitter account from the specified channel.
