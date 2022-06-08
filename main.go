package main

import (
	"fmt"
	"log"

	"github.com/mpgelliston/ircbot/actions"
	"github.com/mpgelliston/ircbot/bot"
)

func main() {
	fmt.Println("Configuring IRC Bot...")
	options := bot.BotOptions{
		Nick:     "NorthBot",
		User:     "NorthBot",
		Name:     "NorthBot",
		Password: "MAGA2024",
		Server:   "irc.libera.chat",
		Port:     6667,
		Channel:  "#testnorthbot",
		Admins:   map[string]bool{"matt1982": true, "lux0r": true, "mrbalihai": true},
		Debug:    true,
		Verbose:  true,
	}

	fmt.Println("Starting IRC Bot")

	b, err := bot.NewBot(options)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Attaching Actions")
	b.AddAction(bot.BotAction{
		Name:   "Join",
		Action: actions.JoinAction,
	})
	b.AddAction(bot.BotAction{
		Name:   "Welcome",
		Action: actions.WelcomeAction,
	})
	b.AddAction(bot.BotAction{
		Name:   "Sexy",
		Action: actions.SexyAction,
	})
	b.AddAction(bot.BotAction{
		Name:   "List Actions",
		Action: actions.ListActionsAction,
	})
	b.AddAction(bot.BotAction{
		Name:   "Uptime",
		Action: actions.UptimeAction,
	})
	b.AddAction(bot.BotAction{
		Name:   "Crypto",
		Action: actions.CryptoPriceAction,
	})

	// Create the client
	b.Connect()
}
