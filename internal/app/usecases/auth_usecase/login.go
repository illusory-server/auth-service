package authUseCase

import (
	"context"
	appDto "github.com/illusory-server/auth-service/internal/app/app_dto"
)

const LoginOrPasswordIncorrect = "incorrect login or password"

func (u *useCase) Login(ctx context.Context, data *appDto.LoginData) (*AuthResult, error) {
	//has, err := u.userRepository.HasUserByLogin(ctx, data.Login)
	//if err != nil {
	//	return nil, err
	//}
	//if !has {
	//	u.log.ErrorContext(ctx, "login error", "login_input_data", data)
	//	return nil, domain.NewErr(domain.ErrForbiddenCode, LoginOrPasswordIncorrect)
	//}
	//
	//userDb, err := u.userRepository.GetByLogin(ctx, data.Login)
	//if err != nil {
	//	return nil, err
	//}
	//isEqualPassword := bcrypt.CompareHashAndPassword([]byte(userDb.Password), []byte(data.Password))
	//if isEqualPassword != nil {
	//	u.log.ErrorContext(ctx, "login error", "login_input_data", data)
	//	return nil, domain.NewErr(domain.ErrForbiddenCode, LoginOrPasswordIncorrect)
	//}
	//tokens, err := u.tokenService.Generate(ctx, tokenService.JwtUserData{Id: userDb.Id, Role: userDb.Role})
	//if err != nil {
	//	return nil, err
	//}
	//_, err = u.tokenService.Save(ctx, appDto.SaveTokenServiceDto{Id: userDb.Id, RefreshToken: tokens.RefreshToken})
	//if err != nil {
	//	return nil, err
	//}
	//
	//mapper := &appMapper.UserMapper{}
	//pureUser := mapper.ToPureUser(userDb)
	//return &AuthResult{
	//	User:   pureUser,
	//	Tokens: tokens,
	//}, nil
	return nil, nil
}
