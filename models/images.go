package models

import "github.com/jmoiron/sqlx"

type Image struct {
	ID   int64  `db:"id" json:"id"`
	Hash string `db:"hash" json:"hash"`
	URL  string `db:"url" json:"url"`
}

func getImage(id interface{}) (*Image, error) {
	image := new(Image)
	err := db.Get(image, "select id, hash, url from images where id=?", id)
	if err != nil {
		return nil, err
	}
	return image, nil
}

func createImage(tx *sqlx.Tx, image *Image) error {
	result, err := tx.Exec("insert into images (hash, url) values (?,?)", image.Hash, image.URL)
	if err != nil {
		return err
	}

	image.ID, _ = result.LastInsertId()
	return nil
}

func deleteImage(tx *sqlx.Tx, id interface{}) error {
	if _, err := tx.Exec("DELETE FROM images WHERE id=?", id); err != nil {
		return err
	}

	return nil
}
