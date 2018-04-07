package main

import (

	"github.com/bobilev/golang-chat-bot-vk/longpoll"
	"fmt"
)

func main() {
	//method := "getLongPollServer"
	access_token := "b25e0478970ebcde8977b7c7b9b8562e28cce81c9f80518b0fa72196fdc0588d833ff6f298a821d12ba18"
	//url := &url.URL{
	//	Scheme:   "https",
	//	Host:     "api.vk.com",
	//	Path:     "method/groups.",
	//}
	//
	//url.Path += method
	//q := url.Query()
	//q.Set("access_token", access_token)
	//q.Add("v", "5.74")
	//q.Add("group_id", "164670950")
	//url.RawQuery = q.Encode()
	//fmt.Println(url)

	////////1 - указывать токен
	bot := longpoll.InitBot(access_token)
	fmt.Println(bot.GetById)

	////////2 - указывать уровень отображения инфы
	////////3 - цикл update
}

