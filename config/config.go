package config

const (
	URLScheme = "https"
	HOST = "api.vk.com"
	PATH = "method/groups."
)

type ResponseGetById struct {
	Response []GetById `json:"response"`

}
type GetById struct {
	Id int `json:"id"`
	Name string `json:"name"`
}