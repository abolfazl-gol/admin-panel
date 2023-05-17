package models

import (
	"fmt"
	"os"
	"time"
)

type Answer struct {
	ID         int64      `db:"id" json:"id"`
	Text       string     `db:"text" json:"text"`
	Enabled    bool       `db:"enabled" json:"enabled"`
	ImageID    int64      `db:"image_id" json:"image_id"`
	QuestionID int64      `db:"question_id" json:"question_id"`
	Correct    bool       `db:"correct" json:"correct"`
	CreatedAt  *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  *time.Time `db:"updated_at" json:"updated_at"`
}

func (a *Answer) String() string {
	return fmt.Sprintf("&Answer{id:%d, text:%s, question_id:%d, image_id:%d, enabled:%v, correct:%v, created_at:%s, updated_at:%s",
		a.ID, a.Text, a.QuestionID, a.ImageID, a.Enabled, a.Correct, a.CreatedAt, a.UpdatedAt,
	)
}

func GetAnswer(id interface{}) (*Answer, error) {
	answer := new(Answer)
	if err := db.Get(answer, "SELECT * FROM answers WHERE id=?", id); err != nil {
		return nil, err
	}

	return answer, nil
}

func GetAnswers() ([]*Answer, error) {
	answers := make([]*Answer, 0)
	if err := db.Select(&answers, "SELECT * FROM answers"); err != nil {
		return nil, err
	}

	return answers, nil
}

func CreateAnswer(answer *Answer, image *Image) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	if image != nil {
		if err := createImage(tx, image); err != nil {
			tx.Rollback()
			return err
		}
		answer.ImageID = image.ID
	}

	result, err := tx.Exec("INSERT INTO answers (text, question_id, image_id, enabled, correct, created_at) VALUES(?,?,?,?,?,?)",
		answer.Text, answer.QuestionID, answer.ImageID, answer.Enabled, answer.Correct, answer.CreatedAt,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	answer.ID, _ = result.LastInsertId()
	return nil
}

func UpdatedAnswer(answer *Answer, image *Image) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	if image != nil {
		oldImage, err := getImage(answer.ImageID)
		if err == nil {
			os.Remove("public/images/" + oldImage.Hash)
		}

		if err := deleteImage(tx, answer.ImageID); err != nil {
			tx.Rollback()
			return err
		}

		if err := createImage(tx, image); err != nil {
			tx.Rollback()
			return err
		}
		answer.ImageID = image.ID
	}

	if _, err = tx.Exec("UPDATE answers SET text=?, question_id=?, image_id=?, enabled=?, correct=? where id=?",
		answer.Text, answer.QuestionID, answer.ImageID, answer.Enabled, answer.Correct, answer.ID); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func DeleteAnswer(id interface{}, imageID interface{}) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	if _, err := tx.Exec("DELETE FROM answers WHERE id=?", id); err != nil {
		tx.Rollback()
		return err
	}

	if imageID != nil {
		if err := deleteImage(tx, imageID); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
