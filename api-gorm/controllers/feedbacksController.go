package controllers

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"api-gorm/utils"
	"api-gorm/models"
)

func PostFeedback(w http.ResponseWriter, r *http.Request) {
	var feedback models.Feedback
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &feedback)
	if err != nil {
		utils.ToError(w, err, http.StatusUnprocessableEntity)
		return
	}
	rows, err := models.NewFeedback(feedback)
	if err != nil {
		utils.ToError(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, rows, http.StatusCreated)	
}

func GetFeedbacks(w http.ResponseWriter, r *http.Request) {
	feedbacks := models.GetAll(models.FEEDBACKS)
	utils.ToJson(w, feedbacks, http.StatusOK)
}