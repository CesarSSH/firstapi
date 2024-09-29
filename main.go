package main

import (
	"net/http" // Paquete para manejar peticiones HTTP

	"github.com/CesarSSH/go-gorm-api/db"
	"github.com/CesarSSH/go-gorm-api/models"
	"github.com/CesarSSH/go-gorm-api/routes" //Consumo un paquete llamado routes
	"github.com/gorilla/mux"                 // Paquete para manejar rutas de manera más flexible
)

func main() {
	//Ejecuto esta funcion para conectar a la DB
	db.DBConnection()

	//Ejecuto las migraciones de las tablas
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	// Crea un nuevo enrutador que va a manejar las rutas (URLs)
	router := mux.NewRouter()

	// Define lo que debe pasar cuando alguien visite la ruta principal "/" y las rutas configuradas
	router.HandleFunc("/", routes.HomeHandler)
	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/user", routes.PostUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	//Tasks routes
	router.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	router.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	router.HandleFunc("/task", routes.PostTaskHandler).Methods("POST")
	router.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")


	// Inicia el servidor web en el puerto 3000 y usa el enrutador que definimos
	http.ListenAndServe(":3000", router)
}
