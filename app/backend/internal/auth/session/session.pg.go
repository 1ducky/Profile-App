package session

import (
	"context"
	"database/sql"
	"profile-template/internal/auth"
	"profile-template/internal/db"
)

type PGSQLService struct {
	db db.DBTX
}

func NewPGSQLRepository(db db.DBTX) auth.SessionRepositoryContract {
	return &PGSQLService{db: db}
}

func (s *PGSQLService) Create(ctx context.Context, payload auth.CreateSessionPayload) error {
	query := `INSERT INTO session (id, user_id, access_token, refresh_token, revoke_at, version) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.ExecContext(ctx, query, payload.SessionID, payload.UserID, payload.AccessToken, payload.RefreshToken, payload.RevokeAt, 1)
	if err != nil {
		return err
	}
	return nil
}
func (s *PGSQLService) Rotate(ctx context.Context, payload auth.UpdateSessionPayload) error {
	query := `UPDATE session SET access_token = $1, refresh_token = $2, version = version + 1, revoke_at = $3 WHERE user_id = $4 AND refresh_token = $5`
	res, err := s.db.ExecContext(ctx, query, payload.AccessToken, payload.RefreshToken, payload.RevokeAt, payload.UserID, payload.OldrefreshToken)
	if err != nil {
		return err
	}
	effected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if effected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (s *PGSQLService) Delete(ctx context.Context, payload auth.DeleteSessionPayload) error {
	query := `DELETE FROM session WHERE id = $1 AND refresh_token = $2 AND user_id = $3`
	_, err := s.db.ExecContext(ctx, query, payload.SessionID, payload.RefreshToken, payload.UserID)
	if err != nil {
		return err
	}
	return nil
}
func (s *PGSQLService) GetVersion(ctx context.Context, accessToken string) (int, error) {
	query := `SELECT version FROM session WHERE access_token = $1`
	var version int
	err := s.db.QueryRowContext(ctx, query, accessToken).Scan(&version)
	if err != nil {
		return 0, err
	}
	return version, nil
}
func (s *PGSQLService) GetByAccessToken(ctx context.Context, accessToken string) (auth.SessionEntity, error) {
	query := `SELECT * FROM session WHERE access_token = $1`
	var session auth.SessionEntity
	err := s.db.QueryRowContext(ctx, query, accessToken).Scan(&session.ID, &session.UserID, &session.AccessToken, &session.RefreshToken, &session.RevokeAt, &session.Version)
	if err != nil {
		return auth.SessionEntity{}, err
	}
	return session, nil
}

func (s *PGSQLService) GetByRefreshToken(ctx context.Context, refreshToken string) (auth.SessionEntity, error) {
	query := `SELECT * FROM session WHERE refresh_token = $1`
	var session auth.SessionEntity
	err := s.db.QueryRowContext(ctx, query, refreshToken).Scan(&session.ID, &session.UserID, &session.AccessToken, &session.RefreshToken, &session.RevokeAt, &session.Version)
	if err != nil {
		return auth.SessionEntity{}, err
	}
	return session, nil
}
