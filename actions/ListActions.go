package actions

import (
	"fmt"

	"github.com/go-irc/irc"
	"github.com/mpgelliston/ircbot/bot"
)

// ListActionsAction - Lists each action that is loaded on the bot.
func ListActionsAction(c *irc.Client, m *irc.Message, b *bot.Bot) {
	if m.Command == "PRIVMSG" && b.Admins[m.Prefix.Name] && m.Trailing() == "LIST ACTIONS" {
		c.WriteMessage(&irc.Message{
			Command: "PRIVMSG",
			Params: []string{
				m.Prefix.Name,
				"The Following Actions are Currently Enabled:",
			},
		})
		for _, action := range b.Actions {
			c.WriteMessage(&irc.Message{
				Command: "PRIVMSG",
				Params: []string{
					m.Prefix.Name,
					fmt.Sprintf("•%s", action.Name),
				},
			})
		}
	}
}
