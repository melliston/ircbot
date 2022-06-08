package actions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-irc/irc"
	"github.com/mpgelliston/ircbot/bot"
)

const (
	CRYPTO_TRIGGER string = "crypto"
	CRYPTO_HELP    string = "help"
	CRYPTO_BUY     string = "buy"
	CRYPTO_SELL    string = "sell"
)

type CryptoErrorResponse struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

type CryptoCoinbaseResponse struct {
	Data struct {
		Base     string `json:"base"`
		Currency string `json:"currency"`
		Amount   string `json:"amount"`
	} `json:"data"`
	Errors []CryptoErrorResponse `json:"errors"`
}

func CryptoPriceAction(c *irc.Client, m *irc.Message, b *bot.Bot) {
	var messages []string

	commands := strings.Split(strings.ToLower(m.Trailing()), " ")
	if len(commands) == 2 && commands[0] == CRYPTO_TRIGGER {
		if commands[1] == CRYPTO_HELP {
			messages = append(messages, "CRYPTO HELP")
			messages = append(messages, "The following commands are available:")
			messages = append(messages, "• CRYPTO HELP - (This command)")
			messages = append(messages, "• CRYPTO BUY <From-To> - Buy Price for a crypto pair eg CRYPTO BUY USD-BTC")
			messages = append(messages, "• CRYPTO SELL <From-To> - Sell Price for a crypto pair eg CRYPTO SELL BTC-USD")
		}
	} else if len(commands) == 3 {
		// Handle the BUY and SELL prices
		action := commands[1]
		pair := commands[2]
		if action == CRYPTO_BUY || action == CRYPTO_SELL {
			url := fmt.Sprintf("https://api.coinbase.com/v2/prices/%s/%s", pair, action) // Use the free Coinbase account

			httpClient := http.Client{
				Timeout: time.Second * 2,
			}

			req, err := http.NewRequest(http.MethodGet, url, nil)
			if err != nil {
				log.Fatal(err)
				messages = append(messages, err.Error())
			}

			res, getErr := httpClient.Do(req)
			if getErr != nil {
				messages = append(messages, getErr.Error())
			}

			if res.Body != nil {
				defer res.Body.Close()
			}

			body, readErr := ioutil.ReadAll(res.Body)
			if readErr != nil {
				messages = append(messages, readErr.Error())
			}

			cryptoPrice := CryptoCoinbaseResponse{}
			jsonErr := json.Unmarshal(body, &cryptoPrice)
			if jsonErr != nil {
				messages = append(messages, jsonErr.Error())
			}
			if len(cryptoPrice.Errors) > 0 {
				messages = append(messages, "There was an error processing your response:")
				for _, err := range cryptoPrice.Errors {
					messages = append(messages, fmt.Sprintf("•%s.", err.Message))
				}

			} else {
				messages = append(messages, fmt.Sprintf("%s: %s - %s", action, pair, cryptoPrice.Data.Amount))
			}
		}
	}
	if len(messages) > 0 {
		for _, message := range messages {
			c.WriteMessage(&irc.Message{
				Command: "PRIVMSG",
				Params: []string{
					b.Channel,
					message,
				},
			})
		}
	}

}
