package server

import (
	"adminpanel/models"
	"adminpanel/pkg/convert"
	"adminpanel/pkg/validate"
	"adminpanel/proto"
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/thoas/go-funk"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Topic
func (s *Service) GetTopic(ctx context.Context, req *proto.GetTopicRequest) (*proto.Topic, error) {
	topic, err := models.GetTopic(req.Id)
	if err != nil {
		if err == models.ErrNotFound {
			return nil, status.Error(codes.NotFound, "topic not found")
		}

		return nil, status.Errorf(codes.Internal, "internal: %v", err)
	}

	return convert.TopicToProto(topic), nil
}

func (s *Service) ListTopic(ctx context.Context, req *proto.ListTopicRequest) (*proto.ListTopicResponse, error) {
	topics, err := models.GetTopics()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal: %v", err)
	}

	protoTopics := make([]*proto.Topic, 0)
	for _, topic := range topics {
		protoTopics = append(protoTopics, convert.TopicToProto(topic))
	}
	return &proto.ListTopicResponse{Topics: protoTopics}, nil
}

func (s *Service) CreateTopic(ctx context.Context, req *proto.CreateTopicRequest) (*proto.Topic, error) {
	if err := validate.CreateTopicRequest(req); err != nil {
		return nil, err
	}
	topic := &models.Topic{
		Name:      req.Name,
		Enabled:   req.Enabled,
		CreatedAt: &t,
	}

	if err := models.CreateTopic(topic); err != nil {
		return nil, fmt.Errorf("can't create topic : %v", err)
	}

	return convert.TopicToProto(topic), nil
}

var permittedCols = []string{"name", "enabled", "updated_at"}

func (s *Service) UpdateTopic(ctx context.Context, req *proto.UpdateTopicRequest) (*proto.Topic, error) {
	if req.Topic == nil || req.Topic.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "topic not set")
	}
	_, err := models.GetTopic(req.Topic.Id)
	if err != nil {
		if err == models.ErrNotFound {
			return nil, status.Error(codes.NotFound, "topic not found")
		}

		return nil, status.Errorf(codes.Internal, "internal: %v", err)
	}

	topic := &models.Topic{
		ID:        req.Topic.Id,
		Name:      req.Topic.Name,
		Enabled:   req.Topic.Enabled,
		UpdatedAt: &t,
	}

	cols := make([]string, 0)
	for _, col := range req.UpdateMask {
		if funk.ContainsString(permittedCols, col) {
			cols = append(cols, col)
		}
	}

	if err := models.UpdateTopic(topic, cols...); err != nil {
		return nil, status.Errorf(codes.Internal, "internal: %v", err)
	}

	return convert.TopicToProto(topic), nil
}

func (s *Service) DeleteTopic(ctx context.Context, req *proto.DeleteTopicRequest) (*empty.Empty, error) {
	topic, err := models.GetTopic(req.Id)
	if err != nil {
		if err == models.ErrNotFound {
			return nil, status.Errorf(codes.NotFound, "topic not found")
		}

		return nil, status.Errorf(codes.Internal, "internal: %v", err)
	}

	if err := models.DeleteTopic(topic.ID); err != nil {
		return nil, status.Errorf(codes.Internal, "internal: %v", err)
	}

	return &empty.Empty{}, nil
}
