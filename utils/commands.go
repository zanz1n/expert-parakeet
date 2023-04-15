package utils

import "github.com/zanz1n/bot-inocente/commands"

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
