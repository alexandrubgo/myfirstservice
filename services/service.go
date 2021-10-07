package services

import "mysql/model"

type Service interface {
	GetPlayers() ([]*model.Player, error)
	GetPlayer(id uint64) (*model.Player, error)
	CreatePlayer(id int, nickname, superpower string) (*model.Player, error)
	DeletePlayer(id int)
}
