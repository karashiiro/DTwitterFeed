package application

import (
	"context"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/karashiiro/DTwitterFeed/configuration"
	twitterscraper "github.com/n0madic/twitter-scraper"
)

// ServeTweets fetches tweets and publishes them to the appropriate channels.
func ServeTweets(client *discordgo.Session, config *configuration.Configuration) {
	for {
		for i := range config.Subscriptions {
			tweets := make(chan *twitterscraper.Result)
			go getTweets(&config.Subscriptions[i], tweets)
			postTweets(client, &config.Subscriptions[i], tweets)
		}

		config.Save()
		time.Sleep(5000)
	}
}

func getTweets(sub *configuration.Subscription, tweets chan *twitterscraper.Result) {
	for tweet := range twitterscraper.GetTweets(context.Background(), sub.UserID, 1) {
		if sub.LastTweetID == tweet.ID {
			break
		}

		log.Println("Got tweet:", tweet.Text)

		tweets <- tweet
		sub.LastTweetID = tweet.ID
	}

	close(tweets)
}

func postTweets(client *discordgo.Session, sub *configuration.Subscription, tweets chan *twitterscraper.Result) {
	for {
		tweet, ok := <-tweets
		if !ok {
			break
		}

		if _, err := client.ChannelMessageSend(sub.ChannelID, tweet.PermanentURL); err != nil {
			log.Println(err)
		}
	}
}
