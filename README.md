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
	bot.Log = 2 // 1,2,3 - уровни отображения логов
	updates := bot.StartLongPollServer()

	for update := range updates {
		fmt.Println("UserID:",update.Object.FromId,"Text Message:",update.Object.Text)
		if update.Object.Text == "hi" {
			res , _ := bot.SendMessage(update.Object.FromId,"Hello")
			fmt.Println("[res]",res.MessageID)
		}
		if update.Object.Text == "text" {
			var Attach vkchatbot.Attachment
			Attach.TypeDoc = "photo"       //Тип документа photo,video,audio,doc,wall,market
			Attach.MediaId = 123456789     // id документа
			Attach.OwnerId = 123456789     // id группы владельца документа
			Attach.AccessKey = "AccessKey" // AccessKey документа (пригодится если группа закрытая)

			res , _ := bot.SendDoc(update.Object.FromId,Attach,"эльфийка")
			fmt.Println("[res]",res.MessageID)
		}
	}
}
```