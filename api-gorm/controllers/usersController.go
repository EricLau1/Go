package controllers

import (
	"io/ioutil"
	"encoding/json"
	"net/http"
	"strconv"
	"api-gorm/utils"
	"api-gorm/models"
	"github.com/gorilla/mux"
)

func PostUser(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var user models.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		utils.ToError(w, err, http.StatusUnprocessableEntity)
		return
	}
	rows, err := models.NewUser(user)
	if err != nil {
		utils.ToError(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, rows, http.StatusCreated)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := models.GetAll(models.USERS)
	utils.ToJson(w, users, http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := models.GetById(models.USERS, vars["id"])
	utils.ToJson(w, user, http.StatusOK)
}

func PutUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 32)
	var user models.User
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		utils.ToError(w, err, http.StatusUnprocessableEntity)
		return
	}
	user.Id = uint32(id)
	rows, err := models.UpdateUser(user)
	if err != nil {
		utils.ToError(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, rows, http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := models.Delete(models.USERS, vars["id"])
	if err != nil {
		utils.ToError(w, err, http.StatusUnprocessableEntity)
		return	
	}
	w.WriteHeader(http.StatusNoContent)
}
