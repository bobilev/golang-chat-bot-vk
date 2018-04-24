package vkchatbot

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"net/url"
	"encoding/json"
	"strings"
	"strconv"
	"time"
	"log"
)
type BotVkApiGroup struct {
	AccessToken string
	GetById int
	Log int
	Url url.URL
}
func InitBot(accessToken string) *BotVkApiGroup {
	bot := new(BotVkApiGroup)
	bot.AccessToken = accessToken
	bot.Url = url.URL{
		Host:     "https://api.vk.com",
		Path:     "/method/",
	}
	bot.GetById = bot.GetGroupID()
	return bot
}
func (bot BotVkApiGroup) constructURL(method string,params ...string) url.URL {
	urlConfig := bot.Url
	urlConfig.Path += method

	q := urlConfig.Query()
	q.Set("access_token", bot.AccessToken)
	q.Add("v", "5.74")
	for _,val := range params {
		values := strings.Split(val,"=")
		q.Add(values[0],values[1])
	}
	urlConfig.RawQuery = q.Encode()
	//fmt.Println("++++++",q)
	//fmt.Println("------",urlConfig.Query())
	return urlConfig
}
func (bot *BotVkApiGroup) GetGroupID() int {
	method := "groups.getById"
	urlConfig := bot.constructURL(method)

	jsonGetById := ResponseGetById{}
	bot.CallMethod(urlConfig,&jsonGetById)

	return jsonGetById.Response[0].Id
}
func (bot *BotVkApiGroup) InitLongPollServer(LPC *LongPollConfig) {
	method := "groups.getLongPollServer"
	groupId := "group_id=" + strconv.Itoa(bot.GetById)
	urlConfig := bot.constructURL(method,groupId)

	jsonGetLongPollServer := ResponseGetLongPollServer{}
	bot.CallMethod(urlConfig,&jsonGetLongPollServer)

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
			updateLP := new(UpdateLP)

			err := bot.CallMethod(LPC.ConstructURL(),&updateLP)
			if err != nil {
				log.Println("[ERR]CallMethod Reconnect 3 sec\n",err)
				time.Sleep(time.Second * 3)
				continue
			}
			if updateLP.Failed == 1 || updateLP.Ts == "" {
				continue
			} else if updateLP.Failed == 2 || updateLP.Failed == 3 {
				if bot.Log == 1 {
					log.Println("[Failed:",updateLP.Failed,"] ReInitLongPollServer")
				}
				bot.InitLongPollServer(LPC)
				continue
			}
			LPC.Ts , _ = strconv.Atoi(updateLP.Ts)

			for _, update := range updateLP.Updates {
				ch <- update
			}
		}
	}(ch)

	return ch
}
func (bot BotVkApiGroup) CallMethod(url url.URL, result interface{}) error {
	resultReq , err := Call(url)
	switch bot.Log {
	case 1:
		log.Println("[Respons]",resultReq)
	case 2:
		log.Println("[Request]",url)
		log.Println("[Respons]",resultReq)
	}
	if err != nil {
		return err
	}
	jsonRes := []byte(resultReq)
	json.Unmarshal(jsonRes,result)
	return nil
}
func Call(url url.URL) (string, error) {
	fmt.Println(url)
	fmt.Println(url.Query())
	//res, err := http.Get(urlString)
	urlString := url.Host+url.Path
	fmt.Println("urlString",urlString)
	res, err := http.PostForm(urlString,url.Query())
	fmt.Println(res)
	//res.Header.Set()

	if err != nil {
		return "",err
	}
	defer res.Body.Close()
	resultReq, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "",err
	}

	resultString := strings.Replace(fmt.Sprintf("%s", resultReq),"\n","",-1)
	return resultString, nil
}
