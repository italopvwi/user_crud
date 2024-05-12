package services

import (
	"context"
	"user_crud/internal/repositories"
	"user_crud/pkg/models"
)

type UserService interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id int) (*models.User, error)
	GetAll(ctx context.Context) ([]*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id int) error
}

type userServiceImpl struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userServiceImpl{userRepo: userRepo}
}

func (s *userServiceImpl) Create(ctx context.Context, user *models.User) error {
	return s.userRepo.Create(ctx, user)
}

func (s *userServiceImpl) GetByID(ctx context.Context, id int) (*models.User, error) {
	return s.userRepo.GetByID(ctx, id)
}

func (s *userServiceImpl) GetAll(ctx context.Context) ([]*models.User, error) {
	return s.userRepo.GetAll(ctx)
}

func (s *userServiceImpl) Update(ctx context.Context, user *models.User) error {
	return s.userRepo.Update(ctx, user)
}

func (s *userServiceImpl) Delete(ctx context.Context, id int) error {
	return s.userRepo.Delete(ctx, id)
}
