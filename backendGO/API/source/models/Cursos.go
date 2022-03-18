package models

import (
	"API/source/database"
	"log"
)

type Cursos struct {
	Id_Curso    int    `json:"id_curso"`
	Nombre      string `json:"nombre"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func Insert(nombre string, description string, image string) (Cursos, bool) {
	db := database.GetConnection()

	var curso_id int
	db.QueryRow("INSERT INTO cursos (nombre, description, image) VALUES ($1, $2, $3) RETURNING id_curso", nombre, description, image).Scan(&curso_id)

	if curso_id == 0 {
		return Cursos{}, false
	}

	return Cursos{curso_id, nombre, description, image}, true
}

func Get(id string) (Cursos, bool) {
	db := database.GetConnection()
	row := db.QueryRow("SELECT * FROM cursos WHERE id_curso = $1", id)

	var id_curso int
	var nombre string
	var description string
	var image string
	err := row.Scan(&id_curso, &description, &image, &nombre)
	if err != nil {
		return Cursos{}, false
	}
	return Cursos{id_curso, nombre, description, image}, true
}

func GetAll() []Cursos {
	db := database.GetConnection()
	rows, err := db.Query("SELECT * FROM cursos ORDER BY id_curso")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var cursos []Cursos

	for rows.Next() {
		c := Cursos{}

		var id_curso int
		var nombre string
		var description string
		var image string

		err = rows.Scan(&id_curso, &description, &image, &nombre)

		if err != nil {
			log.Fatal(err)
		}
		c.Id_Curso = id_curso
		c.Nombre = nombre
		c.Description = description
		c.Image = image
		cursos = append(cursos, c)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return cursos
}

func Delete(id string) (Cursos, bool) {
	db := database.GetConnection()

	var id_curso int
	db.QueryRow("DELETE FROM cursos WHERE id_curso = $1 RETURNING id_curso", id).Scan(&id_curso)
	if id_curso == 0 {
		return Cursos{}, false
	}

	return Cursos{id_curso, "", "", ""}, true
}

func Update(id string, description string, image string, nombre string) (Cursos, bool) {
	db := database.GetConnection()

	var id_curso int
	db.QueryRow("UPDATE cursos SET description = $1, image = $2, nombre =$3 WHERE id_curso = $4 RETURNING id_curso", description, image, nombre, id).Scan(&id_curso)
	if id_curso == 0 {
		return Cursos{}, false
	}

	return Cursos{id_curso, nombre, description, image}, true
}
