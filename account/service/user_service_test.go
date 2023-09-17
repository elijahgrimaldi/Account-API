package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/elijahgrimaldi/Account-API/model"
	"github.com/elijahgrimaldi/Account-API/model/mocks"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockUserResp := &model.User{
			UID:   uid,
			Email: "bob@bob.com",
			Name:  "Bobby Bobson",
		}

		mockUserRepository := new(mocks.MockUserRepository)
		us := NewUserService(&USConfig{
			UserRepository: mockUserRepository,
		})
		mockUserRepository.On("FindByID", mock.Anything, uid).Return(mockUserResp, nil)
		ctx := context.TODO()
		u, err := us.Get(ctx, uid)
		assert.Equal(t, mockUserResp, u)
		assert.NoError(t, err)
		mockUserRepository.AssertExpectations(t) // assert that UserService.Get was called
	})

	t.Run("Error", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockUserRepository := new(mocks.MockUserRepository)
		us := NewUserService(&USConfig{
			UserRepository: mockUserRepository,
		})
		mockUserRepository.On("FindByID", mock.Anything, uid).Return(nil, fmt.Errorf("Some error down the call chain"))
		ctx := context.TODO()
		u, err := us.Get(ctx, uid)
		assert.Nil(t, u)
		assert.Error(t, err)
		mockUserRepository.AssertExpectations(t)
	})
}
