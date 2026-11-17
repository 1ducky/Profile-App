package session

import (
	"context"
	"database/sql"
	"profile-template/internal/auth"
	"profile-template/internal/db"
)

type MYSQLService struct {
	db db.DBTX
}

func NewMySQLRepository(db db.DBTX) auth.SessionRepositoryContract {
	return &MYSQLService{db: db}
}

func (s *MYSQLService) Create(ctx context.Context, payload auth.CreateSessionPayload) error {
	query := `INSERT INTO session (id, user_id, access_token, refresh_token, revoke_at, version) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := s.db.ExecContext(ctx, query, payload.SessionID, payload.UserID, payload.AccessToken, payload.RefreshToken, payload.RevokeAt, 1)
	if err != nil {
		return err
	}
	return nil
}
func (s *MYSQLService) Rotate(ctx context.Context, payload auth.UpdateSessionPayload) error {
	query := `UPDATE session SET access_token = ?, refresh_token = ?, version = version + 1, revoke_at = ? WHERE user_id = ? AND refresh_token = ?`
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

func (s *MYSQLService) Delete(ctx context.Context, payload auth.DeleteSessionPayload) error {
	query := `DELETE FROM session WHERE id = ? AND refresh_token = ? AND user_id = ?`
	_, err := s.db.ExecContext(ctx, query, payload.SessionID, payload.RefreshToken, payload.UserID)
	if err != nil {
		return err
	}
	return nil
}
func (s *MYSQLService) GetVersion(ctx context.Context, accessToken string) (int, error) {
	query := `SELECT version FROM session WHERE access_token = ?`
	var version int
	err := s.db.QueryRowContext(ctx, query, accessToken).Scan(&version)
	if err != nil {
		return 0, err
	}
	return version, nil
}
func (s *MYSQLService) GetByAccessToken(ctx context.Context, accessToken string) (auth.SessionEntity, error) {
	query := `SELECT * FROM session WHERE access_token = ?`
	var session auth.SessionEntity
	err := s.db.QueryRowContext(ctx, query, accessToken).Scan(&session.ID, &session.UserID, &session.AccessToken, &session.RefreshToken, &session.RevokeAt, &session.Version)
	if err != nil {
		return auth.SessionEntity{}, err
	}
	return session, nil
}

func (s *MYSQLService) GetByRefreshToken(ctx context.Context, refreshToken string) (auth.SessionEntity, error) {
	query := `SELECT * FROM session WHERE refresh_token = ?`
	var session auth.SessionEntity
	err := s.db.QueryRowContext(ctx, query, refreshToken).Scan(&session.ID, &session.UserID, &session.AccessToken, &session.RefreshToken, &session.RevokeAt, &session.Version)
	if err != nil {
		return auth.SessionEntity{}, err
	}
	return session, nil
}
