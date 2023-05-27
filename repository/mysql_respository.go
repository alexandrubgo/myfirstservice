package repository

import (
	"database/sql"
	"fmt"
	"mysql/model"
)

type UserMysqlRepository struct {
	db *sql.DB
}

func NewUserMysqlRepository(db *sql.DB) *UserMysqlRepository {
	return &UserMysqlRepository{
		db: db,
	}
}

func (u *UserMysqlRepository) GetPlayers() ([]*model.Player, error) {
	rows, err := u.db.Query(getPlayersStatement)
	if err != nil {
		fmt.Println("u.db.Query", err)
		return nil, err
	}

	players := []*model.Player{}

	for rows.Next() {
		tempUser := &model.Player{}

		if err := rows.Scan(
			&tempUser.ID,
			&tempUser.Nickname,
			&tempUser.Superpower,
		); err != nil {
			fmt.Println("rows.Scan", err)
			return nil, err
		}

		players = append(players, tempUser)
	}

	return players, nil
}

func (u *UserMysqlRepository) GetPlayer(id uint64) (*model.Player, error) {
	user := &model.Player{}

	if err := u.db.QueryRow(getPlayerStatement, id).Scan(
		&user.ID,
		&user.Nickname,
		&user.Superpower,
	); err != nil {
		fmt.Println("u.db.QueryRow", err)
		return nil, err
	}

	return user, nil
}

func (u *UserMysqlRepository) CreatePlayer(player *model.Player) (*model.Player, error) {
	if err := u.db.QueryRow(insertPlayerStatement,
		player.ID,
		player.Nickname,
		player.Superpower,
	).Err(); err != nil {
		fmt.Println("u.db.QueryRow", err)
		return nil, err
	}

	return player, nil
}

func (u *UserMysqlRepository) DeletePlayer(id uint64) {
	u.db.QueryRow(deletePlayerStatement, id)
}
