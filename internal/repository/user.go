package repository

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/Final-Project-Kelompok-3/authentications/internal/dto"
	"github.com/Final-Project-Kelompok-3/authentications/internal/middleware"
	"github.com/Final-Project-Kelompok-3/authentications/internal/model"

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
	Login(ctx context.Context, email, password string) (string, error)
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

	err := u.Db.WithContext(ctx).Where("id = ?", ID).Model(&model.User{}).Updates(data).Error
	return err
}

func (u *user) Delete(ctx context.Context, ID uint) error {
	err := u.Db.WithContext(ctx).Where("id = ?", ID).Delete(&model.User{}).Error
	return err
}

func (u *user) Login(ctx context.Context, email, password string) (string, error) {
	
	user, err := u.FindByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("email and password don't match")
	}

	return middleware.CreateToken(strconv.FormatUint(uint64(user.ID), 10))
}