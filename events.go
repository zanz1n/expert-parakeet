package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func onReady(s *discordgo.Session, e *discordgo.Ready) {
	log.Println("Connected")
	PostCommands(s)
}

func onInteraction(s *discordgo.Session, e *discordgo.InteractionCreate) {
	log.Println(e.AppID)
	if e.Type == discordgo.InteractionApplicationCommand {
		if cmd, find := GetCommand(e.ApplicationCommandData().Name); find {
			cmd.Handler(s, e)
			return
		}
	}
}
