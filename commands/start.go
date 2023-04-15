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
					Required:    true,
				},
				{
					Name:        "times",
					Type:        discordgo.ApplicationCommandOptionInteger,
					Description: "Quantas vezes você deseja mover",
					Required:    true,
				},
			},
		},
		Handler: startCommandHandler(cm),
	}
}

func startCommandHandler(cm *utils.CallJobManager) func(c *discordgo.Session, e *discordgo.InteractionCreate) error {
	return func(c *discordgo.Session, e *discordgo.InteractionCreate) error {
		var (
			user  *discordgo.User
			times int = 10
		)

		for _, opt := range e.ApplicationCommandData().Options {
			if opt.Name == "user" && opt.Type == discordgo.ApplicationCommandOptionUser {
				user = opt.UserValue(c)
			}
			if opt.Name == "times" && opt.Type == discordgo.ApplicationCommandOptionInteger {
				times = int(opt.IntValue())
			}
		}

		if user == nil {
			return fmt.Errorf("failed to get user command option")
		}

		cm.Start(&utils.CallJob{
			GuildId: e.GuildID,
			User:    user,
			Times:   times,
		})

		c.InteractionResponseEdit(e.Interaction, utils.BasicResponseEdit("Ok!"))
		return nil
	}
}
