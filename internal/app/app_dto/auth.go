package appDto

import "github.com/illusory-server/auth-service/internal/domain"

type (
	PureUser struct {
		Id        domain.Id
		Login     string
		Email     string
		IsBanned  bool
		BanReason *string
		Role      string
	}

	RegistrationData struct {
		Login    string
		Password string
		Email    string
	}

	LoginData struct {
		Login    string
		Password string
	}

	SaveTokenServiceDto struct {
		Id           string
		RefreshToken string
	}
)
