package service

import (
	"context"
	"fmt"

	"github.com/halllllll/MaGRO/entity"

	"golang.org/x/crypto/bcrypt"
)

type RegisterUser struct {
	Repo UserRegister
}

func (ru *RegisterUser) RegisterUser(ctx context.Context, name, password, role string) (*entity.User_T, error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot hash password: %v", err.Error())
	}
	u := &entity.User_T{
		Name:     name,
		Password: string(pw),
		Role:     role,
	}
	newUser, err := ru.Repo.RegisterUser(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("failed to register user: %v", err.Error())
	}
	return newUser, nil
}
