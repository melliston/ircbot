package actions

import (
	"github.com/go-irc/irc"
	"github.com/mpgelliston/ircbot/bot"
)

// Join - Joins the inital channel the bot should reside in.
func JoinAction(c *irc.Client, m *irc.Message, b *bot.Bot) {
	if m.Command == "001" {
		// 001 is a welcome event, so we join channels there
		// c.Write("IDENTIFY NorthBot MAGA2024")
		c.Write("JOIN " + b.Channel)
	}
}
