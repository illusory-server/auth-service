package tokenService

import (
	"context"
)

func (s *service) HasByValue(ctx context.Context, refreshToken string) (bool, error) {
	has, err := s.tokenRepository.HasByValue(ctx, refreshToken)
	if err != nil {
		return false, err
	}

	return has, nil
}
