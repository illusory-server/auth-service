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
	HasById = `
		SELECT EXISTS(SELECT 1 FROM jwt_tokens WHERE id = $1);
	`
	HasByValue = `
		SELECT EXISTS(SELECT 1 FROM jwt_tokens WHERE value = $1);
	`
)
