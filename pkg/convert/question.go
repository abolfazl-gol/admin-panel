package convert

import (
	"adminpanel/models"
	"adminpanel/proto"
)

func QuestionToProto(question *models.Question) *proto.Question {
	return &proto.Question{
		Id:      question.ID,
		Text:    question.Text,
		TopicId: question.TopicID,
		Enabled: question.Enabled,
		Shows:   question.Shows,
	}
}
