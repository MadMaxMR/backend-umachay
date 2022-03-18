package main

import (
	"API/source/api"
	"net/http"

	"github.com/gorilla/mux"

	"fmt"
)

func main() {
	var port string = "8000"

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/").Subrouter()
	//Rutas de la API de cursos
	apiRouter.HandleFunc("/cursos", api.GetAllCursos).Methods("GET")
	apiRouter.HandleFunc("/cursos/create", api.CreateCursos).Methods("POST")
	apiRouter.HandleFunc("/cursos/{id}", api.GetCursos).Methods("GET")
	apiRouter.HandleFunc("/cursos/update/{id}", api.UpdateCursos).Methods("PUT")
	apiRouter.HandleFunc("/cursos/delete/{id}", api.DeleteCursos).Methods("DELETE")
	//Rutas de la API de temas
	apiRouter.HandleFunc("/temas", api.GetAllTemas).Methods("GET")
	apiRouter.HandleFunc("/temas/create", api.CreateTema).Methods("POST")
	apiRouter.HandleFunc("/temas/{id}", api.GetTema).Methods("GET")
	apiRouter.HandleFunc("/temas/update/{id}", api.UpdateTema).Methods("PUT")
	apiRouter.HandleFunc("/temas/delete/{id}", api.DeleteTema).Methods("DELETE")
	apiRouter.HandleFunc("/temas/{id}/cursos", api.GetAllTemasCursos).Methods("GET")
	//Rutas de la API de usuarios

	fmt.Printf("Servidor corriendo en el puerto: %s", port)
	http.ListenAndServe(":"+port, router)
}
