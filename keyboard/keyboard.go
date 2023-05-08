package keyboard

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CreateKeyboardAnimalPhoto() tgbotapi.ReplyKeyboardMarkup {
	var keyboardAnimalPhoto = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("૮₍ • ᴥ • ₎ა"),
			tgbotapi.NewKeyboardButton("ฅ^._.^ฅ"),
			tgbotapi.NewKeyboardButton("🔙"),
		),
	)
	return keyboardAnimalPhoto
}

func CreateKeyboardRockPaperScissors() tgbotapi.ReplyKeyboardMarkup {
	var keyboardRockPaperScissors = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🪨"),
			tgbotapi.NewKeyboardButton("⬜"),
			tgbotapi.NewKeyboardButton("✂️"),
			tgbotapi.NewKeyboardButton("🔙"),
		),
	)
	return keyboardRockPaperScissors

}

func CreateKeyboardWeather() tgbotapi.ReplyKeyboardMarkup {
	var keyboardWeather = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Москва"),
			tgbotapi.NewKeyboardButton("Санкт-Петербург"),
			tgbotapi.NewKeyboardButton("🔙"),
		),
	)
	return keyboardWeather
}
func CreateKeyboardCryptho() tgbotapi.ReplyKeyboardMarkup {
	var keyboardWeather = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("BTC ₿"),
			tgbotapi.NewKeyboardButton("ETH ♦️"),
			tgbotapi.NewKeyboardButton("🔙"),
		),
	)
	return keyboardWeather
}
func CreateMainMenu() tgbotapi.ReplyKeyboardMarkup {
	var mainMenu = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🐶Получить фото🐱"),
			tgbotapi.NewKeyboardButton("Камень-Ножницы-Бумага"),
			tgbotapi.NewKeyboardButton("Погода СПБ, Москва"),
			tgbotapi.NewKeyboardButton("Курс Криптовалют"),
		),
	)
	return mainMenu
}
