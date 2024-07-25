package tokenService

import (
	"context"
)

func (s *service) DeleteByValue(ctx context.Context, value string) error {
	err := s.tokenRepository.DeleteByValue(ctx, value)
	if err != nil {
		return err
	}

	return nil
}
