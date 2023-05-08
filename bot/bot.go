package bot

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"github.com/delapaska/tgbot/animals"
	"github.com/delapaska/tgbot/cryptho"
	"github.com/delapaska/tgbot/keyboard"
	"github.com/delapaska/tgbot/weather"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const KELVIN = 272.1

func StartBot() {

	const apiKey = "5928241038:AAFMNVHpYADdw1Y11uEQehZFw7eArlCpEwc"

	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		log.Panic(err)
		return
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updChannel := bot.GetUpdatesChan(u)

	if err != nil {
		log.Panic(err)
	}

	var update tgbotapi.Update

	for {

		update = <-updChannel

		if update.Message != nil {
			/*
				if update.Message.IsCommand() {
					cmdText := update.Message.Command()
					if cmdText == "start" {

						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, это бот, написанный для отбора на стажировку в VK\nОн умеет:\n*Отправлять картинки собак и котов\n*Играть в камень-ножницы-бумага\n*Показывать погоду в Москве и Санкт-Петербурге\n*Показывать курс BTC и ETH ")
						msg.ReplyMarkup = keyboard.CreateMainMenu()
						bot.Send(msg)
					}

				}
			*/
			if update.Message.Command() == "start" {

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, это бот, написанный для отбора на стажировку в VK\nОн умеет:\n*Отправлять картинки собак и котов\n*Играть в камень-ножницы-бумага\n*Показывать погоду в Москве и Санкт-Петербурге\n*Показывать курс BTC и ETH ")
				msg.ReplyMarkup = keyboard.CreateMainMenu()
				bot.Send(msg)

			} else if update.Message.Text == keyboard.CreateMainMenu().Keyboard[0][0].Text {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выбери животное:")
				msg.ReplyMarkup = keyboard.CreateKeyboardAnimalPhoto()
				bot.Send(msg)

				for {
					update = <-updChannel
					if update.Message.Text == keyboard.CreateKeyboardAnimalPhoto().Keyboard[0][2].Text {

						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Возвращаемся")
						msg.ReplyMarkup = keyboard.CreateMainMenu()
						bot.Send(msg)
						break

					} else if update.Message.Text == keyboard.CreateKeyboardAnimalPhoto().Keyboard[0][0].Text {

						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Случайная Собака")
						str := animals.GetRandomDogImageUrl()
						photo := tgbotapi.NewPhoto(update.Message.From.ID, tgbotapi.FileURL(str))
						bot.Send(photo)
						bot.Send(msg)

					} else if update.Message.Text == keyboard.CreateKeyboardAnimalPhoto().Keyboard[0][1].Text {

						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Случайная кошка")
						status := animals.GetRandomCatStatusCode()
						st := strconv.Itoa(status)
						str := "https://http.cat/" + st
						//	fmt.Println(status)
						photo := tgbotapi.NewPhoto(update.Message.From.ID, tgbotapi.FileURL(str))
						bot.Send(photo)
						bot.Send(msg)

					}
				}
			} else if update.Message.Text == keyboard.CreateMainMenu().Keyboard[0][1].Text {

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Это игра в камень-ножницы-бумага:")
				msg.ReplyMarkup = keyboard.CreateKeyboardRockPaperScissors()
				bot.Send(msg)
				scorePlayer := 0
				scoreEnemy := 0

				for {
					update = <-updChannel
					if update.Message.Text == keyboard.CreateKeyboardRockPaperScissors().Keyboard[0][3].Text {

						scorePlayer = 0
						scoreEnemy = 0
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Возвращаемся")
						msg.ReplyMarkup = keyboard.CreateMainMenu()
						bot.Send(msg)
						break

					}
					if update.Message.Text == keyboard.CreateKeyboardRockPaperScissors().Keyboard[0][0].Text {

						randomIndex := rand.Intn(3)
						pick := keyboard.CreateKeyboardRockPaperScissors().Keyboard[0][randomIndex].Text
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, pick)
						bot.Send(msg)

						if randomIndex == 0 {
						} else if randomIndex == 1 {
							scoreEnemy++
						} else if randomIndex == 2 {
							scorePlayer++
						}

						msg1 := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("счёт: %d : %d", scorePlayer, scoreEnemy))
						bot.Send(msg1)

					} else if update.Message.Text == keyboard.CreateKeyboardRockPaperScissors().Keyboard[0][1].Text {

						randomIndex := rand.Intn(3)
						pick := keyboard.CreateKeyboardRockPaperScissors().Keyboard[0][randomIndex].Text
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, pick)
						bot.Send(msg)

						if randomIndex == 1 {
						} else if randomIndex == 2 {
							scoreEnemy++
						} else if randomIndex == 0 {
							scorePlayer++
						}

						msg1 := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("счёт: %d : %d", scorePlayer, scoreEnemy))
						bot.Send(msg1)

					} else if update.Message.Text == keyboard.CreateKeyboardRockPaperScissors().Keyboard[0][2].Text {

						randomIndex := rand.Intn(3)
						pick := keyboard.CreateKeyboardRockPaperScissors().Keyboard[0][randomIndex].Text
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, pick)
						bot.Send(msg)

						if randomIndex == 2 {
							fmt.Println("Draw")
						} else if randomIndex == 0 {
							scoreEnemy++
						} else if randomIndex == 1 {
							scorePlayer++
						}

						msg1 := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("счёт: %d : %d", scorePlayer, scoreEnemy))
						bot.Send(msg1)

					}

					if scorePlayer == 3 {

						msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Ты победил(а), счёт обнулился"))
						bot.Send(msg)
						scorePlayer = 0
						scoreEnemy = 0

					} else if scoreEnemy == 3 {

						msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Ты проиграл(а), счёт обнулился:"))
						bot.Send(msg)
						scorePlayer = 0
						scoreEnemy = 0

					}
				}
			} else if update.Message.Text == keyboard.CreateMainMenu().Keyboard[0][2].Text {

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Погода Москва-Санкт-Петербург сейчас:")
				msg.ReplyMarkup = keyboard.CreateKeyboardWeather()
				bot.Send(msg)

				for {
					update = <-updChannel
					if update.Message.Text == keyboard.CreateKeyboardWeather().Keyboard[0][2].Text {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Возвращаемся")
						msg.ReplyMarkup = keyboard.CreateMainMenu()
						bot.Send(msg)
						break
					} else if update.Message.Text == keyboard.CreateKeyboardWeather().Keyboard[0][0].Text {
						Temp := weather.TakeTempMoscow() - KELVIN
						Wind := weather.TakeWindSpeedMoscow()
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Температура в Москве: %.2f C\nСкорость Ветра в Москве: %.2f м/c", Temp, Wind))
						bot.Send(msg)

					} else if update.Message.Text == keyboard.CreateKeyboardWeather().Keyboard[0][1].Text {
						Temp := weather.TakeTempSPB() - KELVIN
						Wind := weather.TakeWindSpeedSPB()
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Температура в Санкт-Петербурге: %.2f C\nСкорость Ветра в Санкт-Петербурге: %.2f м/c", Temp, Wind))
						bot.Send(msg)

					}

				}

			} else if update.Message.Text == keyboard.CreateMainMenu().Keyboard[0][3].Text {

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "BTC и ETH")
				msg.ReplyMarkup = keyboard.CreateKeyboardCryptho()
				bot.Send(msg)
				for {
					update = <-updChannel
					if update.Message.Text == keyboard.CreateKeyboardCryptho().Keyboard[0][2].Text {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Возвращаемся")
						msg.ReplyMarkup = keyboard.CreateMainMenu()
						bot.Send(msg)
						break
					} else if update.Message.Text == keyboard.CreateKeyboardCryptho().Keyboard[0][0].Text {
						price, _ := cryptho.GetPrice("BTC")
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Курс BTC = %.2f$", price))
						bot.Send(msg)

					} else if update.Message.Text == keyboard.CreateKeyboardCryptho().Keyboard[0][1].Text {
						price, _ := cryptho.GetPrice("ETH")
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Курс ETH = %.2f$", price))
						bot.Send(msg)

					}
				}
			}

		}
	}
}
