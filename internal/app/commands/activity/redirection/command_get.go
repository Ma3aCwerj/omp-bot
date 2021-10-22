package redirection

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ActivityRedirectionCommander) Get(msg *tgbotapi.Message) {

	id, err := strconv.ParseUint(msg.CommandArguments(), 10, 64)

	if msg.CommandArguments() == "" || err != nil {
		log.Printf("\n\nwrong argument %d, error: %s\n\n", id, err)
		ans := tgbotapi.NewMessage(msg.Chat.ID, "Wrong arguments or not a valid id")
		c.bot.Send(ans)
		return
	}

	redirection, err := c.redirectionService.Describe(id)
	if err != nil {
		log.Printf("\n\n%s\n\n", err)
		ans := tgbotapi.NewMessage(msg.Chat.ID, "Could not get entity")
		c.bot.Send(ans)
		return
	}

	ans := tgbotapi.NewMessage(msg.Chat.ID, redirection.String())
	c.bot.Send(ans)
}
