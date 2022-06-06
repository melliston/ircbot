package bot

type BotOptions struct {
	Nick     string
	Channel  string
	Password string
	Ident    string
	Server   string
	Port     int
}

type Bot struct {
	Nick     string
	Channel  string
	Password string
	Ident    string
	Server   string
	Port     int
}

func NewBot(opt BotOptions) (*Bot, error) {
	bot := &Bot{
		Nick:     opt.Nick,
		Channel:  opt.Channel,
		Password: opt.Password,
		Ident:    opt.Ident,
		Server:   opt.Server,
		Port:     opt.Port,
	}
	return bot, nil
}
