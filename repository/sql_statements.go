package repository

var (
	getPlayersStatement   = "SELECT * FROM players;"
	getPlayerStatement    = "SELECT * FROM players WHERE id=?;"
	insertPlayerStatement = "INSERT INTO players (id, nickname, superpower) VALUES (?, ?, ?);"
	deletePlayerStatement = "DELETE FROM players WHERE id=?;"
)
