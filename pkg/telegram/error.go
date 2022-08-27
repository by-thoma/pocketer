package telegram

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	errInvalidURL   = errors.New("url is invalid")
	errUnauthorized = errors.New("user is not authorized")
	errUnableToSave = errors.New("unable to save")
)

//msg.Text = "Ты не авторизирован"
//msg.Text = "Это невалидная ссылка"
//msg.Text = "Не удалость сохранить ссылку"
func (b *Bot) handleError(chatID int64, err error) {
	msg := tgbotapi.NewMessage(chatID, "Произошла неизвестная ошибка")

	switch err {
	case errInvalidURL:
		msg.Text = "Это невалидная ссылка"
		b.bot.Send(msg)
	case errUnauthorized:
		msg.Text = "Ты не авторизирован"
		b.bot.Send(msg)
	case errUnableToSave:
		msg.Text = "Не удалость сохранить ссылку"
		b.bot.Send(msg)
	default:
		b.bot.Send(msg)

	}

}
