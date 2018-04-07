package longpoll

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
	"net/url"
	"github.com/bobilev/golang-chat-bot-vk/config"
	"encoding/json"
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
		Scheme:   config.URLScheme,
		Host:     config.HOST,
		Path:     config.PATH,
	}

	url.Path += method
	q := url.Query()
	q.Set("access_token", access_token)
	q.Add("v", "5.74")
	url.RawQuery = q.Encode()


	jsonGetById := config.ResponseGetById{}
	Call(url.String(),&jsonGetById)


	fmt.Println(jsonGetById.Response)
	fmt.Println(jsonGetById)

	return "No"
}
func Call(urlString string,result interface{}) string {
	res, err := http.Get(urlString)
	defer res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}
	resultReq, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	//jsonGetById := config.ResponseGetById{}
	jsonRes := []byte(resultReq)
	json.Unmarshal(jsonRes,result)

	//fmt.Println(jsonGetById.Response[0].Name)
	//fmt.Println(jsonGetById.Response[0].Id)
	//fmt.Println(result)

	resultString := fmt.Sprintf("%s", resultReq)
	return resultString
}
