package application

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/karashiiro/DTwitterFeed/commands"
	"github.com/karashiiro/DTwitterFeed/configuration"
)

// CreateMessageHandler curries the message creation delegate with the provided application resources.
func CreateMessageHandler(config *configuration.Configuration) func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(client *discordgo.Session, message *discordgo.MessageCreate) {
		messageCreateInternal(client, message, config)
	}
}

func messageCreateInternal(client *discordgo.Session, message *discordgo.MessageCreate, config *configuration.Configuration) {
	if message.Author.Bot {
		return
	}

	member, err := client.GuildMember(message.GuildID, message.Author.ID)
	if err != nil {
		log.Println(err)
		return
	}

	if !rolesOk(member.Roles, config.ModRoles) {
		return
	}

	if message.Content[0:1] != config.Prefix {
		return
	}

	content := message.Content[1:]
	args := strings.Split(content, " ")[1:]

	if strings.HasPrefix(content, "addtwitter") {
		commands.AddTwitter(client, message, args, config)
	} else if strings.HasPrefix(content, "removetwitter") {
		commands.RemoveTwitter(client, message, args, config)
	}
}

func rolesOk(memberRoles []string, modRoles []string) bool {
	ok := false
	for _, role := range memberRoles {
		for _, modRole := range modRoles {
			if modRole == role {
				ok = true
				break
			}
		}
	}
	return ok
}
