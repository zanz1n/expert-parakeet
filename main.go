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
	token   string
	signf   = make(chan os.Signal)
	forever = make(chan bool)
	client  *discordgo.Session
)

func init() {
	godotenv.Load()
	token = os.Getenv("TOKEN")
}

func handleInterrupt(sigch chan os.Signal) {
	<-sigch
	log.Println("Stopping ...")
	client.Close()
	forever <- true
}

func main() {
	var err error
	client, err = discordgo.New("Bot " + token)

	if err != nil {
		log.Panicln("Failed to connect to discord\n" + err.Error())
	}

	if err = client.Open(); err != nil {
		log.Panicln("Failed to connect to discord\n" + err.Error())
	}

	signal.Notify(signf, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	go handleInterrupt(signf)

	<-forever
}
