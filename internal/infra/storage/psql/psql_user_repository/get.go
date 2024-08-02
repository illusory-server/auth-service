package psqlUserRepository

import (
	"context"
	"github.com/OddEer0/Eer0/eerror"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/internal/domain/query"
	"github.com/illusory-server/auth-service/pkg/etrace"
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
		tr := traceUserRepository.OfName("GetById").
			OfCauseMethod("db.QueryRow").
			OfCauseParams(etrace.FuncParams{
				"id": id,
			})
		return nil, eerror.
			Err(err).
			Code(eerror.ErrInternal).
			Msg(eerror.MsgInternal).
			Stack(tr).
			Err()
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
		tr := traceUserRepository.OfName("GetByLogin").
			OfCauseMethod("db.QueryRow").
			OfCauseParams(etrace.FuncParams{
				"login": login,
			})

		return nil, eerror.
			Err(err).
			Code(eerror.ErrInternal).
			Msg(eerror.MsgInternal).
			Stack(tr).
			Err()
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
		tr := traceUserRepository.OfName("GetByQuery").
			OfCauseMethod("db.Query").
			OfParams(etrace.FuncParams{
				"query": query,
			})
		return nil, eerror.
			Err(err).
			Code(eerror.ErrInternal).
			Msg(eerror.MsgInternal).
			Stack(tr).
			Err()
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
			tr := traceUserRepository.OfName("GetByQuery").
				OfCauseMethod("rows.Scan").
				OfParams(etrace.FuncParams{
					"query": query,
				})
			return nil, eerror.
				Err(err).
				Code(eerror.ErrInternal).
				Msg(eerror.MsgInternal).
				Stack(tr).
				Err()
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		tr := traceUserRepository.OfName("GetByQuery").
			OfCauseMethod("rows.Err").
			OfParams(etrace.FuncParams{
				"query": query,
			})
		return nil, eerror.
			Err(err).
			Code(eerror.ErrInternal).
			Msg(eerror.MsgInternal).
			Stack(tr).
			Err()
	}
	return users, nil
}

func (u *userRepository) GetPageCount(ctx context.Context, limit uint) (uint, error) {
	db := u.getQuery(ctx)

	var pageCount uint
	err := db.QueryRow(ctx, `SELECT COUNT(*) / $1 FROM users`, limit).Scan(&pageCount)
	if err != nil {
		tr := traceUserRepository.OfName("GetPageCount").
			OfCauseMethod("db.QueryRow").
			OfCauseParams(etrace.FuncParams{
				"limit": limit,
			})
		return 0, eerror.
			Err(err).
			Code(eerror.ErrInternal).
			Msg(eerror.MsgInternal).
			Stack(tr).
			Err()
	}

	return pageCount, nil
}
