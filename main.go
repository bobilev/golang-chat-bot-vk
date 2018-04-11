package main

import (

	"github.com/bobilev/golang-chat-bot-vk/longpoll"
	"fmt"
)

func main() {
	accessToken := "b25e0478970ebcde8977b7c7b9b8562e28cce81c9f80518b0fa72196fdc0588d833ff6f298a821d12ba18"

	bot := longpoll.InitBot(accessToken)

	updates := bot.StartLongPollServer()

	for update := range updates {
		//if update.Body == "" {
		//	continue
		//}
		if update.Body == "hi" {
			res , _ := bot.SendMessage(update.UserId,"Hello")
			fmt.Println("[res]",res.MessageID)
		}
		if update.Body == "sex" {
			res , _ := bot.SendPhoto(update.UserId,456239017,"секси эльфийка")
			fmt.Println("[res]",res.MessageID)
		}
		fmt.Println(update)
		fmt.Println("Text Message:",update.Body)
	}
	////////2 - указывать уровень отображения инфы
	////////3 - цикл update
}

