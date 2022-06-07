package actions

import (
	"github.com/go-irc/irc"
	"github.com/mpgelliston/ircbot/bot"
)

func SexyAction(c *irc.Client, m *irc.Message, b *bot.Bot) {
	if m.Command == "PRIVMSG" && b.Admins[m.Prefix.Name] {
		c.WriteMessage(&irc.Message{
			Command: "PRIVMSG",
			Params: []string{
				b.Channel,
				"You're looking kinda sexy, " + m.Prefix.Name,
			},
		})
	}
}
