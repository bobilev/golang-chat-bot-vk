package longpoll


type LongPollConfig struct {
	Key string
	Server string
	Ts int
	Wait int
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
	MessageID  int  `json:"кesponse"`//идентификатор сообщения;
}