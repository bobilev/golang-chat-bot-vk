package vkchatbot

import (
	"net/url"
	"strconv"
)

type LongPollConfig struct {
	Key string
	Server string
	Ts int
	Wait int
}
func (LPC *LongPollConfig) ConstructURL() url.URL{
	Url := url.URL{
		Host:     LPC.Server,
	}
	q := Url.Query()
	q.Set("act", "a_check")
	q.Add("key", LPC.Key)
	q.Add("ts", strconv.Itoa(LPC.Ts))
	q.Add("wait", strconv.Itoa(LPC.Wait))

	Url.RawQuery = q.Encode()

	return Url
}
type ResponseGetById struct {
	Response []GetById `json:"response"`

}
type GetById struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type ResponseGetLongPollServer struct {
	Response GetLongPollServer
}
type GetLongPollServer struct {
	Key string
	Server string
	Ts int
}
type UpdateLP struct {
	Ts string
	Updates []ObjectUpdate
	Failed int
}
type ObjectUpdate struct {
	Type string
	Update `json:"object"`
	GroupId int `json:"group_id"`
}
type Update struct {
	Id int
	Date int
	Out int
	UserId int `json:"user_id"`
	ReadState int `json:"read_state"`
	Title string
	Body string
}
//type UpdatesChannel <-chan ObjectUpdate
type ResSendMessage struct {
	MessageID  int  `json:"response"`//идентификатор сообщения;
}
type Attachment struct {
	MediaId int
	TypeDoc string
	OwnerId int
}