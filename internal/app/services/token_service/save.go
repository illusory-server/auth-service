package tokenService

import (
	"context"
	appDto "github.com/illusory-server/auth-service/internal/app/app_dto"
	"github.com/illusory-server/auth-service/internal/domain/model"
)

func (s *service) Save(ctx context.Context, data appDto.SaveTokenServiceDto) (*model.Token, error) {
	save, err := s.tokenRepository.Save(ctx, data.Id, data.RefreshToken)
	if err != nil {
		return nil, err
	}
	return save, nil
}
