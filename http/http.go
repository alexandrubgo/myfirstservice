package http

import (
	"encoding/json"
	"fmt"
	"mysql/model"
	"mysql/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserHttp struct {
	userService services.Service
}

func NewUserHttp(userService services.Service) *UserHttp {
	return &UserHttp{userService: userService}
}

func (u *UserHttp) GetPlayers(w http.ResponseWriter, r *http.Request) {
	players, err := u.userService.GetPlayers()
	if err != nil {
		http.Error(w, "Heheh", http.StatusBadRequest)
		return
	}

	_ = json.NewEncoder(w).Encode(players)
}

func (u *UserHttp) GetPlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		http.Error(w, "no id provided", http.StatusBadRequest)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "cannot caca", http.StatusInternalServerError)
		return
	}

	players, err := u.userService.GetPlayer(uint64(idInt))
	if err != nil {
		http.Error(w, "Heheh", http.StatusBadRequest)
		return
	}
	_ = json.NewEncoder(w).Encode(players)
}

func (u *UserHttp) CreatePlayer(w http.ResponseWriter, r *http.Request) {

	fmt.Println("create player")

	p := &model.Player{}
	if err := json.NewDecoder(r.Body).Decode(p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(p)
	// vars := mux.Vars(r)
	// id := vars["id"]
	// nickname := vars["nickname"]
	// superpower := vars["superpower"]

	// if id == "" || nickname == "" || superpower == "" {
	// 	http.Error(w, "arguments missing", http.StatusBadRequest)
	// 	return
	// }

	// idInt, err := strconv.Atoi(id)
	// if err != nil {
	// 	http.Error(w, "internal server errror", http.StatusInternalServerError)
	// 	return
	// }

	players, err := u.userService.CreatePlayer(p.ID, p.Nickname, p.Superpower)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_ = json.NewEncoder(w).Encode(players)
}

func (u *UserHttp) DeletePlayer(w http.ResponseWriter, r *http.Request) {

}
