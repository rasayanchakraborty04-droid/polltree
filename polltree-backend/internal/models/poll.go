package models

type Poll struct {
	ID int `json:"id"`
	Question string `json:"question"`
	Options []Option `json:"Options"`
}

type Option struct {
	ID int `json:"id"`
	Text string `json:"text"`
}

type CreatePollRequest struct {
	Question string `json:"question"`
	Options []string `json:"options"`
}