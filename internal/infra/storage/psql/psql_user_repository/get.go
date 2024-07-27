package psqlUserRepository

import (
	"context"
	"fmt"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/internal/domain/query"
	"github.com/illusory-server/auth-service/pkg/eerr"
	"strings"
)

func (u *userRepository) GetById(ctx context.Context, id domain.Id) (*model.User, error) {
	db := u.getQuery(ctx)

	user := &model.User{}
	err := db.QueryRow(ctx, GetByIdQuery, id).Scan(
		&user.Id,
		&user.Login,
		&user.Email,
		&user.Role,
		&user.Password,
		&user.UpdatedAt,
		&user.CreatedAt,
	)
	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Str("id", string(id)).
			Msg("get user by id query failed")

		return nil, eerr.Wrap(err, "[userRepository.GetById] db.QueryRow")
	}

	return user, nil
}

func (u *userRepository) GetByLogin(ctx context.Context, login string) (*model.User, error) {
	db := u.getQuery(ctx)

	user := &model.User{}
	err := db.QueryRow(ctx, GetByLoginQuery, login).Scan(
		&user.Id,
		&user.Login,
		&user.Email,
		&user.Role,
		&user.Password,
		&user.UpdatedAt,
		&user.CreatedAt,
	)
	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Str("login", login).
			Msg("get user by login query failed")

		return nil, eerr.Wrap(err, "[userRepository.GetByLogin] db.QueryRow")
	}

	return user, nil
}

func (u *userRepository) GetByQuery(ctx context.Context, query *query.PaginationQuery) ([]*model.User, error) {
	db := u.getQuery(ctx)
	offset := query.Limit * (query.Page - 1)
	queryStr := strings.Builder{}
	queryStr.WriteString(GetByQuerySelect)
	queryStr.WriteString("ORDER BY ")
	queryStr.WriteString(query.SortBy)
	queryStr.WriteString(" ")
	queryStr.WriteString(string(query.SortOrder))
	queryStr.WriteString(GetByQueryLimit)

	rows, err := db.Query(ctx, queryStr.String(), query.Limit, offset)
	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Interface("query", query).
			Msg("get user by query failed")

		return nil, eerr.Wrap(err, "[userRepository.GetByQuery] db.Query")
	}
	defer rows.Close()

	users := make([]*model.User, 0, query.Limit)
	for rows.Next() {
		user := &model.User{}
		err := rows.Scan(
			&user.Id,
			&user.Login,
			&user.Email,
			&user.Role,
			&user.Password,
			&user.UpdatedAt,
			&user.CreatedAt,
		)
		if err != nil {
			u.log.Error(ctx).
				Err(err).
				Interface("query", query).
				Msg("get user by query scan failed")

			return nil, eerr.Wrap(err, "[userRepository.GetByQuery] rows.Scan")
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		u.log.Error(ctx).
			Err(err).
			Interface("query", query).
			Msg("get user by query rows failed")

		return nil, eerr.Wrap(err, "[userRepository.GetByQuery] rows.Err")
	}
	return users, nil
}

func (u *userRepository) GetPageCount(ctx context.Context, limit uint) (uint, error) {
	db := u.getQuery(ctx)

	var pageCount uint
	err := db.QueryRow(ctx, `SELECT COUNT(*) / $1 FROM users`, limit).Scan(&pageCount)
	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Str("limit", fmt.Sprint(limit)).
			Msg("get users page count query failed")

		return 0, eerr.Wrap(err, "[userRepository.GetPageCount] db.QueryRow")
	}

	return pageCount, nil
}
