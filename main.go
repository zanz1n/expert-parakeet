package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	token  string
	signf  = make(chan os.Signal)
	client *discordgo.Session
)

func handleInterrupt(sigch chan os.Signal) {
	<-sigch
	log.Println("Stopping ...")
	client.Close()
}

func init() {
	godotenv.Load()
	token = os.Getenv("TOKEN")
}

func main() {
	var err error
	client, err = discordgo.New("Bot " + token)

	if err != nil {
		log.Panicln("Failed to connect to discord\n" + err.Error())
	}

	client.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)

	client.AddHandler(onReady)
	client.AddHandler(onInteraction)

	if err = client.Open(); err != nil {
		log.Panicln("Failed to connect to discord\n" + err.Error())
	}

	signal.Notify(signf, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	handleInterrupt(signf)
}
