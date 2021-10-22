package redirection

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ActivityRedirectionCommander) Help(msg *tgbotapi.Message) {

	text := "This is help to domain - activity and subdomain - redirection\n\nList supports commands:\n" +
		"/help__activity__redirection - this message, list supports commands\n" +
		"/list__activity__redirection - list activity redirection\n" +
		"/get__activity__redirection - get activity redirection by id.\n\tuse\n\t/get__activity__redirection 3\n" +
		"/new__activity__redirection - add new activity redirection. Required fields: title, description, from, to.\n" +
		"\tuse\n\t/new__activity__redirection {\"title\": \"Title\", \"description\": \"Description\", \"from\": \"from_redirected\", \"to\": \"to_redirect\"}\n\n" +
		"/edit__activity__redirection - edit activity redirection. Required fields: id, title, description, from, to\n" +
		"\tuse\n\t/edit__activity__redirection {\"id\": 3, title: \"New title\", \"description\": \"Description\", \"from\": \"from_redirected\", \"to\": \"to_redirect\"}\n\n" +
		"/delete__activity__redirection - delete an exist activity redirection by id.\n\tuse\n\t/delete__activity__redirection 3\n"

	c.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, text))
}
