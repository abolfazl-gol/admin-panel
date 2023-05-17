package models

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Topic struct {
	ID        int64      `db:"id" json:"id"`
	Name      string     `db:"name" json:"name"`
	ImageID   int64      `db:"image_id" json:"image_id"`
	Enabled   bool       `db:"enabled" json:"enabled"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}

func (t *Topic) String() string {
	return fmt.Sprintf("&Topics{id: %d, name: %s, image_id: %d, enabled: %v, created_at: %s, updated_at: %s", t.ID, t.Name, t.ImageID, t.Enabled, t.CreatedAt, t.UpdatedAt)
}

func GetTopic(id interface{}) (*Topic, error) {
	topic := new(Topic)
	if err := db.Get(topic, "SELECT * FROM topics where id= ?", id); err != nil {
		return nil, err
	}

	return topic, nil
}

func GetTopics() ([]*Topic, error) {
	topics := make([]*Topic, 0)
	if err := db.Select(&topics, "SELECT * FROM topics"); err != nil {
		return nil, err
	}

	return topics, nil

}

func CreateTopic(topic *Topic) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	// if image != nil {
	// 	if err := createImage(tx, image); err != nil {
	// 		tx.Rollback()
	// 		return err
	// 	}
	// 	topic.ImageID = image.ID
	// }

	result, err := tx.Exec(
		"INSERT INTO topics (name, enabled, created_at) VALUES (?, ?, ?)",
		topic.Name, topic.Enabled, topic.CreatedAt,
	)

	if err != nil {
		tx.Rollback()
		if strings.Contains(err.Error(), "1062") {
			return errors.New("duplicate topic name")
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	topic.ID, _ = result.LastInsertId()
	return nil
}

func UpdateTopic(topic *Topic, cols ...string) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	// if image != nil {
	// 	// delete from folder --> /public/images
	// 	oldImage, err := getImage(topic.ImageID)
	// 	if err == nil {
	// 		os.Remove("public/images/" + oldImage.Hash)
	// 	}

	// 	// delete from db
	// 	if err := deleteImage(tx, topic.ImageID); err != nil {
	// 		tx.Rollback()
	// 		return err
	// 	}

	// 	// created new image in db
	// 	if err := createImage(tx, image); err != nil {
	// 		tx.Rollback()
	// 		return err
	// 	}
	// 	topic.ImageID = image.ID
	// }

	var updates string
	for _, col := range cols {
		updates += fmt.Sprintf("%s = :%s", col, col) + ","
	}
	updates = updates[:len(updates)-1]
	query := fmt.Sprintf("UPDATE topics SET %s where id=:id", updates)

	_, err = tx.NamedExec(query, topic)
	if err != nil {
		tx.Rollback()
		if strings.Contains(err.Error(), "1062") {
			return errors.New("duplicate topic name")
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return err
}

func DeleteTopic(id interface{}) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	if _, err := tx.Exec("DELETE FROM topics WHERE id=?", id); err != nil {
		tx.Rollback()
		return err
	}
	// if imageId != nil {
	// 	if err := deleteImage(tx, imageId); err != nil {
	// 		tx.Rollback()
	// 		return err
	// 	}
	// }
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
