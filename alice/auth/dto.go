package auth

type yandexAuthInfoDTO struct {
	Id           string   `json:"id"`
	Login        string   `json:"login"`
	ClientId     string   `json:"client_id"`
	DefaultEmail string   `json:"default_email"`
	Emails       []string `json:"emails"`
	DefaultPhone struct {
		Id     int    `json:"id"`
		Number string `json:"number"`
	} `json:"default_phone"`
	Avatar string `json:"default_avatar_id"`
	PSuid  string `json:"psuid"`
}
