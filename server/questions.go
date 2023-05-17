package server

import (
	"adminpanel/models"
	"adminpanel/pkg/convert"
	"adminpanel/proto"
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) ListQuestion(ctx context.Context, req *proto.ListQuestionRequest) (*proto.ListQuestionResponse, error) {
	questions, err := models.GetQuestions()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal: %v", err)
	}

	protoQuestion := make([]*proto.Question, 0)
	for _, question := range questions {
		protoQuestion = append(protoQuestion, convert.QuestionToProto(question))
	}

	return &proto.ListQuestionResponse{Questions: protoQuestion}, nil
}

func (s *Service) GetQuestion(ctx context.Context, req *proto.GetQuestionRequest) (*proto.Question, error) {
	question, err := models.GetQuestion(req.Id)
	if err != nil {
		if err == models.ErrNotFound {
			return nil, status.Error(codes.NotFound, "question not found")
		}

		return nil, status.Errorf(codes.Internal, "internal: %v", err)
	}

	return convert.QuestionToProto(question), nil
}

func (s *Service) CreateQuestion(ctx context.Context, req *proto.CreateQuestionRequest) (*proto.Question, error) {
	if req.Text == "" {
		return nil, status.Error(codes.InvalidArgument, "field required")
	}
	question := &models.Question{
		Text:      req.Text,
		TopicID:   req.TopicId,
		Enabled:   req.Enabled,
		CreatedAt: &t,
	}

	if err := models.CreateQuention(question); err != nil {
		return nil, fmt.Errorf("can't create question: %v", err)
	}

	return convert.QuestionToProto(question), nil
}

func (s *Service) UpdateQuestion(ctx context.Context, req *proto.UpdateQuestionRequest) (*proto.Question, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQuestion not implemented")
}
func (s *Service) DeleteQuestion(ctx context.Context, req *proto.DeleteQuestionRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQuestion not implemented")
}
