package models

type ErrResponse struct {
	Error string `json:"msg"`
}

type MsgResponse struct {
	Msg string `json:"msg"`
}
