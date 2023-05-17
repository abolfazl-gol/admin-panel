package convert

import (
	"adminpanel/models"
	"adminpanel/proto"
)

func TopicToProto(topic *models.Topic) *proto.Topic {
	return &proto.Topic{
		Id:      topic.ID,
		Name:    topic.Name,
		Enabled: topic.Enabled,
	}
}
