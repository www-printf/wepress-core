package usecases

import (
	"context"
	"encoding/json"
	"errors"

	"golang.org/x/oauth2"

	"github.com/www-printf/wepress-core/config"
	"github.com/www-printf/wepress-core/modules/auth/domains"
	"github.com/www-printf/wepress-core/modules/auth/dto"
	"github.com/www-printf/wepress-core/utils"
)

type OauthStrategy interface {
	GenerateOauthSession() (*domains.OauthSession, error)
	ExchangeToken(*dto.OauthCallBackTransfer) (*oauth2.Token, error)
	GetUserInfo(*oauth2.Token) (*dto.OauthUserResponse, error)
}

// Github
type githubOauthStrategy struct {
	conf          *oauth2.Config
	uinfoEndpoint string
}

func NewGithubOauthStrategy(conf *config.GithubOauthConfig) OauthStrategy {
	return &githubOauthStrategy{
		conf: &oauth2.Config{
			ClientID:     conf.ClientID,
			ClientSecret: conf.ClientSecret,
			RedirectURL:  conf.RedirectURL,
			Scopes:       conf.Scopes,
			Endpoint: oauth2.Endpoint{
				AuthURL:  conf.AuthURL,
				TokenURL: conf.TokenURL,
			},
		},
		uinfoEndpoint: conf.UserInfoURL,
	}
}

func (s *githubOauthStrategy) GenerateOauthSession() (*domains.OauthSession, error) {
	state, err := utils.GenerateSecret()
	if err != nil {
		return nil, err
	}

	url := s.conf.AuthCodeURL(state, oauth2.SetAuthURLParam("prompt", "select_account"))

	return &domains.OauthSession{
		State:    state,
		URL:      url,
		Provider: "github",
	}, nil
}

func (s *githubOauthStrategy) ExchangeToken(req *dto.OauthCallBackTransfer) (*oauth2.Token, error) {
	ctx := context.Background()
	tok, err := s.conf.Exchange(ctx, req.Code)
	if err != nil {
		return nil, err
	}
	return tok, nil
}

func (s *githubOauthStrategy) GetUserInfo(tok *oauth2.Token) (*dto.OauthUserResponse, error) {
	client := s.conf.Client(context.Background(), tok)
	resp, err := client.Get(s.uinfoEndpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("failed to get user info")
	}

	var emails []dto.GithubEmail
	err = json.NewDecoder(resp.Body).Decode(&emails)
	if err != nil {
		return nil, err
	}

	return &dto.OauthUserResponse{
		Email: emails[0].Email,
	}, nil
}

// Google
type googleOauthStrategy struct {
	conf          *oauth2.Config
	uinfoEndpoint string
}

func NewGoogleOauthStrategy(conf *config.GoogleOauthConfig) OauthStrategy {
	return &googleOauthStrategy{
		conf: &oauth2.Config{
			ClientID:     conf.ClientID,
			ClientSecret: conf.ClientSecret,
			RedirectURL:  conf.RedirectURL,
			Scopes:       conf.Scopes,
			Endpoint: oauth2.Endpoint{
				AuthURL:  conf.AuthURL,
				TokenURL: conf.TokenURL,
			},
		},
		uinfoEndpoint: conf.UserInfoURL,
	}
}

func (s *googleOauthStrategy) GenerateOauthSession() (*domains.OauthSession, error) {
	return nil, nil
}

func (s *googleOauthStrategy) ExchangeToken(*dto.OauthCallBackTransfer) (*oauth2.Token, error) {
	return nil, nil
}

func (s *googleOauthStrategy) GetUserInfo(*oauth2.Token) (*dto.OauthUserResponse, error) {
	return nil, nil
}

// Facebook
type facebookOauthStrategy struct {
	conf          *oauth2.Config
	uinfoEndpoint string
}

func NewFacebookOauthStrategy(conf *config.FacebookOauthConfig) OauthStrategy {
	return &facebookOauthStrategy{
		conf: &oauth2.Config{
			ClientID:     conf.ClientID,
			ClientSecret: conf.ClientSecret,
			RedirectURL:  conf.RedirectURL,
			Scopes:       conf.Scopes,
			Endpoint: oauth2.Endpoint{
				AuthURL:  conf.AuthURL,
				TokenURL: conf.TokenURL,
			},
		},
		uinfoEndpoint: conf.UserInfoURL,
	}
}

func (s *facebookOauthStrategy) GenerateOauthSession() (*domains.OauthSession, error) {
	return nil, nil
}

func (s *facebookOauthStrategy) ExchangeToken(*dto.OauthCallBackTransfer) (*oauth2.Token, error) {
	return nil, nil
}

func (s *facebookOauthStrategy) GetUserInfo(*oauth2.Token) (*dto.OauthUserResponse, error) {
	return nil, nil
}
