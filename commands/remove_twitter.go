package commands

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/karashiiro/DTwitterFeed/configuration"
)

// RemoveTwitter removes a Twitter subscription from the configuration.
func RemoveTwitter(client *discordgo.Session, message *discordgo.MessageCreate, args []string, config *configuration.Configuration) {
	if len(args) < 2 {
		if _, err := client.ChannelMessageSend(message.ChannelID, fmt.Sprintf("<@%s>, too few arguments!", message.Author.ID)); err != nil {
			log.Println(err)
		}
		return
	}

	userID := args[0]
	channelID := args[1]

	i := -1
	for j, sub := range config.Subscriptions {
		if sub.ChannelID == channelID && strings.ToLower(sub.UserID) == strings.ToLower(userID) {
			i = j
		}
	}
	if i != -1 {
		config.Subscriptions = *splice(config.Subscriptions, i)
	} else {
		if _, err := client.ChannelMessageSend(message.ChannelID, fmt.Sprintf("<@%s>, no subscription exists for that channel with that Twitter account!", message.Author.ID)); err != nil {
			log.Println(err)
		}
		return
	}

	config.Save()

	log.Println("Subscription to", userID, "removed from channel", channelID)
	if _, err := client.ChannelMessageSend(message.ChannelID, fmt.Sprintf("<@%s>, the subscription was removed!", message.Author.ID)); err != nil {
		log.Println(err)
	}
}

func splice(slice []configuration.Subscription, i int) *[]configuration.Subscription {
	newSlice := make([]configuration.Subscription, len(slice)-1)
	copy(newSlice, slice[0:i])
	copy(newSlice[i:], slice[i+1:])
	return &newSlice
}
