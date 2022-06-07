package main

import (
	"fmt"
	"log"

	"github.com/mpgelliston/ircbot/action"
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
	a := bot.BotAction{
		Name:   "Join",
		Action: action.Join,
	}
	b.AddAction(a)

	// Create the client
	b.Connect()
}
