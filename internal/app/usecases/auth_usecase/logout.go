package authUseCase

import (
	"context"
)

func (u *useCase) Logout(ctx context.Context, refreshToken string) error {
	err := u.tokenService.DeleteByValue(ctx, refreshToken)
	if err != nil {
		return err
	}
	return nil
}
