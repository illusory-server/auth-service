package psqlBanRepository

const (
	CreateQuery = `
		INSERT INTO user_ban (id, is_banned, ban_reason, updated_at, created_at)
		VALUES ($1, $2, $3, $4, $5);
	`
	HasByIdQuery = `
		SELECT EXISTS(SELECT 1 FROM user_ban WHERE id = $1);
	`
	IsBannedByIdQuery = `
		SELECT EXISTS(SELECT 1 FROM user_ban WHERE id = $1 AND is_banned = true);
	`
	DeleteByIdQuery = `
		DELETE FROM user_ban WHERE id = $1;
	`
	GetBanReasonByIdQuery = `
		SELECT id, is_banned, ban_reason, updated_at, created_at FROM user_ban WHERE id = $1;
	`
	BanByIdQuery = `
		UPDATE user_ban SET is_banned = true, ban_reason = $2, updated_at = $3 WHERE id = $1
		RETURNING id, is_banned, ban_reason, updated_at, created_at;
	`
	UnbanByIdQuery = `
		UPDATE user_ban SET is_banned = false, ban_reason = NULL, updated_at = $2 WHERE id = $1
		RETURNING id, is_banned, ban_reason, updated_at, created_at;
	`
)
