package service

import (
	"strings"

	"github.com/revandpratama/go-simple-blog-api/dto"
	"github.com/revandpratama/go-simple-blog-api/entity"
	"github.com/revandpratama/go-simple-blog-api/errorhandler"
	"github.com/revandpratama/go-simple-blog-api/helper"
	"github.com/revandpratama/go-simple-blog-api/repository"
)

type UserService interface {
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
	Register(req *dto.RegisterRequest) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(r repository.UserRepository) *userService {
	return &userService{
		userRepository: r,
	}
}

func (s *userService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	var data dto.LoginResponse
	var user *entity.User
	var err error

	emailPrefix := "@"

	if strings.Contains(req.EmailOrUsername, emailPrefix) {
		user, err = s.userRepository.GetUserByEmail(req.EmailOrUsername)
		if err != nil {
			return nil, &errorhandler.NotFoundError{Message: "wrong credentials"}
		}
	} else {
		user, err = s.userRepository.GetUserByUsername(req.EmailOrUsername)
		if err != nil {
			return nil, &errorhandler.NotFoundError{Message: "wrong credentials"}
		}
	}

	if err := helper.VerifyPassword(req.Password, user.Password); err != nil {
		return nil, err
	}

	token, err := helper.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	data = dto.LoginResponse{
		ID:   user.ID,
		Name: user.Username,
		Token: token,
	}

	return &data, nil

}

func (s *userService) Register(req *dto.RegisterRequest) error {

	if emailExists := s.userRepository.EmailExists(req.Email); emailExists {
		return &errorhandler.BadRequestError{Message: "this email is already in use"}
	}

	if usernameExists := s.userRepository.UsernameExists(req.Username); usernameExists {
		return &errorhandler.BadRequestError{Message: "this username is already taken"}
	}

	if req.Password != req.PasswordConfirm {
		return &errorhandler.BadRequestError{Message: "password and confirmation password do not match"}
	}

	hashedPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		return &errorhandler.InternalServerError{Message: "Could not create password hash"}
	}

	user := entity.User{
		Name:     req.Name,
		Username: req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := s.userRepository.Register(&user); err != nil {
		return &errorhandler.InternalServerError{Message: "failed to register the new account"}
	}

	return nil

}
