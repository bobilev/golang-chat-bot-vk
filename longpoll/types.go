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
