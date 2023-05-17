package convert

import (
	"adminpanel/models"
	"adminpanel/proto"
)

func UserToProto(user *models.User) *proto.User {
	return &proto.User{
		Id:    user.ID,
		Email: user.Email,
		Token: user.Token,
	}
}
