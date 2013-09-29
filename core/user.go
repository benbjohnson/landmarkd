package core

type User struct {
	Id     string                 `json:"id"`
	Traits map[string]interface{} `json:"traits"`
}
