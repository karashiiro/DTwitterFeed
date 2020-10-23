package commands

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/karashiiro/DTwitterFeed/configuration"
)

// AddTwitter adds a Twitter subscription to the configuration.
func AddTwitter(client *discordgo.Session, message *discordgo.MessageCreate, args []string, config *configuration.Configuration) {
	if len(args) < 2 {
		if _, err := client.ChannelMessageSend(message.ChannelID, fmt.Sprintf("<@%s>, too few arguments!", message.Author.ID)); err != nil {
			log.Println(err)
		}
		return
	}

	userID := args[0]
	channelID := args[1]

	for _, sub := range config.Subscriptions {
		if sub.ChannelID == channelID && strings.ToLower(sub.UserID) == strings.ToLower(userID) {
			if _, err := client.ChannelMessageSend(message.ChannelID, fmt.Sprintf("<@%s>, that subscription already exists!", message.Author.ID)); err != nil {
				log.Println(err)
			}
			return
		}
	}

	config.Subscriptions = append(config.Subscriptions, configuration.Subscription{
		UserID:        userID,
		ChannelID:     channelID,
		LastTweetTime: 0,
	})

	config.Save()

	log.Println("Subscription created for Twitter account", userID, "on channel", channelID)
	if _, err := client.ChannelMessageSend(message.ChannelID, fmt.Sprintf("<@%s>, the subscription was created!", message.Author.ID)); err != nil {
		log.Println(err)
	}
}
