package main

import (
	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var numericKeyboard = tgBotApi.NewReplyKeyboard(
	tgBotApi.NewKeyboardButtonRow(
		tgBotApi.NewKeyboardButton("New reminder"),
	),
	tgBotApi.NewKeyboardButtonRow(
		tgBotApi.NewKeyboardButton("My reminders"),
	),
	tgBotApi.NewKeyboardButtonRow(
		tgBotApi.NewKeyboardButton("My profile"),
	),
)

func main() {

	//db := config.Open()
	bot, err := tgBotApi.NewBotAPI("5801953953:AAEkU4a_UrCAp4AqUfMXv-Yp7y9IUf8fN2c")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgBotApi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil { // If we got a message

			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgBotApi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			msg.ReplyMarkup = numericKeyboard
			//msg.ReplyMarkup = tgBotApi.NewRemoveKeyboard(true)

			bot.Send(msg)
		}
	}
}
