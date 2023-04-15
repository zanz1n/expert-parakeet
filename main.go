package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/zanz1n/bot-inocente/commands"
	"github.com/zanz1n/bot-inocente/utils"
)

var (
	token  string
	signf  chan os.Signal
	client *discordgo.Session
)

func init() {
	godotenv.Load()
	signf = make(chan os.Signal)
	token = os.Getenv("DISCORD_TOKEN")
}

func main() {
	var err error
	client, err = discordgo.New("Bot " + token)

	if err != nil {
		log.Panicln("Failed to connect to discord\n" + err.Error())
	}

	client.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)

	cjm := utils.NewCallJobManager(client)

	for i := 0; i < 10; i++ {
		go cjm.AttachListenner()
	}

	client.AddHandler(onReady)
	client.AddHandler(onInteraction)

	AddCommand(commands.NewStartCommand(cjm))

	if err = client.Open(); err != nil {
		log.Panicln("Failed to connect to discord\n" + err.Error())
	}

	defer client.Close()

	signal.Notify(signf, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-signf

	log.Println("Stopping ...")
}
