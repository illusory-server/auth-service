package psqlUserRepository

const (
	CreateQuery = `
		INSERT INTO users (id, login, email, role, password, updated_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7);
	`
	DeleteByIdQuery = `
		DELETE FROM users WHERE id = $1;
	`
	UpdateByIdQuery = `
		UPDATE users SET login = $2, email = $3, role = $4, updated_at = $5 WHERE id = $1;
	`
	UpdateRoleByIdQuery = `
		UPDATE users SET role = $2, updated_at = $3 WHERE id = $1
		RETURNING (id, login, email, role, password, updated_at, created_at);
	`
	UpdatePasswordByIdQuery = `
		UPDATE users SET password = $2, updated_at = $3 WHERE id = $1
		RETURNING (id, login, email, role, password, updated_at, created_at);
	`
	HasByIdQuery = `
		SELECT EXISTS(SELECT 1 FROM users WHERE id = $1); 
	`
	HasByLoginQuery = `
		SELECT EXISTS(SELECT 1 FROM users WHERE login = $1);
	`
	HasByEmailQuery = `
		SELECT EXISTS(SELECT 1 FROM users WHERE email = $1);
	`
	CheckUserRoleQuery = `
		SELECT EXISTS(SELECT 1 FROM users WHERE id = $1 AND role = $2);
	`
	GetByIdQuery = `
		SELECT id, login, email, role, password, updated_at, created_at FROM users WHERE id = $1;
	`
	GetByLoginQuery = `
		SELECT id, login, email, role, password, updated_at, created_at FROM users WHERE login = $1;
	`
	GetByQuerySelect = `SELECT id, login, email, role, password, updated_at, created_at FROM users `
	GetByQueryLimit  = ` LIMIT $1 OFFSET $2;`
)
