package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/karashiiro/DTwitterFeed/application"
	"github.com/karashiiro/DTwitterFeed/configuration"
)

func main() {
	client, err := discordgo.New("Bot " + os.Getenv("DTWITTERFEED_BOT_TOKEN"))
	if err != nil {
		log.Fatalln(err)
	}

	config := configuration.LoadConfig()

	messageCreate := application.CreateMessageHandler(config)
	client.AddHandler(messageCreate)

	if err = client.Open(); err != nil {
		log.Fatalln(err)
	}

	go application.ServeTweets(client, config)

	user, err := client.User("@me")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Logged in as", user.Username+"#"+user.Discriminator)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	client.Close()
}
