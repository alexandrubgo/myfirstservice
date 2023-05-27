// main.go
package main

import (
	"database/sql"
	"log"
	geluHttp "mysql/delivery/http"
	"mysql/repository"
	"mysql/services"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	//PRIMUL FEATURE ADAUGAT
	db, err := sql.Open("mysql", "root:123qweasdzxc@tcp(127.0.0.1:3306)/database")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	userMysqlRepo := repository.NewUserMysqlRepository(db)
	userServices := services.NewUserService(userMysqlRepo)
	userHttp := geluHttp.NewUserHttp(userServices)

	r := mux.NewRouter()

	// GET
	r.HandleFunc("/players", userHttp.GetPlayers).Methods("GET")
	// GET
	r.HandleFunc("/players/{id}", userHttp.GetPlayer).Methods("GET")
	// POST
	r.HandleFunc("/playersx", userHttp.CreatePlayer).Methods("POST")
	//DELETE
	r.HandleFunc("/players/remove/{id}", userHttp.DeletePlayer).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
