package longpoll

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
	"net/url"
)
type BotVkApiGroup struct {
	Access_token string
	GetById string
}
func InitBot(access_token string) *BotVkApiGroup {
	bot := new(BotVkApiGroup)
	bot.Access_token = access_token
	bot.GetById = GetGroupID(access_token)
	return bot
}
func GetGroupID(access_token string) string {
	method := "getById"
	url := &url.URL{
		Scheme:   "https",
		Host:     "api.vk.com",
		Path:     "method/groups.",
	}

	url.Path += method
	q := url.Query()
	q.Set("access_token", access_token)
	q.Add("v", "5.74")
	url.RawQuery = q.Encode()

	fmt.Println(Call(url.String()))

	return "No"
}
func Call(urlString string) string {
	res, err := http.Get(urlString)
	defer res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	resultString := fmt.Sprintf("%s", result)
	return resultString
}