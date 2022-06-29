package repository

import (
	"context"
	"strings"

	"github.com/Final-Project-Kelompok-3/authentications/internal/dto"
	"github.com/Final-Project-Kelompok-3/authentications/internal/model"

	"gorm.io/gorm"
)

type Role interface {
	FindAll(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.Role, *dto.PaginationInfo, error)
	FindByID(ctx context.Context, ID uint) (model.Role, error)
	Create(ctx context.Context, role model.Role) error
	Update(ctx context.Context, ID uint, data map[string]interface{}) error
	Delete(ctx context.Context, ID uint) error
}

type role struct {
	Db *gorm.DB
}

func NewRole(db *gorm.DB) *role {
	return &role{db}
}

func (r *role) FindAll(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.Role, *dto.PaginationInfo, error) {
	
	var (
		roles []model.Role;
		count int64;
	)
	
	query := r.Db.WithContext(ctx).Model(&model.Role{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(title) LIKE ? ", search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := dto.GetLimitOffset(paginate)

	err := query.Limit(limit).Offset(offset).Find(&roles).Error

	return roles, dto.CheckInfoPagination(paginate,count), err
}

func (r *role) FindByID(ctx context.Context, ID uint) (model.Role, error) {
	
	var role model.Role
	err := r.Db.WithContext(ctx).Model(&role).Where("id = ?", ID).First(&role).Error

	return role, err
}

func (r *role) Create(ctx context.Context, role model.Role) error {

	return r.Db.WithContext(ctx).Model(&model.Role{}).Create(&role).Error
}

func (r *role) Update(ctx context.Context, ID uint, data map[string]interface{}) error {

	err := r.Db.WithContext(ctx).Where("id = ?", ID).Model(&model.Role{}).Updates(data).Error
	return err
}

 func (r *role) Delete(ctx context.Context, ID uint) error {
	err := r.Db.WithContext(ctx).Where("id = ?", ID).Delete(&model.Role{}).Error
	return err
 }