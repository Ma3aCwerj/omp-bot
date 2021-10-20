package redirection

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/activity/redirection"
)

type RedirectionCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)
	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)
}

type ActivityRedirectionCommander struct {
	bot                *tgbotapi.BotAPI
	redirectionService *redirection.DummyActivityRedirectionService
}

func NewRedirectionCommander(bot *tgbotapi.BotAPI) *ActivityRedirectionCommander {

	service := redirection.NewDummyActivityRedirectionService()

	return &ActivityRedirectionCommander{bot: bot, redirectionService: service}
}

func (c *ActivityRedirectionCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("ActivityRedirectionCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *ActivityRedirectionCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	case "delete":
		c.Delete(msg)
	case "get":
		c.Get(msg)
	default:
		c.Default(msg)
	}
}
