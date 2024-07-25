package appMapper

import (
	appDto "github.com/illusory-server/auth-service/internal/app/app_dto"
	"github.com/illusory-server/auth-service/internal/domain/model"
)

type UserMapper struct {
}

func (u *UserMapper) ToPureUser(user *model.User, ban *model.Ban) *appDto.PureUser {
	return &appDto.PureUser{
		Id:        user.Id,
		Login:     user.Login,
		Email:     user.Email,
		IsBanned:  ban.IsBanned,
		BanReason: &ban.BanReason,
		Role:      user.Role,
	}
}

func (u *UserMapper) ToPureUserCopy(user *model.User, ban *model.Ban) appDto.PureUser {
	return appDto.PureUser{
		Id:        user.Id,
		Login:     user.Login,
		Email:     user.Email,
		IsBanned:  ban.IsBanned,
		BanReason: &ban.BanReason,
		Role:      user.Role,
	}
}
