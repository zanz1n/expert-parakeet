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
			s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
			})

			err := cmd.Handler(s, e)

			if err != nil {
				errStr := err.Error()
				log.Printf(
					"Exception caught when executing command %s\t - %s",
					cmd.Data.Name, errStr,
				)
				s.InteractionResponseEdit(e.Interaction, &discordgo.WebhookEdit{
					Content: &errStr,
				})
			}

			return
		}
	}
}
