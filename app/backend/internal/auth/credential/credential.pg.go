package credential

import (
	"context"
	"profile-template/internal/auth"
	"profile-template/internal/db"
)

type PGRepository struct {
	DB db.DBTX
}

func NewPGRepository(db db.DBTX) auth.CredentialRepositoryContract {
	return &PGRepository{DB: db}
}

// func (r *PGRepository) FindByEmail(ctx context.Context, email string) (AuthEntity, error) {
// 	var user AuthEntity
// 	query := `SELECT id,username, role, version FROM users WHERE email = $1`
// 	row := r.DB.QueryRowContext(ctx, query, email)
// 	if err := row.Scan(&user.ID, &user.Username, &user.Role, &user.Version); err != nil {
// 		return AuthEntity{}, err
// 	}
// 	return user, nil
// }

func (r *PGRepository) Authentication(ctx context.Context, email string) (auth.AuthLogin, error) {
	var credential auth.AuthLogin
	query := `SELECT user_id,email, password,role FROM credential WHERE email = $1`
	row := r.DB.QueryRowContext(ctx, query, email)
	if err := row.Scan(&credential.UserID, &credential.Email, &credential.Password, &credential.Role); err != nil {
		return auth.AuthLogin{}, err
	}
	return credential, nil
}

func (r *PGRepository) Registration(ctx context.Context, entity auth.RegisterPayload) error {
	query := `INSERT INTO credential (email,password,user_id,id) VALUES ($1,$2,$3,$4)`
	_, err := r.DB.ExecContext(ctx, query, entity.Email, entity.Password, entity.UserID, entity.CredentialID)
	if err != nil {
		return err
	}
	return nil
}

func (r *PGRepository) GetRoleByUserId(ctx context.Context, userId string) (auth.Role, error) {
	var role auth.Role
	query := `SELECT role FROM credential WHERE user_id = $1`
	row := r.DB.QueryRowContext(ctx, query, userId)
	if err := row.Scan(&role); err != nil {
		return auth.Role("invalid"), err
	}
	return role, nil
}
