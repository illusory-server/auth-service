package appDto

import "github.com/illusory-server/auth-service/internal/domain"

type (
	PureUser struct {
		Id         domain.Id
		Login      string
		Email      string
		IsBanned   bool
		BanReason  *string
		IsActivate bool
		Role       string
	}

	LoginData struct {
		Login    string
		Password string
	}

	JwtTokens struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
)
