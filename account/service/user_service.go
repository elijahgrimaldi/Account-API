package service

import (
	"context"
	"log"

	"github.com/elijahgrimaldi/Account-API/model"
	"github.com/elijahgrimaldi/Account-API/model/apperrors"
	"github.com/google/uuid"
)

type UserService struct {
	UserRepository model.UserRepository
}

// SignUp implements model.UserService.
func (s *UserService) Signup(ctx context.Context, u *model.User) error {

	newPassword, err := hashPassword(u.Password)
	if err != nil {
		log.Printf("Unable to signup user for email: %v\n", u.Email)
		return apperrors.NewInternal()
	}
	u.Password = newPassword
	if err := s.UserRepository.Create(ctx, u); err != nil {
		return err
	}
	return nil
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
