package repository

import "mysql/model"

type UserRepository interface {
	GetPlayers() ([]*model.Player, error)
	GetPlayer(id uint64) (*model.Player, error)
	CreatePlayer(player *model.Player) (*model.Player, error)
	DeletePlayer(id uint64)
}
