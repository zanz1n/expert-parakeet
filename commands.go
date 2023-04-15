package main

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/zanz1n/bot-inocente/commands"
)

var (
	cmds = make(map[string]*commands.Command)
)

func GetCommand(name string) (*commands.Command, bool) {
	cmd, ok := cmds[name]
	return cmd, ok
}

func AddCommand(cmd *commands.Command) {
	cmds[cmd.Data.Name] = cmd
}

func PostCommands(c *discordgo.Session) {
	guild := os.Getenv("GUILD_ID")
	i, ec := 0, 0

	for _, cmd := range cmds {
		_, err := c.ApplicationCommandCreate(c.State.User.ID, guild, cmd.Data)
		if err != nil {
			log.Printf("failed to post command: %s", err.Error())
			ec++
			continue
		}
		i++
	}

	log.Printf("%b commands posted on %s, %b failed", i, guild, ec)
}
