// user_service.go

package services

import (
	"cg-test/github.com/uumkon/cg-test/proto/user"
	"cg-test/internal/entities"
	"cg-test/internal/storage/repo"
	"context"
	"errors"
	"regexp"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceServer struct {
	repo *repo.UserRepo
	user.UnimplementedUserServiceServer
}

func NewUserServiceServer(repo *repo.UserRepo) *UserServiceServer {
	return &UserServiceServer{repo: repo}
}

func (s *UserServiceServer) Register(ctx context.Context, req *user.RegisterRequest) (*user.RegisterResponse, error) {
	login, password, fio := req.GetLogin(), req.GetPassword(), req.GetFio()

	if len(password) < 6 || !regexp.MustCompile(`[.,!_]`).MatchString(password) {
		return nil, errors.New("password must be at least 6 characters long and contain one of the characters: . , ! _")
	}

	apiToken := uuid.New().String()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := entities.User{
		Login:          login,
		HashedPassword: string(hashedPassword),
		FIO:            fio,
		APIToken:       apiToken,
	}

	err = s.repo.SaveUser(newUser)
	if err != nil {
		return nil, err
	}

	return &user.RegisterResponse{Fio: newUser.FIO, ApiToken: newUser.APIToken}, nil
}

func (s *UserServiceServer) Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	login, password := req.GetLogin(), req.GetPassword()

	userEntity, err := s.repo.FindUserByLogin(login)
	if err != nil {
		return nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userEntity.HashedPassword), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	return &user.LoginResponse{Fio: userEntity.FIO, ApiToken: userEntity.APIToken}, nil
}
