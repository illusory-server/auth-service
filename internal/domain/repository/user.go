package repository

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/internal/domain/query"
)

type UserRepository interface {
	Create(context.Context, *model.User) (*model.User, error)

	GetById(context.Context, domain.Id) (*model.User, error)
	GetByLogin(context.Context, string) (*model.User, error)
	GetByQuery(context.Context, *query.PaginationQuery) ([]*model.User, error)
	GetPageCount(ctx context.Context, limit uint) (uint, error)

	UpdateById(context.Context, *model.User) (*model.User, error)
	UpdateRoleById(ctx context.Context, id string, role string) (*model.User, error)
	UpdatePasswordById(ctx context.Context, id string, password string) (*model.User, error)

	DeleteById(context.Context, domain.Id) error

	HasById(context.Context, domain.Id) (bool, error)
	HasByLogin(context.Context, string) (bool, error)
	HasByEmail(context.Context, string) (bool, error)
	CheckUserRole(ctx context.Context, id, rol string) (bool, error)
}
