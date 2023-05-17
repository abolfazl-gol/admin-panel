package server

import (
	"adminpanel/proto"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) ListAnswer(ctx context.Context, req *proto.ListAnswerRequest) (*proto.ListAnswerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAnswer not implemented")
}
func (s *Service) GetAnswer(ctx context.Context, req *proto.GetAnswerRequest) (*proto.Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnswer not implemented")
}
func (s *Service) CreateAnswer(ctx context.Context, req *proto.CreateAnswerRequest) (*proto.Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAnswer not implemented")
}
func (s *Service) UpdateAnswer(ctx context.Context, req *proto.UpdateAnswerRequest) (*proto.Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAnswer not implemented")
}
func (s *Service) DeleteAnswer(ctx context.Context, req *proto.DeleteAnswerRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAnswer not implemented")
}
