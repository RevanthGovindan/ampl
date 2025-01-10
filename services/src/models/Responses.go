package models

type ErrResponse struct {
	Error string `json:"msg"`
}

type MsgResponse struct {
	Msg string `json:"msg"`
}

type LoginResponse struct {
	Name  string `json:"name"`
	Token string `json:"token"`
	Type  string `json:"type"`
}
