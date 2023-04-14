package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func onReady(s *discordgo.Session, e *discordgo.Ready) {
	log.Println("Connected")
}
