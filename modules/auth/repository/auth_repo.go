package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/www-printf/wepress-core/modules/auth/domains"
)

type AuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*domains.User, error)
	GetUserByID(ctx context.Context, id string) (*domains.User, error)
	InsertUser(ctx context.Context, user *domains.User) error
	InsertKeyPair(ctx context.Context, user *domains.User, keyPair map[string]string) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) GetUserByEmail(ctx context.Context, email string) (*domains.User, error) {
	var user domains.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &user, nil
}

func (r *authRepository) GetUserByID(ctx context.Context, id string) (*domains.User, error) {
	var user domains.User
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRepository) InsertKeyPair(
	ctx context.Context, user *domains.User, keyPair map[string]string,
) error {
	err := r.db.WithContext(ctx).Model(user).Updates(keyPair).Error
	if err != nil {
		return err
	}
	if err := r.db.WithContext(ctx).First(user, user.ID).Error; err != nil {
		return err
	}
	return nil
}

func (r *authRepository) InsertUser(ctx context.Context, user *domains.User) error {
	err := r.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
