package authUseCase

import (
	"context"
	appDto "github.com/illusory-server/auth-service/internal/app/app_dto"
)

func (u *useCase) Registration(ctx context.Context, data *appDto.RegistrationData) (*AuthResult, error) {
	//user, err := u.userService.Create(ctx, data)
	//if err != nil {
	//	return nil, err
	//}
	//
	//tokens, err := u.tokenService.Generate(ctx, tokenService.JwtUserData{
	//	Id:   user.Id,
	//	Role: user.Role,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//_, err = u.tokenService.Save(ctx, appDto.SaveTokenServiceDto{Id: user.Id, RefreshToken: tokens.RefreshToken})
	//if err != nil {
	//	return nil, err
	//}
	//
	//mapper := appMapper.UserMapper{}
	//pureUser := mapper.ToPureUser(user)
	//return &AuthResult{User: pureUser, Tokens: tokens}, nil
	return nil, nil
}
