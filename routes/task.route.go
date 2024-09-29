package routes

import (
	"encoding/json"
	"github.com/CesarSSH/go-gorm-api/db"
	"github.com/CesarSSH/go-gorm-api/models"
	"github.com/gorilla/mux"
	"net/http"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)
	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}
	json.NewEncoder(w).Encode(task)
}
func PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	// Validar si el userID existe
	var user models.User
	db.DB.First(&user, task.UserID)
	if user.ID == 0 {
		w.WriteHeader(http.StatusBadRequest) // Error 400
		w.Write([]byte("UserID does not exist"))
		return
	}
	// Crea la tarea
	createdTask := db.DB.Create(&task)
	err := createdTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //Error 400
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&task)

}
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)
	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}
	// db.DB.Delete(&task) //Esta funcion me desactivar la task
	db.DB.Unscoped().Delete(&task)      //Esta funcion si me borra el task
	w.WriteHeader(http.StatusNoContent) //2024

}
