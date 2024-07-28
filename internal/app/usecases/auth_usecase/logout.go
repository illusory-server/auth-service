package authUseCase

import (
	"context"
)

func (u *useCase) Logout(ctx context.Context, refreshToken string) error {
	err := u.tokenRepository.DeleteByValue(ctx, refreshToken)
	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Str("refresh_token", refreshToken).
			Msg("failed to delete refresh token")

		return err
	}
	return nil
}
