package redirection

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ActivityRedirectionCommander) New(msg *tgbotapi.Message) {

	redirection, err := GetArguments(msg.CommandArguments())

	if err != nil {
		log.Println(err)
		ans := tgbotapi.NewMessage(msg.Chat.ID, err.Error())
		c.bot.Send(ans)
		return
	}

	id, err := c.redirectionService.Create(redirection)
	if err != nil {
		log.Println(err)
		ans := tgbotapi.NewMessage(msg.Chat.ID, "Could not create redirection")
		c.bot.Send(ans)
		return
	}

	text := fmt.Sprintf("Created redirection with id %d, tittle %s", id, string(redirection.Title))
	ans := tgbotapi.NewMessage(msg.Chat.ID, text)
	c.bot.Send(ans)
}
