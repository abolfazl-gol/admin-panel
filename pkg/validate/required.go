package validate

import (
	"adminpanel/proto"
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

var (
	v = validator.New()
)

func RegisterRequest(req *proto.RegisterRequest) error {

	if err := v.Var(req.Email, "required,email"); err != nil {
		return fmt.Errorf("email is not valid")
	}
	if err := v.Var(req.Password, "required,gte=6"); err != nil {
		return fmt.Errorf("password is not valid")
	}
	return nil
}

func CreateTopicRequest(req *proto.CreateTopicRequest) error {
	if err := v.Var(req.Name, "required"); err != nil {
		return fmt.Errorf("name is required")
	}
	return nil
}
