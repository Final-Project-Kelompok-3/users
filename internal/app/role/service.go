package role

import (
	"context"

	"github.com/Final-Project-Kelompok-3/users/pkg/constant"
	res "github.com/Final-Project-Kelompok-3/users/pkg/util/response"

	"github.com/Final-Project-Kelompok-3/users/internal/dto"
	"github.com/Final-Project-Kelompok-3/users/internal/factory"
	"github.com/Final-Project-Kelompok-3/users/internal/model"
	"github.com/Final-Project-Kelompok-3/users/internal/repository"
)

type Service interface {
	FindAll(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.Role], error)
	FindByID(ctx context.Context, payload *dto.ByIDRequest) (*model.Role, error)
	Create(ctx context.Context, payload *dto.CreateRoleRequest) (string, error)
	Update(ctx context.Context, ID uint, payload *dto.UpdateRoleRequest) (string, error)
	Delete(ctx context.Context, ID uint) (*model.Role, error)
}

type service struct {
	RoleRepository repository.Role
}

func NewService(f *factory.Factory) Service {
	return &service{
		RoleRepository: f.RoleRepository,
	}
}

func (s *service) FindAll(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.Role], error) {
	
	Roles, info, err := s.RoleRepository.FindAll(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	
	result := new(dto.SearchGetResponse[model.Role])
	result.Datas = Roles
	result.PaginationInfo = *info

	return result, nil
}

func (s *service) FindByID(ctx context.Context, payload *dto.ByIDRequest) (*model.Role, error) {

	data, err := s.RoleRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil
}

func (s *service) Create(ctx context.Context, payload *dto.CreateRoleRequest) (string, error) {
	
	var role = model.Role{
		Name:       payload.Name,
		Description: payload.Description,
	}

	err := s.RoleRepository.Create(ctx, role)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}

func (s *service) Update(ctx context.Context, ID uint, payload *dto.UpdateRoleRequest) (string, error) {
	var data = make(map[string]interface{})

	if payload.Name != nil {
		data["title"] = payload.Name
	}
	if payload.Description != nil {
		data["description"] = payload.Description
	}

	err := s.RoleRepository.Update(ctx, ID, data)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}

func (s *service) Delete(ctx context.Context, ID uint) (*model.Role, error) {
	data, err := s.RoleRepository.FindByID(ctx, ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	err = s.RoleRepository.Delete(ctx, ID)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil
}