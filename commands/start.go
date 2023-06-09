package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/zanz1n/expert-parakeet/utils"
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
				{
					Name:        "except",
					Type:        discordgo.ApplicationCommandOptionChannel,
					Description: "Seleciona um canal para não moves o usuário",
					Required:    false,
				},
			},
		},
		Handler: startCommandHandler(cm),
	}
}

func startCommandHandler(cm *utils.CallJobManager) func(c *discordgo.Session, e *discordgo.InteractionCreate) error {
	return func(c *discordgo.Session, e *discordgo.InteractionCreate) error {
		startOpt := &utils.CallJob{
			GuildId:  e.GuildID,
			ExceptCh: nil,
		}

		for _, opt := range e.ApplicationCommandData().Options {
			if opt.Name == "user" && opt.Type == discordgo.ApplicationCommandOptionUser {
				startOpt.User = opt.UserValue(c)
			}
			if opt.Name == "times" && opt.Type == discordgo.ApplicationCommandOptionInteger {
				startOpt.Times = int(opt.IntValue())
			}
			if opt.Name == "except" && opt.Type == discordgo.ApplicationCommandOptionChannel {
				startOpt.ExceptCh = &opt.ChannelValue(c).ID
			}
		}

		if startOpt.User == nil {
			return fmt.Errorf("failed to get user command option")
		}

		cm.Start(startOpt)

		c.InteractionResponseEdit(e.Interaction, utils.BasicResponseEdit("Ok!"))
		return nil
	}
}
