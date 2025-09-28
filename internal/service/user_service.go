package service

import (
	"context"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"

	"social_network/internal/domain"
	"social_network/internal/repository"
)

type UserService interface {
	Register(ctx context.Context, email, password string) (*domain.User, error)
	CreateProfile(ctx context.Context, firstName, secondName string, birthdate time.Time, biography, city, password string) (*domain.User, error)
	GetByID(ctx context.Context, id string) (*domain.User, error)
	Login(ctx context.Context, id, password string) (*domain.User, error)
}

type userService struct {
	users repository.UserRepository
}

func NewUserService(users repository.UserRepository) UserService {
	return &userService{users: users}
}

func (s *userService) Register(ctx context.Context, email, password string) (*domain.User, error) {
	// бизнес-валидации поверх DTO (если нужны)
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return s.users.Create(ctx, email, string(hash))
}

func (s *userService) CreateProfile(ctx context.Context, firstName, secondName string, birthdate time.Time, biography, city, password string) (*domain.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	log.Printf("CreateProfile called: firstName=%s, secondName=%s, birthdate=%s, biography=%s, city=%s, password=%s", firstName, secondName, birthdate, biography, city, password)
	return s.users.CreateProfile(ctx, firstName, secondName, birthdate, biography, city, string(hash))
}

func (s *userService) GetByID(ctx context.Context, id string) (*domain.User, error) {
	return s.users.GetById(ctx, id)
}

func (s *userService) Login(ctx context.Context, id, password string) (*domain.User, error) {
	user, err := s.users.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}
