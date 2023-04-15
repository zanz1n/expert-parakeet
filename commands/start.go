package commands

import (
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
		var userId string
		if e.User != nil {
			userId = e.User.ID
		} else {
			userId = e.Member.User.ID
		}
		cm.Start(&utils.CallJob{
			GuildId: e.GuildID,
			UserId:  userId,
			Times:   10,
		})

		c.InteractionRespond(e.Interaction, utils.BasicResponse("Hello World!"))
		return nil
	}
}
