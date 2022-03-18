package api

import (
	"API/source/helpers"
	"API/source/models"
	"encoding/json"

	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Data struct {
	Success bool            `json:"success"`
	Data    []models.Cursos `json:"data"`
	Errors  []string        `json:"errors"`
}

type Data1 struct {
	Success bool           `json:"success"`
	Data    []models.Temas `json:"data"`
	Errors  []string       `json:"errors"`
}

type Data2 struct {
	Success bool            `json:"success"`
	Data    []models.Videos `json:"data"`
	Errors  []string        `json:"errors"`
}

//API Cursos
func CreateCursos(w http.ResponseWriter, req *http.Request) {
	bodyCursos, success := helpers.DecodeBody(req)

	if success != true {
		http.Error(w, "Error al leer el body", http.StatusBadRequest)
		return
	}

	var data Data = Data{Errors: make([]string, 0)}
	bodyCursos.Description = strings.TrimSpace(bodyCursos.Description)
	bodyCursos.Nombre = strings.TrimSpace(bodyCursos.Nombre)
	bodyCursos.Image = strings.TrimSpace(bodyCursos.Image)
	if !helpers.IsValid(bodyCursos.Description, bodyCursos.Nombre) {
		data.Success = false
		data.Errors = append(data.Errors, "invalid data")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(json)
		return

	}
	//insertamos los cursos si son validos
	cursos, success := models.Insert(bodyCursos.Nombre, bodyCursos.Description, bodyCursos.Image)
	if success != true {
		data.Errors = append(data.Errors, "error al insertar cursos")
	}
	//si todo sale bien
	data.Success = true
	//dentro del dato llenamos con el  curso creado
	data.Data = append(data.Data, cursos)
	//creación del Json con los datos dentro
	json, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(json)
	return
}

func GetCursos(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id_curso := vars["id"]

	var data Data

	var cursos models.Cursos
	var success bool

	cursos, success = models.Get(id_curso)
	if success != true {
		data.Success = false
		data.Errors = append(data.Errors, "Curso no encontrado")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(json)
		return
	}

	data.Success = true
	data.Data = append(data.Data, cursos)

	json, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func GetAllCursos(w http.ResponseWriter, req *http.Request) {
	var cursos []models.Cursos = models.GetAll()

	var data = Data{true, cursos, make([]string, 0)}
	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func DeleteCursos(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id_curso := vars["id"]

	var data = Data{Errors: make([]string, 0)}

	cursos, success := models.Delete(id_curso)

	if success != true {
		data.Errors = append(data.Errors, "Curso no eliminado")
	}

	data.Success = true
	data.Data = append(data.Data, cursos)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func UpdateCursos(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id_curso := vars["id"]

	bodyCursos, success := helpers.DecodeBody(req)

	if success != true {
		http.Error(w, "Error al leer el body", http.StatusBadRequest)
		return
	}

	var data Data = Data{Errors: make([]string, 0)}

	bodyCursos.Description = strings.TrimSpace(bodyCursos.Description)
	bodyCursos.Nombre = strings.TrimSpace(bodyCursos.Nombre)
	bodyCursos.Image = strings.TrimSpace(bodyCursos.Image)

	if !helpers.IsValid(bodyCursos.Description, bodyCursos.Nombre) {
		data.Success = false
		data.Errors = append(data.Errors, "invalid data")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(json)
		return

	}

	cursos, success := models.Update(id_curso, bodyCursos.Nombre, bodyCursos.Description, bodyCursos.Image)

	if success != true {
		data.Errors = append(data.Errors, "Curso no actualizado")
	}

	data.Success = true
	data.Data = append(data.Data, cursos)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(json)
	return
}

///API Temas

func CreateTema(w http.ResponseWriter, req *http.Request) {
	bodyTemas, success := helpers.DecodeBodyTemas(req)

	if success != true {
		http.Error(w, "Error al leer el body", http.StatusBadRequest)
		return
	}

	var data Data1 = Data1{Errors: make([]string, 0)}

	bodyTemas.Description = strings.TrimSpace(bodyTemas.Description)
	bodyTemas.Title = strings.TrimSpace(bodyTemas.Title)

	if !helpers.IsValidTemas(bodyTemas.Description, bodyTemas.Title, bodyTemas.Curso_Id) {
		data.Success = false
		data.Errors = append(data.Errors, "invalid data")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(json)
		return

	}
	//insertamos los cursos si son validos
	temas, success := models.InsertTema(bodyTemas.Title, bodyTemas.Description, bodyTemas.Curso_Id)
	if success != true {
		data.Errors = append(data.Errors, "error al insertar tema")
	}
	//si todo sale bien
	data.Success = true
	//dentro del dato llenamos con el  curso creado
	data.Data = append(data.Data, temas)
	//creación del Json con los datos dentro
	json, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(json)
	return
}

func GetTema(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id_tema := vars["id"]

	var data Data1

	var temas models.Temas
	var success bool

	temas, success = models.GetTema(id_tema)
	if success != true {
		data.Success = false
		data.Errors = append(data.Errors, "Tema no encontrado")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(json)
		return
	}

	data.Success = true
	data.Data = append(data.Data, temas)

	json, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func GetAllTemas(w http.ResponseWriter, req *http.Request) {
	var temas []models.Temas = models.GetAllTemas()

	var data = Data1{true, temas, make([]string, 0)}
	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func DeleteTema(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id_tema := vars["id"]

	var data = Data1{Errors: make([]string, 0)}

	cursos, success := models.DeleteTemas(id_tema)

	if success != true {
		data.Errors = append(data.Errors, "Tema no eliminado")
	}

	data.Success = true
	data.Data = append(data.Data, cursos)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func UpdateTema(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id_tema := vars["id"]

	bodyTemas, success := helpers.DecodeBodyTemas(req)

	if success != true {
		http.Error(w, "Error al leer el body", http.StatusBadRequest)
		return
	}

	var data Data1 = Data1{Errors: make([]string, 0)}

	bodyTemas.Description = strings.TrimSpace(bodyTemas.Description)
	bodyTemas.Title = strings.TrimSpace(bodyTemas.Title)

	if !helpers.IsValidTemas(bodyTemas.Description, bodyTemas.Title, bodyTemas.Curso_Id) {
		data.Success = false
		data.Errors = append(data.Errors, "invalid data")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(json)
		return

	}

	temas, success := models.UpdateTema(id_tema, bodyTemas.Title, bodyTemas.Curso_Id, bodyTemas.Description)

	if success != true {
		data.Errors = append(data.Errors, "Tema no actualizado")
	}

	data.Success = true
	data.Data = append(data.Data, temas)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(json)
	return
}

func GetAllTemasCursos(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id_curso := vars["id"]

	var temas []models.Temas = models.GetAllTemasCursos(id_curso)
	var data = Data1{true, temas, make([]string, 0)}
	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

//API Videos

func CreateVideo(w http.ResponseWriter, req *http.Request) {
	bodyVideos, success := helpers.DecodeBodyVideos(req)

	if success != true {
		http.Error(w, "Error al leer el body", http.StatusBadRequest)
		return
	}

	var data Data2 = Data2{Errors: make([]string, 0)}

	bodyVideos.Url_Video = strings.TrimSpace(bodyVideos.Url_Video)

	if !helpers.IsValidVideos(bodyVideos.Url_Video, bodyVideos.Nivel, bodyVideos.Id_Tema_Id) {
		data.Success = false
		data.Errors = append(data.Errors, "invalid data")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(json)
		return
	}
	//insertamos los videos si son validos
	videos, success := models.InsertVideo(bodyVideos.Url_Video, bodyVideos.Nivel, bodyVideos.Id_Tema_Id)
	if success != true {
		data.Errors = append(data.Errors, "error al insertar tema")
	}
	//si todo sale bien
	data.Success = true
	//dentro del dato llenamos con el  video creado
	data.Data = append(data.Data, videos)
	//creación del Json con los datos dentro
	json, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(json)
	return
}

func GetVideo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id_tema := vars["id"]

	var data Data1

	var temas models.Temas
	var success bool

	temas, success = models.GetTema(id_tema)
	if success != true {
		data.Success = false
		data.Errors = append(data.Errors, "Tema no encontrado")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(json)
		return
	}

	data.Success = true
	data.Data = append(data.Data, temas)

	json, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func GetAllVideos(w http.ResponseWriter, req *http.Request) {
	var temas []models.Temas = models.GetAllTemas()

	var data = Data1{true, temas, make([]string, 0)}
	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func DeleteVideo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id_tema := vars["id"]

	var data = Data1{Errors: make([]string, 0)}

	cursos, success := models.DeleteTemas(id_tema)

	if success != true {
		data.Errors = append(data.Errors, "Tema no eliminado")
	}

	data.Success = true
	data.Data = append(data.Data, cursos)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func UpdateVideo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id_tema := vars["id"]

	bodyTemas, success := helpers.DecodeBodyTemas(req)

	if success != true {
		http.Error(w, "Error al leer el body", http.StatusBadRequest)
		return
	}

	var data Data1 = Data1{Errors: make([]string, 0)}

	bodyTemas.Description = strings.TrimSpace(bodyTemas.Description)
	bodyTemas.Title = strings.TrimSpace(bodyTemas.Title)

	if !helpers.IsValidTemas(bodyTemas.Description, bodyTemas.Title, bodyTemas.Curso_Id) {
		data.Success = false
		data.Errors = append(data.Errors, "invalid data")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(json)
		return

	}

	temas, success := models.UpdateTema(id_tema, bodyTemas.Title, bodyTemas.Curso_Id, bodyTemas.Description)

	if success != true {
		data.Errors = append(data.Errors, "Tema no actualizado")
	}

	data.Success = true
	data.Data = append(data.Data, temas)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(json)
	return
}

func GetAllVideosTema(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id_curso := vars["id"]

	var temas []models.Temas = models.GetAllTemasCursos(id_curso)
	var data = Data1{true, temas, make([]string, 0)}
	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
