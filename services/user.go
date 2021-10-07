package services

import (
	"fmt"
	"mysql/model"
	"mysql/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (u *UserService) GetPlayers() ([]*model.Player, error) {
	fmt.Println("Get Players Method")

	players, err := u.userRepo.GetPlayers()
	if err != nil {
		fmt.Println("u.userRepo.GetPlayers", err)
		return nil, err
	}

	return players, nil
}

func (u *UserService) GetPlayer(id uint64) (*model.Player, error) {
	fmt.Println("Get Player Method")

	player, err := u.userRepo.GetPlayer(uint64(id))
	if err != nil {
		fmt.Println("u.userRepo.GetPlayer", err)
		return nil, err
	}
	return player, nil
}

func (u *UserService) CreatePlayer(id int, nickname, superpower string) (*model.Player, error) {
	fmt.Println("Create Player Method")
	player := &model.Player{
		ID:         id,
		Nickname:   nickname,
		Superpower: superpower,
	}

	playerDB, err := u.userRepo.CreatePlayer(player)
	if err != nil {
		fmt.Println("u.userRepo.CreatePlayer", err)
		return nil, err
	}

	return playerDB, nil
}

func (u *UserService) DeletePlayer(id int) {
	fmt.Println("Delete Player Method")
	u.userRepo.DeletePlayer(uint64(id))
}
