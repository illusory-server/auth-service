package repository

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain/model"
)

type ActivateRepository interface {
	Create(ctx context.Context, userId string) (*model.Activate, error)

	GetByUserId(ctx context.Context, userId string) (*model.Activate, error)

	DeleteById(ctx context.Context, userId string) error

	Update(ctx context.Context, activate *model.Activate) (*model.Activate, error)
	ActivateUserById(ctx context.Context, userId string) (*model.Activate, error)
	ActivateUserByLink(ctx context.Context, link string) (*model.Activate, error)

	IsActivateById(ctx context.Context, userId string) (bool, error)
	HasById(ctx context.Context, userId string) (bool, error)
}
