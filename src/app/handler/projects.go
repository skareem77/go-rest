package handler

import (
	"encoding/json"
	"net/http"
	"rest-api/src/app/model"

	"github.com/jinzhu/gorm"
)

// GetAllProjects get all projects
func GetAllProjects(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	projects := []model.Project{}
	db.Find(&projects)
	ResponseJSON(w, http.StatusOK, projects)
}

// SaveProject creats a new project in DB
func SaveProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	project := model.Project{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&project)

	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	if err = db.Save(&project).Error; err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJSON(w, http.StatusCreated, project)
}
