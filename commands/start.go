package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/zanz1n/bot-inocente/utils"
)

func NewStartCommand(cm *utils.CallJobManager) *Command {
	return &Command{
		Data: &discordgo.ApplicationCommand{
			Name:        "start",
			Description: "Inicia o inferno na terra com algum membro do servidor",
			Type:        discordgo.ChatApplicationCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "user",
					Type:        discordgo.ApplicationCommandOptionUser,
					Description: "O usuário",
				},
			},
		},
		Handler: startCommandHandler(cm),
	}
}

func startCommandHandler(cm *utils.CallJobManager) func(c *discordgo.Session, e *discordgo.InteractionCreate) error {
	return func(c *discordgo.Session, e *discordgo.InteractionCreate) error {
		var user *discordgo.User

		for _, opt := range e.ApplicationCommandData().Options {
			if opt.Name == "user" && opt.Type == discordgo.ApplicationCommandOptionUser {
				user = opt.UserValue(c)
			}
		}

		if user == nil {
			return fmt.Errorf("failed to get user command option")
		}

		cm.Start(&utils.CallJob{
			GuildId: e.GuildID,
			User:    user,
			Times:   10,
		})

		c.InteractionResponseEdit(e.Interaction, utils.BasicResponseEdit("Hello World!"))
		return nil
	}
}
