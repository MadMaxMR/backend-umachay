package models

import (
	"API/source/database"
	"log"
)

type Videos struct {
	Id_Video   int    `json:"id_video"`
	Url_Video  string `json:"url_video"`
	Nivel      int    `json:"nivel"`
	Id_Tema_Id int    `json:"id_tema"`
}

func InsertVideo(url_video string, nivel int, id_tema_id int) (Videos, bool) {
	db := database.GetConnection()

	var video_id int
	db.QueryRow("INSERT INTO videos (url_video, nivel, id_tema_id) VALUES ($1, $2, $3) RETURNING id_video", url_video, nivel, id_tema_id).Scan(&video_id)

	if video_id == 0 {
		return Videos{}, false
	}

	return Videos{video_id, url_video, nivel, id_tema_id}, true
}

func GetVideo(id string) (Videos, bool) {
	db := database.GetConnection()
	row := db.QueryRow("SELECT * FROM videos WHERE id_video = $1", id)

	var id_video int
	var url_video string
	var nivel int
	var id_tema_id uint

	err := row.Scan(&id_video, &url_video, &nivel, &id_tema_id)
	if err != nil {
		return Videos{}, false
	}
	return Videos{id_video, url_video, nivel, int(id_tema_id)}, true
}

func GetAllVideos() []Videos {
	db := database.GetConnection()
	rows, err := db.Query("SELECT * FROM videos ORDER BY id_video ASC")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var videos []Videos

	for rows.Next() {
		t := Videos{}

		var id_video int
		var url_video string
		var nivel int
		var id_tema_id uint

		err = rows.Scan(&id_video, &url_video, &nivel, &id_tema_id)

		if err != nil {
			log.Fatal("Error de Scan", err)
		}
		t.Id_Video = id_video
		t.Url_Video = url_video
		t.Nivel = nivel
		t.Id_Tema_Id = int(id_tema_id)
		videos = append(videos, t)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return videos
}

func GetAllVideosTema(id string) []Videos {
	db := database.GetConnection()
	row, err := db.Query("SELECT * FROM videos WHERE id_tema_id = $1", id)

	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var videos []Videos

	for row.Next() {
		t := Videos{}

		var id_video int
		var url_video string
		var nivel int
		var id_tema_id uint

		err = row.Scan(&id_video, &url_video, &nivel, &id_tema_id)

		if err != nil {
			log.Fatal("error de Scan", err)
		}
		t.Id_Video = id_video
		t.Url_Video = url_video
		t.Nivel = nivel
		t.Id_Tema_Id = int(id_tema_id)
		videos = append(videos, t)
	}
	err = row.Err()
	if err != nil {
		log.Fatal(err)
	}
	return videos
}

func DeleteVideo(id string) (Videos, bool) {
	db := database.GetConnection()

	var id_video int
	db.QueryRow("DELETE FROM videos WHERE id_video = $1 RETURNING id_tema", id).Scan(&id_video)
	if id_video == 0 {
		return Videos{}, false
	}

	return Videos{id_video, "", 0, 0}, true
}

func UpdateVideo(id string, url_video string, nivel int, id_tema_id int) (Videos, bool) {
	db := database.GetConnection()

	var id_video int
	db.QueryRow("UPDATE videos SET url_video =$1, nivel =$2, id_tema_id = $3 WHERE id_video = $4 RETURNING id_video", url_video, nivel, id_tema_id, id).Scan(&id_video)
	if id_video == 0 {
		return Videos{}, false
	}

	return Videos{id_video, url_video, nivel, id_tema_id}, true
}
