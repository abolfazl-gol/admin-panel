package models

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Question struct {
	ID        int64      `db:"id" json:"id"`
	Text      string     `db:"text" json:"text"`
	Shows     int32      `db:"shows" json:"shows"`
	Enabled   bool       `db:"enabled" json:"enabled"`
	TopicID   int64      `db:"topic_id" json:"topic_id"`
	ImageID   int64      `db:"image_id" json:"image_id"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}

func (q *Question) String() string {
	return fmt.Sprintf("&Question{id: %d, text: %s, image_id: %d, enabled: %v, shows: %d, created_at: %s, updated_at: %s", q.ID, q.Text, q.ImageID, q.Enabled, q.Shows, q.CreatedAt, q.UpdatedAt)
}

func GetQuestion(id interface{}) (*Question, error) {
	question := new(Question)
	if err := db.Get(question, "SELECT * FROM questions WHERE id=?", id); err != nil {
		return nil, err
	}

	return question, nil
}

func GetQuestions() ([]*Question, error) {
	questions := make([]*Question, 0)
	if err := db.Select(&questions, "SELECT * FROM questions"); err != nil {
		return nil, err
	}

	return questions, nil
}

func CreateQuention(question *Question) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	// if image != nil {
	// 	if err := createImage(tx, image); err != nil {
	// 		tx.Rollback()
	// 		return err
	// 	}
	// 	question.ImageID = image.ID
	// }

	result, err := tx.Exec("INSERT INTO questions (text, topic_id, image_id, enabled, created_at) values (?, ?, ?, ?, ?)",
		question.Text, question.TopicID, question.ImageID, question.Enabled, question.CreatedAt,
	)

	if err != nil {
		tx.Rollback()
		if strings.Contains(err.Error(), "1062") {
			return errors.New("duplicate question text")
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	question.ID, _ = result.LastInsertId()
	return nil
}

func UpdateQuestion(question *Question, image *Image) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	if image != nil {
		oldImage, err := getImage(question.ImageID)
		if err == nil {
			os.Remove("public/images/" + oldImage.Hash)
		}

		if err := deleteImage(tx, question.ImageID); err != nil {
			tx.Rollback()
			return err
		}
		if err := createImage(tx, image); err != nil {
			tx.Rollback()
			return err
		}
		question.ImageID = image.ID
	}

	_, err = tx.Exec("UPDATE questions SET text=?, topic_id=?, image_id=?, enabled=?, updated_at=?  where id=?", question.Text, question.TopicID, question.ImageID, question.Enabled, question.UpdatedAt, question.ID)
	if err != nil {
		tx.Rollback()
		if strings.Contains(err.Error(), "1062") {
			return errors.New("duplicate question text")
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func DeleteQuestion(id interface{}, imageId interface{}) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	if _, err := tx.Exec("DELETE FROM questions WHERE id=?", id); err != nil {
		tx.Rollback()
		return err
	}

	if imageId != nil {
		if err := deleteImage(tx, imageId); err != nil {
			tx.Rollback()
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
