package keyboard

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func CreateMainKeyboard() tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Кнопка 1"),
			tgbotapi.NewKeyboardButton("Кнопка 2"),
			tgbotapi.NewKeyboardButton("Кнопка 3"),
			tgbotapi.NewKeyboardButton("Кнопка 4"),
		),
	)

	return keyboard
}
