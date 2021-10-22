package redirection

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ActivityRedirectionCommander) Delete(msg *tgbotapi.Message) {

	id, err := strconv.ParseUint(msg.CommandArguments(), 10, 64)

	if msg.CommandArguments() == "" || err != nil {
		log.Printf("\n\nwrong argument %d, error: %s\n\n", id, err)
		ans := tgbotapi.NewMessage(msg.Chat.ID, "Wrong arguments or not a valid id")
		c.bot.Send(ans)
		return
	}

	ok, err := c.redirectionService.Remove(id)
	text := "Redirection with id: " + strconv.Itoa(int(id)) + " deleted"

	if !ok {
		text = err.Error()
	}
	c.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, text))
}
