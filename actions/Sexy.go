package actions

import (
	"github.com/go-irc/irc"
	"github.com/mpgelliston/ircbot/bot"
)

const (
	SEXY_TRIGGER string = "PRIVMSG"
)

func SexyAction(c *irc.Client, m *irc.Message, b *bot.Bot) {
	phrase := "Hi " + b.Nick
	if m.Command == SEXY_TRIGGER && m.Trailing() == phrase {
		c.WriteMessage(&irc.Message{
			Command: "PRIVMSG",
			Params: []string{
				b.Channel,
				"You're looking kinda sexy, " + m.Prefix.Name,
			},
		})
	}
}
