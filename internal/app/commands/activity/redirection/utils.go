package redirection

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/activity"
)

const RecorcPerMessage = 5

func GetArguments(args string) (activity.Redirection, error) {
	var redirection activity.Redirection

	if err := json.Unmarshal([]byte(args), &redirection); err != nil {
		log.Println(err)
		return redirection, errors.New("wrong json arguments")
	}

	return redirection, nil
}

func GetKeyboard(prev, next []byte) tgbotapi.InlineKeyboardMarkup {

	prevCallback := path.CallbackPath{
		Domain:       "activity",
		Subdomain:    "redirection",
		CallbackName: "list",
		CallbackData: string(prev),
	}

	nextCallback := path.CallbackPath{
		Domain:       "activity",
		Subdomain:    "redirection",
		CallbackName: "list",
		CallbackData: string(next),
	}

	return tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("<", prevCallback.String()),
		tgbotapi.NewInlineKeyboardButtonData(">", nextCallback.String())))
}

func GetListRedirection(list []activity.Redirection) string {
	text := "List redirections:\n\n"
	for _, s := range list {
		text += fmt.Sprintf("%d - %s\nFrom - %s\tTo - %s\n", s.Id, s.Title, s.From, s.To)
	}
	return text
}
