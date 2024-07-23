package repository

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/model"
)

type BanRepository interface {
	Create(ctx context.Context, ban *model.Ban) (*model.Ban, error)

	GetBanReasonById(ctx context.Context, id domain.Id) (string, error)

	DeleteById(context.Context, domain.Id) error

	BanById(ctx context.Context, id domain.Id, reason string) (*model.Ban, error)
	UnbanById(context.Context, domain.Id) (*model.Ban, error)

	HasById(context.Context, domain.Id) (bool, error)
	IsBannedById(ctx context.Context, id domain.Id) (bool, error)
}
