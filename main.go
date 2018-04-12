package main

import (

	"github.com/bobilev/golang-chat-bot-vk/longpoll"
	"fmt"
)

func main() {
	accessToken := "b25e0478970ebcde8977b7c7b9b8562e28cce81c9f80518b0fa72196fdc0588d833ff6f298a821d12ba18"

	bot := longpoll.InitBot(accessToken)
	bot.Log = 2 // 0,1,2 - уровни отображения логов
	updates := bot.StartLongPollServer()

	for update := range updates {
		fmt.Println("UserID:",update.UserId,"Text Message:",update.Body)
		if update.Body == "hi" {
			res , _ := bot.SendMessage(update.UserId,"Hello")
			fmt.Println("[res]",res.MessageID)
		}
		if update.Body == "sex" {
			res , _ := bot.SendDoc(update.UserId,"photo",456239017,"секси эльфийка")
			fmt.Println("[res]",res.MessageID)
		}

	}
}