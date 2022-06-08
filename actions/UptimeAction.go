package actions

import (
	"strings"

	"github.com/go-irc/irc"
	"github.com/mpgelliston/ircbot/bot"
)

const (
	UPTIME_TRIGGER string = "PRIVMSG"
	UPTIME_PHRASE  string = "UPTIME"
)

func UptimeAction(c *irc.Client, m *irc.Message, b *bot.Bot) {
	if m.Command == "PRIVMSG" && b.Admins[m.Prefix.Name] && strings.ToUpper(m.Trailing()) == UPTIME_PHRASE {
		c.WriteMessage(&irc.Message{
			Command: "PRIVMSG",
			Params: []string{
				b.Channel,
				b.Uptime(),
			},
		})
	}
}
