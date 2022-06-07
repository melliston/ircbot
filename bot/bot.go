package bot

import (
	"fmt"
	"log"
	"net"

	"github.com/go-irc/irc"
)

// BotOptions A struct used to configure a new instance of a Bot
type BotOptions struct {
	Nick     string
	Channel  string
	Password string
	Ident    string
	User     string
	Name     string
	Server   string
	Port     int
	Verbose  bool
	Debug    bool
	Admins   map[string]bool
}

type BotAction struct {
	Name   string
	Action func(*irc.Client, *irc.Message, *Bot)
}

type Bot struct {
	Nick     string
	Channel  string
	Password string
	Ident    string
	User     string
	Name     string
	Server   string
	Port     int
	Client   *irc.Client
	Verbose  bool
	Debug    bool
	Admins   map[string]bool
	Actions  []BotAction
}

func NewBot(opt BotOptions) (*Bot, error) {
	bot := &Bot{
		Nick:     opt.Nick,
		User:     opt.User,
		Name:     opt.Name,
		Channel:  opt.Channel,
		Password: opt.Password,
		Ident:    opt.Ident,
		Server:   opt.Server,
		Port:     opt.Port,
		Verbose:  opt.Verbose,
		Debug:    opt.Debug,
		Admins:   opt.Admins,
	}
	return bot, nil
}

func (b *Bot) Connect() {
	if b.Verbose {
		fmt.Printf("Connecting to Server: %s:%d ", b.Server, b.Port)
	}

	conn, err := net.Dial("tcp", "irc.libera.chat:6667")
	if err != nil {
		log.Fatalln(err)
	}

	if b.Verbose {
		fmt.Println("IRC Bot Connectting...")
	}
	config := irc.ClientConfig{
		Nick:    b.Nick,
		Pass:    b.Password,
		User:    b.User,
		Name:    b.Name,
		Handler: irc.HandlerFunc(b.Handler),
	}

	// Initiate the new client
	client := irc.NewClient(conn, config)

	if b.Verbose {
		fmt.Println("IRC Config Initiated")
	}

	err = client.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

func (b *Bot) AddAction(a BotAction) {
	b.Actions = append(b.Actions, a)
}

func (b *Bot) Handler(c *irc.Client, m *irc.Message) {
	if b.Debug {
		fmt.Println(m.Command)
		fmt.Println(m.User)
		fmt.Println(m.Prefix)
		fmt.Println(m.Params)
	}

	for _, action := range b.Actions {
		action.Action(c, m, b)
	}
	// if m.Command == "PRIVMSG" && c.FromChannel(m) {
	// 	// Create a handler on all messages.
	// 	c.WriteMessage(&irc.Message{
	// 		Command: "PRIVMSG",
	// 		Params: []string{
	// 			m.Params[0],
	// 			m.Trailing(),
	// 		},
	// 	})
	// }
}
