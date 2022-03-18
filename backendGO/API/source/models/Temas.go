package models

import (
	"API/source/database"
	"log"
)

type Temas struct {
	Id_Tema     int    `json:"id_tema"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Curso_Id    int    `json:"curso_id"`
}

func InsertTema(title string, description string, curso_id int) (Temas, bool) {
	db := database.GetConnection()

	var tema_id int
	db.QueryRow("INSERT INTO temas (title, description, curso_id) VALUES ($1, $2, $3) RETURNING id_tema", title, description, curso_id).Scan(&tema_id)

	if tema_id == 0 {
		return Temas{}, false
	}

	return Temas{tema_id, title, description, curso_id}, true
}

func GetTema(id string) (Temas, bool) {
	db := database.GetConnection()
	row := db.QueryRow("SELECT * FROM temas WHERE id_tema = $1", id)

	var id_tema int
	var title string
	var description string
	var curso_id uint
	err := row.Scan(&id_tema, &title, &curso_id, &description)
	if err != nil {
		return Temas{}, false
	}
	return Temas{id_tema, title, description, int(curso_id)}, true
}

func GetAllTemas() []Temas {
	db := database.GetConnection()
	rows, err := db.Query("SELECT * FROM temas ORDER BY id_tema ASC")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var temas []Temas

	for rows.Next() {
		t := Temas{}

		var id_tema int
		var title string
		var description string
		var curso_id uint

		err = rows.Scan(&id_tema, &title, &curso_id, &description)

		if err != nil {
			log.Fatal("Error de Scan", err)
		}
		t.Id_Tema = id_tema
		t.Title = title
		t.Description = description
		t.Curso_Id = int(curso_id)
		temas = append(temas, t)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return temas
}

func GetAllTemasCursos(id string) []Temas {
	db := database.GetConnection()
	row, err := db.Query("SELECT * FROM temas WHERE curso_id = $1", id)

	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var temas []Temas

	for row.Next() {
		t := Temas{}

		var id_tema int
		var title string
		var description string
		var curso_id uint

		err = row.Scan(&id_tema, &title, &curso_id, &description)

		if err != nil {
			log.Fatal("error de Scan", err)
		}
		t.Id_Tema = id_tema
		t.Title = title
		t.Description = description
		t.Curso_Id = int(curso_id)
		temas = append(temas, t)
	}
	err = row.Err()
	if err != nil {
		log.Fatal(err)
	}
	return temas
}

func DeleteTemas(id string) (Temas, bool) {
	db := database.GetConnection()

	var id_tema int
	db.QueryRow("DELETE FROM temas WHERE id_tema = $1 RETURNING id_tema", id).Scan(&id_tema)
	if id_tema == 0 {
		return Temas{}, false
	}

	return Temas{id_tema, "", "", 0}, true
}

func UpdateTema(id string, title string, curso_id int, description string) (Temas, bool) {
	db := database.GetConnection()

	var id_tema int
	db.QueryRow("UPDATE temas SET title =$1, curso_id =$2, description = $3 WHERE id_tema = $4 RETURNING id_tema", title, curso_id, description, id).Scan(&id_tema)
	if id_tema == 0 {
		return Temas{}, false
	}

	return Temas{id_tema, title, description, curso_id}, true
}
