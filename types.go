package vkchatbot

import (
	"net/url"
	"strconv"
)

type LongPollConfig struct {
	Key    string
	Server string
	Ts     string
	Wait   int
}
func (LPC *LongPollConfig) ConstructURL() url.URL{
	Url := url.URL{
		Host:     LPC.Server,
	}
	q := Url.Query()
	q.Set("act", "a_check")
	q.Add("key", LPC.Key)
	q.Add("ts", LPC.Ts)
	q.Add("wait", strconv.Itoa(LPC.Wait))

	Url.RawQuery = q.Encode()

	return Url
}
type ResponseGetById struct {
	Response []GetById `json:"response"`
}
type GetById struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ResponseGetLongPollServer struct {
	Response GetLongPollServer
}
type GetLongPollServer struct {
	Key    string
	Server string
	Ts     string
}
type UpdateLP struct {
	Ts      string
	Updates []ObjectUpdate
	Failed  int
}
type ObjectUpdate struct {
	Type string
	Object struct {
		Date      int    `json:"date"`
		FromId    int    `json:"from_id"`
		Id        int    `json:"id"`
		Out       int    `json:"out"`
		PeerId    int    `json:"peer_id"`
		Text      string `json:"text"`
		//fwd_messages	[]
		Important bool   `json:"important"`
		RandomId  int    `json:"random_id"`
		//attachments	[]
		IsHidden  bool   `json:"is_hidden"`
		ConversationMessageId int `json:"conversation_message_id"`
	} `json:"object"`
	GroupId int `json:"group_id"`
}
//type UpdatesChannel <-chan ObjectUpdate
type ResSendMessage struct {
	MessageID  int  `json:"response"`//идентификатор сообщения;
}
type Attachment struct {
	MediaId   int
	TypeDoc   string
	OwnerId   int
	AccessKey string
}