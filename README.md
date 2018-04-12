# golang-chat-bot-vk

```go
package main

import (
    "github.com/bobilev/golang-chat-bot-vk"
	"fmt"
)

func main() {
	accessToken := "VkBotToken"

	bot := vkchatbot.InitBot(accessToken)
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
```