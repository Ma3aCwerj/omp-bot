package redirection

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ActivityRedirectionCommander) Edit(msg *tgbotapi.Message) {
	redirection, err := GetArguments(msg.CommandArguments())

	if err != nil {
		log.Println(err)
		c.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, err.Error()))
		return
	}

	err = c.redirectionService.Update(redirection.Id, redirection)
	if err != nil {
		log.Println(err)
		c.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, err.Error()))
		return
	}

	text := fmt.Sprintf("Update redirection with id %d, tittle %s", redirection.Id, string(redirection.Title))
	ans := tgbotapi.NewMessage(msg.Chat.ID, text)
	c.bot.Send(ans)
}
