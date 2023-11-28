package domain

type LoginResponse struct {
	Id          int32  `json:"id"`
	AccessToken string `json:"access_token"`
}
