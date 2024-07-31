package psqlActivateRepository

const (
	HasByIdQuery = `
		SELECT EXISTS(SELECT 1 FROM user_activate WHERE id = $1);
	`
	IsActivateByIdQuery = `
		SELECT EXISTS(SELECT 1 FROM user_activate WHERE is_activate = true AND id = $1);
	`
	CreateQuery = `
		INSERT INTO user_activate (id, is_activate, link, updated_at, created_at)
		VALUES ($1, $2, $3, $4, $5);
	`
	DeleteByIdQuery = `
		DELETE FROM user_activate WHERE id = $1;
	`
	GetByUserIdQuery = `
		SELECT id, is_activate, link, updated_at, created_at FROM user_activate WHERE id = $1;
	`
	UpdateQuery = `
		UPDATE user_activate SET is_activate = $2, link = $3, updated_at = $4 WHERE id = $1;
	`
	ActivateUserByIdQuery = `
		UPDATE user_activate SET is_activate = true, link = NULL, updated_at = $2 WHERE id = $1;
	`
	ActivateUserByLinkQuery = `
		UPDATE user_activate SET is_activate = true, link = NULL, updated_at = $2 WHERE link = $1;
	`
)
