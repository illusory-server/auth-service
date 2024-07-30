package psqlTokenRepository

const (
	CreateQuery = `
		INSERT INTO jwt_tokens (id, value, updated_at, created_at)
		VALUES ($1, $2, $3, $4);
	`
	UpdateQuery = `
		UPDATE jwt_tokens SET value = $2, updated_at = $3 WHERE id = $1
		RETURNING id, value, updated_at, created_at;
	`
	HasByIdQuery = `
		SELECT EXISTS(SELECT 1 FROM jwt_tokens WHERE id = $1);
	`
	HasByValueQuery = `
		SELECT EXISTS(SELECT 1 FROM jwt_tokens WHERE value = $1);
	`
	DeleteByValueQuery = `
		DELETE FROM jwt_tokens WHERE value = $1;
	`
	DeleteByIdQuery = `
		DELETE FROM jwt_tokens WHERE id = $1;
	`
	GetByIdQuery = `
		SELECT id, value, updated_at, created_at FROM jwt_tokens WHERE id = $1;
	`
)
