package repository

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain/model"
)

type TokenRepository interface {
	Create(ctx context.Context, id, token string) (*model.Token, error)

	GetById(ctx context.Context, id string) (*model.Token, error)

	DeleteById(ctx context.Context, id string) error
	DeleteByValue(ctx context.Context, value string) error

	UpdateById(ctx context.Context, id, token string) (*model.Token, error)

	HasById(ctx context.Context, id string) (bool, error)
	HasByValue(ctx context.Context, token string) (bool, error)

	Save(ctx context.Context, id, token string) (*model.Token, error)
}
