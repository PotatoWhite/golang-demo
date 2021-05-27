package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var user User
	records := DB.First(&user, vars["id"])
	if records.RowsAffected == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	DB.Delete(&User{}, vars["id"])
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var user User
	DB.First(&user, vars["id"])
	json.NewDecoder(r.Body).Decode(&user)

	w.Header().Set("Content-Type", "application/json")
	if id, _ := strconv.ParseUint(vars["id"], 10,32); uint(id) != user.ID {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid Param")
		return
	}

	DB.Save(&user)
	json.NewEncoder(w).Encode(&user)
}

func CreateUsers(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	DB.Create(&user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	DB.Find(&users)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&users)
}

