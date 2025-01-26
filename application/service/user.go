package service

import (
	"go-api-sample/domain/entity"
	"go-api-sample/infrastructure/repository"

	"github.com/google/uuid"
)

type (
	UserService interface {
		CreateUser(user *entity.User) (*entity.User, error)
		GetUserByID(id string) (*entity.User, error)
		GetUserByEmail(email string) (*entity.User, error)
		GetUsers() ([]entity.User, error)
		UpdateUser(user *entity.User) (*entity.User, error)
		DeleteUser(id string) error
	}
)

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) CreateUser(user *entity.User) (*entity.User, error) {
	user.ID = uuid.New().String()
	return s.userRepository.Create(user)
}

func (s *userService) GetUserByID(id string) (*entity.User, error) {
	return s.userRepository.FindByID(id)
}

func (s *userService) GetUserByEmail(email string) (*entity.User, error) {
	return s.userRepository.FindByEmail(email)
}

func (s *userService) GetUsers() ([]entity.User, error) {
	return s.userRepository.FindAll()
}

func (s *userService) UpdateUser(user *entity.User) (*entity.User, error) {
	return s.userRepository.Save(user)
}

func (s *userService) DeleteUser(id string) error {
	return s.userRepository.Delete(id)
}
