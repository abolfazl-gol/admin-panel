package server

import (
	"adminpanel/models"
	"adminpanel/pkg/convert"
	"adminpanel/pkg/validate"
	"adminpanel/proto"
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	t = time.Now()
)

type Service struct {
}

func Start() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	proto.RegisterApiServer(server, &Service{})
	fmt.Println("server started on port", ":50051")
	return server.Serve(lis)
}

// User
func (s *Service) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.User, error) {
	if err := validate.RegisterRequest(req); err != nil {
		return nil, err
	}
	user := &models.User{
		Email: req.Email,
		Token: uuid.Must(uuid.NewV4()).String(),
	}

	user.GenerateHash(req.Password)

	if err := models.CreateUser(user); err != nil {
		return nil, fmt.Errorf("can't create user: %v", err)
	}

	return convert.UserToProto(user), nil
}

func (s *Service) Login(ctx context.Context, req *proto.LoginRequest) (*proto.User, error) {
	user, err := models.GetUserByEmail(req.Email)
	if err != nil {
		if err == models.ErrNotFound {
			return nil, status.Error(codes.Unauthenticated, "invalid email/password")
		}

		return nil, status.Errorf(codes.Internal, "Internal: %v", err)
	}

	if !user.Authenticate(req.Password) {
		return nil, status.Error(codes.Unauthenticated, "invalid email/password")
	}

	return convert.UserToProto(user), nil
}
