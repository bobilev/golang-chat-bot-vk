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
	"time"
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
	bot.GetById = bot.GetGroupID()
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
func (bot *BotVkApiGroup) GetGroupID() int {
	method := "groups.getById"
	urlConfig := bot.constructURL(method)

	jsonGetById := ResponseGetById{}
	CallMethod(urlConfig,&jsonGetById)

	return jsonGetById.Response[0].Id
}
func (bot *BotVkApiGroup) InitLongPollServer(LPC *LongPollConfig) {
	method := "groups.getLongPollServer"
	group_id := "group_id=" + strconv.Itoa(bot.GetById)
	urlConfig := bot.constructURL(method,group_id)

	jsonGetLongPollServer := ResponseGetLongPollServer{}
	CallMethod(urlConfig,&jsonGetLongPollServer)

	LPC.Key = jsonGetLongPollServer.Response.Key
	LPC.Server = jsonGetLongPollServer.Response.Server
	LPC.Ts = jsonGetLongPollServer.Response.Ts
	LPC.Wait = 25
}
func (bot *BotVkApiGroup) StartLongPollServer() (chan ObjectUpdate) {
	LPC := new(LongPollConfig)
	bot.InitLongPollServer(LPC)
	ch := make(chan ObjectUpdate, 1)

	go func(ch chan ObjectUpdate){
		for {
			log.Println("New request: TS",LPC.Ts)
			updateLP := new(UpdateLP)

			connectLPCurl := LPC.Server+"?act=a_check&key="+LPC.Key+"&ts="+strconv.Itoa(LPC.Ts)+"&wait="+strconv.Itoa(LPC.Wait)
			err := CallMethod(connectLPCurl,&updateLP)
			if err != nil {
				log.Println("[ERR]CallMethod Reconnect 3 sec\n",err)
				time.Sleep(time.Second * 3)

				continue
			}
			if updateLP.Failed == 1 || updateLP.Ts == "" {
				fmt.Println("Failed:",updateLP.Failed)
				fmt.Println("Ts:",updateLP.Ts)
				continue
			} else if updateLP.Failed == 2 || updateLP.Failed == 3 {
				fmt.Println("Failed:",updateLP.Failed)
				bot.InitLongPollServer(LPC)
				continue
			}
			fmt.Println("end switch")
			LPC.Ts , _ = strconv.Atoi(updateLP.Ts)

			for _, update := range updateLP.Updates {
				fmt.Println("ch <- update",update)
				ch <- update
			}
		}
	}(ch)

	fmt.Println(LPC.Key)
	fmt.Println(LPC.Server)
	fmt.Println(LPC.Ts)
	fmt.Println(LPC.Wait)
	return ch
}
func CallMethod(url string, result interface{}) error {
	fmt.Println("START[CallMethod]",url)
	resultReq , err := Call(url)
	if err != nil {
		return err
	}
	jsonRes := []byte(resultReq)
	json.Unmarshal(jsonRes,result)
	return nil
}
func Call(urlString string) (string, error) {
	res, err := http.Get(urlString)
	defer res.Body.Close()

	if err != nil {
		return "",err
	}
	resultReq, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "",err
	}

	resultString := fmt.Sprintf("%s", resultReq)
	log.Println("{Call}",resultString)
	return resultString, nil
}
