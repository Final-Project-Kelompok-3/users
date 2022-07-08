package user

import (
	"context"
	"errors"

	"github.com/Final-Project-Kelompok-3/users/pkg/constant"
	res "github.com/Final-Project-Kelompok-3/users/pkg/util/response"

	"github.com/Final-Project-Kelompok-3/users/internal/dto"
	"github.com/Final-Project-Kelompok-3/users/internal/factory"
	"github.com/Final-Project-Kelompok-3/users/internal/model"
	"github.com/Final-Project-Kelompok-3/users/internal/repository"
)

type service struct {
	UserRepository repository.User
}

type Service interface {
	Register(ctx context.Context, payload *dto.CreateUserRequest) (string, error)
	Login(ctx context.Context, email, password string) (string, error)
	FindAll(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.User], error)
	FindByID(ctx context.Context, payload *dto.ByIDRequest) (*model.User, error)
	Create(ctx context.Context, payload *dto.CreateUserRequest) (string, error)
	Update(ctx context.Context, ID uint, payload *dto.UpdateUserRequest) (string, error)
	Delete(ctx context.Context, ID uint) (*model.User, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		UserRepository: f.UserRepository,
	}
}

func (s *service) Register(ctx context.Context, payload *dto.CreateUserRequest) (string, error) {

	data, _ := s.UserRepository.FindByEmail(ctx, payload.Email)
	if (data != model.User{}) {
		return "", res.ErrorBuilder(&res.ErrorConstant.Duplicate, errors.New("email already exist"))
	}

	var user = model.User{
		RoleID: 	payload.RoleID,
		FirstName: 	payload.FirstName,
		LastName: 	payload.LastName,
		Email:      payload.Email,
		Password: 	payload.Password,
	}

	err := s.UserRepository.Create(ctx, user)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}

func (s *service) Login(ctx context.Context, email, password string) (string, error) {
	
	token, err := s.UserRepository.Login(ctx, email, password)
	if err != nil {
		if err == constant.RecordNotFound {
			return "", res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return token, nil
}

func (s *service) FindAll(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.User], error) {
	
	Users, info, err := s.UserRepository.FindAll(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	
	result := new(dto.SearchGetResponse[model.User])
	result.Datas = Users
	result.PaginationInfo = *info

	return result, nil
}

func (s *service) FindByID(ctx context.Context, payload *dto.ByIDRequest) (*model.User, error) {

	data, err := s.UserRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil
}

func (s *service) Create(ctx context.Context, payload *dto.CreateUserRequest) (string, error) {

	data, _ := s.UserRepository.FindByEmail(ctx, payload.Email)
	if (data != model.User{}) {
		return "", res.ErrorBuilder(&res.ErrorConstant.Duplicate, errors.New("email already exist"))
	}

	var user = model.User{
		RoleID: 	payload.RoleID,
		FirstName: 	payload.FirstName,
		LastName: 	payload.LastName,
		Email:      payload.Email,
		Password: 	payload.Password,
	}

	err := s.UserRepository.Create(ctx, user)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}

func (s *service) Update(ctx context.Context, ID uint, payload *dto.UpdateUserRequest) (string, error) {
	var data = make(map[string]interface{})

	if payload.RoleID != nil {
		data["role_id"] = payload.RoleID
	}
	if payload.FirstName != nil {
		data["first_name"] = payload.FirstName
	}
	if payload.LastName != nil {
		data["last_name"] = payload.LastName
	}
	if payload.Email != nil {
		data["email"] = payload.Email
	}
	if payload.Password != nil {
		data["password"] = payload.Password
	}

	err := s.UserRepository.Update(ctx, ID, data)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}

func (s *service) Delete(ctx context.Context, ID uint) (*model.User, error) {
	data, err := s.UserRepository.FindByID(ctx, ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	err = s.UserRepository.Delete(ctx, ID)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil
}