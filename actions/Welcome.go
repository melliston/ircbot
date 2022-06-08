package actions

import (
	"fmt"

	"github.com/go-irc/irc"
	"github.com/mpgelliston/ircbot/bot"
)

const (
	WELCOME_TRIGGER string = "JOIN"
)

func WelcomeAction(c *irc.Client, m *irc.Message, b *bot.Bot) {
	if m.Command == WELCOME_TRIGGER {
		c.WriteMessage(&irc.Message{
			Command: "PRIVMSG",
			Params: []string{
				b.Channel,
				fmt.Sprintf("Welcome to %s, %s. We hope you enjoy your time here. ", b.Channel, m.Prefix.Name),
			},
		})
	}
}
