package model

import (
	"github.com/illusory-server/auth-service/internal/domain"
	"time"
)

type (
	User struct {
		Id        domain.Id `json:"id"`
		Login     string    `json:"login"`
		Email     string    `json:"email"`
		Role      string    `json:"role"`
		Password  string    `json:"-"`
		UpdatedAt time.Time `json:"updated_at"`
		CreatedAt time.Time `json:"created_at"`
	}

	Token struct {
		Id        domain.Id `json:"id"`
		Value     string    `json:"value"`
		UpdatedAt time.Time `json:"updated_at"`
		CreatedAt time.Time `json:"created_at"`
	}

	Activate struct {
		Id          domain.Id `json:"id"`
		IsActivated bool      `json:"is_activated"`
		Link        string    `json:"link"`
		UpdatedAt   time.Time `json:"updated_at"`
		CreatedAt   time.Time `json:"created_at"`
	}

	Ban struct {
		Id        domain.Id `json:"id"`
		IsBanned  bool      `json:"is_banned"`
		BanReason string    `json:"ban_reason"`
		UpdatedAt time.Time `json:"updated_at"`
		CreatedAt time.Time `json:"created_at"`
	}
)
