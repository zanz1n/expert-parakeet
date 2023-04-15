package utils

import "github.com/bwmarrin/discordgo"

func BasicResponse(msg string) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
		},
	}
}

func BasicResponseEdit(msg string) *discordgo.WebhookEdit {
	return &discordgo.WebhookEdit{
		Content: &msg,
	}
}
