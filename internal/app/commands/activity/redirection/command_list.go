package redirection

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ActivityRedirectionCommander) List(msg *tgbotapi.Message) {
	redirection, err := c.redirectionService.List(1, RecorcPerMessage)
	if err != nil {
		log.Println(err)
		ans := tgbotapi.NewMessage(msg.Chat.ID, "Could not list redirections")
		c.bot.Send(ans)
		return
	}

	if len(redirection) == 0 {
		reply := tgbotapi.NewMessage(msg.Chat.ID, "Redirections is empty")
		c.bot.Send(reply)
		return
	}

	text := GetListRedirection(redirection)

	var prev, next CallbackData
	prevBytes, _ := json.Marshal(&prev)
	next.Offset = RecorcPerMessage + 1
	nextBytes, _ := json.Marshal(&next)

	ans := tgbotapi.NewMessage(msg.Chat.ID, text)

	keyboard := GetKeyboard(prevBytes, nextBytes)

	ans.ReplyMarkup = &keyboard

	c.bot.Send(ans)
}
