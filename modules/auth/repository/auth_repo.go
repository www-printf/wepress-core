package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/www-printf/wepress-core/modules/auth/domains"
)

type AuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*domains.User, error)
}

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{db: db}
}

func (r *AuthRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*domains.User, error) {
	var user domains.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
