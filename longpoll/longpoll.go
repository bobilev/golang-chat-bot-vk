package longpoll

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
	"net/url"
	"github.com/bobilev/golang-chat-bot-vk/config"
	"encoding/json"
	"strings"
	"strconv"
)
type BotVkApiGroup struct {
	Access_token string
	GetById int
	Url url.URL
}
func InitBot(access_token string) *BotVkApiGroup {
	bot := new(BotVkApiGroup)
	bot.Access_token = access_token

	bot.Url = url.URL{
		Scheme:   config.URLScheme,
		Host:     config.HOST,
		Path:     config.PATH,
	}

	bot.GetById = bot.GetGroupID(access_token)

	return bot
}
func (bot BotVkApiGroup) constructURL(method string,params ...string) string {
	urlConfig := bot.Url
	urlConfig.Path += method

	q := urlConfig.Query()
	q.Set("access_token", bot.Access_token)
	q.Add("v", "5.74")
	for _,val := range params {
		values := strings.Split(val,"=")
		q.Add(values[0],values[1])
	}
	urlConfig.RawQuery = q.Encode()
	return urlConfig.String()
}
func (bot *BotVkApiGroup) GetGroupID(access_token string) int {
	method := "getById"
	urlConfig := bot.constructURL(method)

	jsonGetById := ResponseGetById{}
	CallMethod(urlConfig,&jsonGetById)

	return jsonGetById.Response[0].Id
}
func (bot *BotVkApiGroup) InitLongPollServer(LPC *LongPollConfig) {
	method := "getLongPollServer"
	group_id := "group_id=" + strconv.Itoa(bot.GetById)
	urlConfig := bot.constructURL(method,group_id)

	jsonGetLongPollServer := ResponseGetLongPollServer{}
	CallMethod(urlConfig,&jsonGetLongPollServer)

	LPC.Key = jsonGetLongPollServer.Response.Key
	LPC.Server = jsonGetLongPollServer.Response.Server
	LPC.Ts = jsonGetLongPollServer.Response.Ts
	LPC.Wait = 25
}
func (bot *BotVkApiGroup) StartLongPollServer() {
	LPC := new(LongPollConfig)
	bot.InitLongPollServer(LPC)

	//fmt.Println(LPC.Key)
	//fmt.Println(LPC.Server)
	//fmt.Println(LPC.Ts)
	//fmt.Println(LPC.Wait)
	//LPC.Ts = 2
	//fmt.Println(LPC.Ts)


}
func CallMethod(url string, result interface{}) {
	resultReq := Call(url)

	jsonRes := []byte(resultReq)
	json.Unmarshal(jsonRes,result)
}
func Call(urlString string,) string {
	res, err := http.Get(urlString)
	defer res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}
	resultReq, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	resultString := fmt.Sprintf("%s", resultReq)
	return resultString
}
