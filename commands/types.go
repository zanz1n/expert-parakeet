package commands

import "github.com/bwmarrin/discordgo"

type Command struct {
	Data    *discordgo.ApplicationCommand
	Handler func(c *discordgo.Session, e *discordgo.InteractionCreate) error
}
