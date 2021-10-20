package activity

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/activity/redirection"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type ActivityCommander struct {
	bot                  *tgbotapi.BotAPI
	redirectionCommander Commander
}

func NewActivityCommander(bot *tgbotapi.BotAPI) *ActivityCommander {
	return &ActivityCommander{bot: bot, redirectionCommander: redirection.NewRedirectionCommander(bot)}
}

func (c *ActivityCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "redirection":
		c.redirectionCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("ActivityCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *ActivityCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "redirection":
		c.redirectionCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("ActivityCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
