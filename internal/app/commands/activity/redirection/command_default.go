package redirection

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ActivityRedirectionCommander) Default(msg *tgbotapi.Message) {

	log.Printf("%s %s", msg.From.UserName, msg.Text)

	text := fmt.Sprintf("Unknown command. %s", msg.Command())
	text += "\n\nUse /help__activity__redirection\n"
	ans := tgbotapi.NewMessage(msg.Chat.ID, text)
	c.bot.Send(ans)
}
