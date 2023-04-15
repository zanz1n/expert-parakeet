package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/zanz1n/bot-inocente/utils"
)

func onReady(s *discordgo.Session, e *discordgo.Ready) {
	log.Println("Connected")
}

func onInteraction(s *discordgo.Session, e *discordgo.InteractionCreate) {
	if e.Type == discordgo.InteractionApplicationCommand {
		if cmd, find := utils.GetCommand(e.ApplicationCommandData().Name); find {
			cmd.Handler(s, e)
		}
	}
}
