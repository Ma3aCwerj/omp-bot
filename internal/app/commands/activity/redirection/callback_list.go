package redirection

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackData struct {
	Offset uint64 `json:"offset"`
}

func (c *ActivityRedirectionCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	callbackAnswer := tgbotapi.NewCallback(callback.ID, "")
	c.bot.AnswerCallbackQuery(callbackAnswer)

	var data CallbackData
	json.Unmarshal([]byte(callbackPath.CallbackData), &data)

	redirection, err := c.redirectionService.List(data.Offset, RecorcPerMessage)
	if err != nil {
		log.Println(err)
		return
	}

	if len(redirection) == 0 {
		c.bot.Send(tgbotapi.NewEditMessageText(callback.Message.Chat.ID, callback.Message.MessageID, "Redirections is empty"))
		return
	}

	text := GetListRedirection(redirection)

	ans := tgbotapi.NewEditMessageText(callback.Message.Chat.ID, callback.Message.MessageID, text)

	var prev, next CallbackData

	if data.Offset <= RecorcPerMessage {
		prev.Offset = 0
	} else {
		prev.Offset = data.Offset - RecorcPerMessage
	}
	prevBytes, _ := json.Marshal(&prev)

	next.Offset = data.Offset + RecorcPerMessage
	nextBytes, _ := json.Marshal(&next)

	keyboard := GetKeyboard(prevBytes, nextBytes)

	ans.ReplyMarkup = &keyboard

	c.bot.Send(ans)
}
