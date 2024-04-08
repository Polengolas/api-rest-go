package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Polengolas/api-rest-go/database"
	"github.com/Polengolas/api-rest-go/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	var p []models.Personality
	database.DB.Find(&p)

	json.NewEncoder(w).Encode(p)
}

func ReturnOnePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p []models.Personality

	database.DB.Find(&p, id)

	json.NewEncoder(w).Encode(p)
}

func CreateNewPersonality(w http.ResponseWriter, r *http.Request) {
	var newPersonality models.Personality

	json.NewDecoder(r.Body).Decode(&newPersonality)

	database.DB.Create(&newPersonality)

	json.NewEncoder(w).Encode(newPersonality)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personality

	database.DB.Delete(&p, id)

	json.NewEncoder(w).Encode(p)
}

func EditPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personality

	database.DB.First(&p, id)

	json.NewDecoder(r.Body).Decode(&p)

	database.DB.Save(&p)

	json.NewEncoder(w).Encode(p)
}
