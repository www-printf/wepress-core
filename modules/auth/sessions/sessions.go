package sessions

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/www-printf/wepress-core/modules/auth/domains"
)

type SessionStorage interface {
	SetUserSession(ctx context.Context, uid string, session *domains.SessionUser, exp time.Duration) error
	GetUserSession(ctx context.Context, uid string) (*domains.SessionUser, error)
	SetOauthSession(ctx context.Context, sess *domains.OauthSession, exp time.Duration) error
	GetOauthSession(ctx context.Context, state string) (*domains.OauthSession, error)
}

type sessionStorage struct {
	redisClient *redis.Client
}

func NewSessionStorage(redisClient *redis.Client) SessionStorage {
	return &sessionStorage{redisClient: redisClient}
}

func (s *sessionStorage) SetUserSession(ctx context.Context, uid string, session *domains.SessionUser, exp time.Duration) error {
	fields := map[string]interface{}{
		"FullName":  session.FullName,
		"Email":     session.Email,
		"CreatedAt": session.CreatedAt.Format(time.RFC3339Nano),
		"PubKey":    session.PubKey,
	}

	err := s.redisClient.HSet(ctx, uid, fields).Err()
	if err != nil {
		return err
	}

	return s.redisClient.Expire(ctx, uid, exp).Err()
}

func (s *sessionStorage) GetUserSession(ctx context.Context, uid string) (*domains.SessionUser, error) {
	user := &domains.SessionUser{}
	err := s.redisClient.HGetAll(ctx, uid).Scan(user)
	return user, err
}

func (s *sessionStorage) SetOauthSession(ctx context.Context, sess *domains.OauthSession, exp time.Duration) error {
	fields := map[string]interface{}{
		"Verifier": sess.Verifier,
		"Provider": sess.Provider,
	}

	err := s.redisClient.HSet(ctx, sess.State, fields).Err()
	if err != nil {
		return err
	}

	return s.redisClient.Expire(ctx, sess.State, exp).Err()
}

func (s *sessionStorage) GetOauthSession(ctx context.Context, state string) (*domains.OauthSession, error) {
	oauth := &domains.OauthSession{}
	err := s.redisClient.HGetAll(ctx, state).Scan(oauth)
	return oauth, err
}
