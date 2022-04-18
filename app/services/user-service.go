package services

import (
	"errors"

	"github.com/aliepba/go-crypto/app/methods"
	"github.com/aliepba/go-crypto/app/models"
	"github.com/aliepba/go-crypto/app/requests"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterInput(input requests.RegisterInput) (models.User, error)
	Login(input requests.LoginInput) (models.User, error)
	IsEmailAvailable(input requests.CheckEmailInput) (bool, error)
	GetUserByID(ID int) (models.User, error)
}

type userService struct {
	method methods.MethodUser
}

func NewServiceUser(method methods.MethodUser) *userService {
	return &userService{method}
}

func (s *userService) RegisterInput(input requests.RegisterInput) (models.User, error) {
	user := models.User{}
	user.Name = input.Name
	user.Email = input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	newUser, err := s.method.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *userService) Login(input requests.LoginInput) (models.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.method.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userService) IsEmailAvailable(input requests.CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.method.FindByEmail(email)

	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *userService) GetUserByID(ID int) (models.User, error) {
	user, err := s.method.FindById(ID)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found with that ID")
	}

	return user, nil
}
