package config

type OauthConfig struct {
	Providers []string `env:"OAUTH_PROVIDERS" envSeparator:","`
	Github    GithubOauthConfig
	Google    GoogleOauthConfig
	Facebook  FacebookOauthConfig
}

type GoogleOauthConfig struct {
	ClientID     string   `env:"GOOGLE_CLIENT_ID"`
	ClientSecret string   `env:"GOOGLE_CLIENT_SECRET"`
	RedirectURL  string   `env:"GOOGLE_REDIRECT_URL"`
	AuthURL      string   `env:"GOOGLE_AUTH_URL"`
	TokenURL     string   `env:"GOOGLE_TOKEN_URL"`
	UserInfoURL  string   `env:"GOOGLE_USER_INFO_URL"`
	Scopes       []string `env:"GOOGLE_SCOPE" envSeparator:","`
}

type FacebookOauthConfig struct {
	ClientID     string   `env:"FACEBOOK_CLIENT_ID"`
	ClientSecret string   `env:"FACEBOOK_CLIENT_SECRET"`
	RedirectURL  string   `env:"FACEBOOK_REDIRECT_URL"`
	AuthURL      string   `env:"FACEBOOK_AUTH_URL"`
	TokenURL     string   `env:"FACEBOOK_TOKEN_URL"`
	UserInfoURL  string   `env:"FACEBOOK_USER_INFO_URL"`
	Scopes       []string `env:"FACEBOOK_SCOPE" envSeparator:","`
}

type GithubOauthConfig struct {
	ClientID     string   `env:"GITHUB_CLIENT_ID"`
	ClientSecret string   `env:"GITHUB_CLIENT_SECRET"`
	RedirectURL  string   `env:"GITHUB_REDIRECT_URL"`
	AuthURL      string   `env:"GITHUB_AUTH_URL"`
	TokenURL     string   `env:"GITHUB_TOKEN_URL"`
	UserInfoURL  string   `env:"GITHUB_USER_INFO_URL"`
	Scopes       []string `env:"GITHUB_SCOPE" envSeparator:","`
}
