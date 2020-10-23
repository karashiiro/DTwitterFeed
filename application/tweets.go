package application

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/bregydoc/gtranslate"

	"github.com/taruti/langdetect"

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
			postTweets(client, config, &config.Subscriptions[i], tweets)
		}

		config.Save()
		time.Sleep(5000)
	}
}

func getTweets(sub *configuration.Subscription, tweets chan *twitterscraper.Result) {
	for tweet := range twitterscraper.GetTweets(context.Background(), sub.UserID, 3) {
		if sub.LastTweetTime >= getTweetTimestamp(tweet.ID).Unix() {
			break
		}

		log.Println("Got tweet:", tweet.Text)

		tweets <- tweet
		sub.LastTweetTime = getTweetTimestamp(tweet.ID).Unix()
	}

	close(tweets)
}

func postTweets(client *discordgo.Session, config *configuration.Configuration, sub *configuration.Subscription, tweets chan *twitterscraper.Result) {
	for {
		tweet, ok := <-tweets
		if !ok {
			break
		}

		// Send the plain Tweet
		if _, err := client.ChannelMessageSend(sub.ChannelID, tweet.PermanentURL); err != nil {
			log.Println(err)
			return
		}

		lang := langdetect.DetectLanguage([]byte(tweet.Text), "ISO88591").String()
		if lang == config.NativeLanguage || config.NativeLanguage == "" {
			return
		}

		// Make and send the translated Tweet
		translatedText, err := gtranslate.TranslateWithParams(tweet.Text, gtranslate.TranslationParams{From: lang, To: config.NativeLanguage})
		if err != nil {
			log.Println(err)
			return
		}

		translatedText = strings.Replace(translatedText, "…", "", -1) // Google Translate sticks an ellipsis at the end of links for some reason, breaking them

		embed := &discordgo.MessageEmbed{
			Title:       "Translated Text",
			Description: translatedText,
			Color:       0x1DA1F2, // Twitter Blue
		}
		if _, err := client.ChannelMessageSendEmbed(sub.ChannelID, embed); err != nil {
			log.Println(err)
		}
	}
}

func getTweetTimestamp(snowflakeStr string) time.Time {
	snowflake, err := strconv.Atoi(snowflakeStr)
	if err != nil {
		log.Fatalln(err)
	}

	offset := 1288834974657
	timestamp := (snowflake >> 22) + offset
	t := time.Unix(int64(timestamp/1000), 0)
	return t
}
