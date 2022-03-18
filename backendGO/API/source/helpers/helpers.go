package helpers

import (
	"API/source/models"
	"encoding/json"
	"net/http"
	"strings"
)

func DecodeBody(req *http.Request) (models.Cursos, bool) {
	var cursos models.Cursos
	err := json.NewDecoder(req.Body).Decode(&cursos)
	if err != nil {
		return models.Cursos{}, false
	}

	return cursos, true
}

func IsValid(description string, nombre string) bool {
	desc := strings.TrimSpace(description)
	nom := strings.TrimSpace(nombre)
	if len(desc) == 0 || len(nom) == 0 {
		return false
	}
	return true
}

func DecodeBodyTemas(req *http.Request) (models.Temas, bool) {
	var temas models.Temas
	err := json.NewDecoder(req.Body).Decode(&temas)
	if err != nil {
		return models.Temas{}, false
	}

	return temas, true
}

func IsValidTemas(description string, title string, curso_id int) bool {
	desc := strings.TrimSpace(description)
	tit := strings.TrimSpace(title)
	curs_id := strings.TrimSpace(string(curso_id))
	if len(desc) == 0 || len(tit) == 0 || len(curs_id) == 0 {
		return false
	}
	return true
}

func DecodeBodyVideos(req *http.Request) (models.Videos, bool) {
	var videos models.Videos
	err := json.NewDecoder(req.Body).Decode(&videos)
	if err != nil {
		return models.Videos{}, false
	}
	return videos, true
}

func IsValidVideos(url_video string, nivel int, id_tema_id int) bool {
	url := strings.TrimSpace(url_video)
	niv := strings.TrimSpace(string(nivel))
	id_tema := strings.TrimSpace(string(id_tema_id))

	if len(url) == 0 || len(niv) == 0 || len(id_tema) == 0 {
		return false
	}
	return true
}
