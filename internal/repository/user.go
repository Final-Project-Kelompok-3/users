package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/Final-Project-Kelompok-3/users/internal/dto"
	"github.com/Final-Project-Kelompok-3/users/internal/model"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type User interface {
	FindAll(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.User, *dto.PaginationInfo, error)
	FindByID(ctx context.Context, ID uint) (model.User, error)
	FindByEmail(ctx context.Context, email string) (model.User, error)
	Create(ctx context.Context, user model.User) error
	Update(ctx context.Context, ID uint, data map[string]interface{}) error
	Delete(ctx context.Context, ID uint) error
}

type user struct {
	Db *gorm.DB
}

func NewUser(db *gorm.DB) *user {
	return &user{db}
}

func (u *user) FindAll(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.User, *dto.PaginationInfo, error) {
	
	var (
		users []model.User;
		count int64;
	)
	
	query := u.Db.WithContext(ctx).Model(&model.User{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("LOWER(name) LIKE ? OR LOWER(email) LIKE ? ", search, search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := dto.GetLimitOffset(paginate)

	err := query.Limit(limit).Offset(offset).Find(&users).Error

	return users, dto.CheckInfoPagination(paginate,count), err
}

func (u *user) FindByID(ctx context.Context, ID uint) (model.User, error) {
	
	var user model.User
	err := u.Db.WithContext(ctx).Model(&user).Where("id = ?", ID).First(&user).Error

	return user, err
}

func (u *user) FindByEmail(ctx context.Context, email string) (model.User, error) {
	
	var user model.User
	err := u.Db.WithContext(ctx).Model(&user).Where("email = ?", email).First(&user).Error

	return user, err
}

func (u *user) Create(ctx context.Context, user model.User) error {

	return u.Db.WithContext(ctx).Model(&model.User{}).Create(&user).Error
}

func (u *user) Update(ctx context.Context, ID uint, data map[string]interface{}) error {
	
	if data["password"] != nil {
		pss := fmt.Sprintf("%v", data["password"])
		bytes, _ := bcrypt.GenerateFromPassword([]byte(pss), bcrypt.DefaultCost)
		data["password"] = string(bytes)
	}

	err := u.Db.WithContext(ctx).Where("id = ?", ID).Model(&model.User{}).Updates(data).Error
	return err
}

func (u *user) Delete(ctx context.Context, ID uint) error {
	
	err := u.Db.WithContext(ctx).Where("id = ?", ID).Delete(&model.User{}).Error
	return err
}