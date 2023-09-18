package service

import (
	"context"

	"github.com/elijahgrimaldi/Account-API/model"
	"github.com/google/uuid"
)

type UserService struct {
	UserRepository model.UserRepository
}

// SignUp implements model.UserService.
func (*UserService) Signup(ctx context.Context, u *model.User) error {
	panic("unimplemented")
}

// Get implements model.UserService.
func (u *UserService) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	user, err := u.UserRepository.FindByID(ctx, uid)
	return user, err
}

type USConfig struct {
	UserRepository model.UserRepository
}

func NewUserService(c *USConfig) model.UserService {
	return &UserService{
		UserRepository: c.UserRepository,
	}
}
