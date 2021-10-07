package model

type Player struct {
	ID         int    `json:"id"`
	Nickname   string `json:"nickname"`
	Superpower string `json:"superpower"`
}
