-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(50) PRIMARY KEY,
    login VARCHAR(512) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    role VARCHAR(128) NOT NULL ,
    password VARCHAR(255) NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS jwt_tokens (
    id VARCHAR(50) PRIMARY KEY,
    value VARCHAR(512) UNIQUE NOT NULL,
    updated_at timestamp NOT NULL,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS user_activate (
    id VARCHAR(50) PRIMARY KEY,
    is_activate BOOLEAN NOT NULL,
    link VARCHAR(1024) NOT NULL,
    updated_at timestamp NOT NULL,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS user_ban (
    id VARCHAR(50) PRIMARY KEY,
    is_banned BOOLEAN NOT NULL,
    ban_reason VARCHAR(1024),
    updated_at timestamp NOT NULL,
    created_at TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
