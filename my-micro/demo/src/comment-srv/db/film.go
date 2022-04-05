package db

import (
	"database/sql"
	"oldboyedu-go/my-micro/demo/src/comment-srv/entity"
)

func SelectFilmImageAndName(filmId int64) (string, string, error) {

	film := entity.Film{}
	err := db.Get(&film, "SELECT `img`,`title_cn` FROM `film` WHERE `movie_id` = ?", filmId)
	if err == sql.ErrNoRows {
		return "", "", nil
	}
	return film.Img, film.TitleCn, err
}
