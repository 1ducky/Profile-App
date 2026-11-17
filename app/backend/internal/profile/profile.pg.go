package profile

import (
	"context"
	"database/sql"
	"profile-template/internal/db"
	"strconv"
	"strings"
)

type PostgresRepository struct {
	db db.DBTX
}

func NewPostgresRepository(db db.DBTX) RepoContract {
	return &PostgresRepository{db: db}
}

// placeholders menghasilkan $start, $start+1, ... sebanyak n.
// Sebaiknya dipindah ke package internal/db agar dipakai bersama.
func placeholders(start, n int) []string {
	out := make([]string, n)
	for i := range out {
		out[i] = "$" + strconv.Itoa(start+i)
	}
	return out
}

var profileColumns = db.MakeColm(ID, USERNAME, TAG, BIO, AVATAR_URL, COVER_URL)

func scanProfile(rows *sql.Rows) (Profile, error) {
	var p Profile
	var bio, avatar, cover sql.NullString
	if err := rows.Scan(&p.ID, &p.Username, &p.Tag, &bio, &avatar, &cover); err != nil {
		return Profile{}, err
	}
	if bio.Valid {
		p.Bio = bio.String
	}
	if avatar.Valid {
		p.AvatarURL = avatar.String
	}
	if cover.Valid {
		p.CoverURL = cover.String
	}
	return p, nil
}

func (r *PostgresRepository) CreateUser(ctx context.Context, req CreateData) error {
	cols := strings.Join(db.MakeColm(ID, USERNAME, TAG), ",")
	ph := strings.Join(placeholders(1, 3), ",")
	query := "INSERT INTO " + string(TABLE) + " (" + cols + ") VALUES (" + ph + ")"
	_, err := r.db.ExecContext(ctx, query, req.ID, req.Username, req.Tag)
	return err
}

func (r *PostgresRepository) UpdateProfile(ctx context.Context, req UpdateRequest) error {
	query := "UPDATE " + string(TABLE) + " SET " +
		string(USERNAME) + " = $1, " +
		string(TAG) + " = $2, " +
		string(BIO) + " = $3 " +
		"WHERE " + string(ID) + " = $4"
	_, err := r.db.ExecContext(ctx, query, req.Username, req.Tag, req.Bio, req.ID)
	return err
}

func (r *PostgresRepository) GetUserByID(ctx context.Context, id string) (Profile, error) {
	query := "SELECT " + strings.Join(profileColumns, ",") +
		" FROM " + string(TABLE) + " WHERE " + string(ID) + " = $1"
	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return Profile{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return Profile{}, err
		}
		return Profile{}, ErrNotFound
	}
	profile, err := scanProfile(rows)
	if err != nil {
		return Profile{}, err
	}
	if err := rows.Err(); err != nil {
		return Profile{}, err
	}
	return profile, nil
}

func (r *PostgresRepository) GetUsersByIDs(ctx context.Context, ids []string) ([]Profile, error) {
	if len(ids) == 0 {
		return []Profile{}, nil
	}

	// Hapus duplikat agar query lebih ringan
	seen := make(map[string]struct{}, len(ids))
	args := make([]any, 0, len(ids))
	for _, id := range ids {
		if _, ok := seen[id]; ok {
			continue
		}
		seen[id] = struct{}{}
		args = append(args, id)
	}

	query := "SELECT " + strings.Join(profileColumns, ",") +
		" FROM " + string(TABLE) +
		" WHERE " + string(ID) + " IN (" + strings.Join(placeholders(1, len(args)), ",") + ")"

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	profiles := make([]Profile, 0, len(args))
	for rows.Next() {
		p, err := scanProfile(rows)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return profiles, nil
}

func (r *PostgresRepository) DeleteUser(ctx context.Context, id string) error {
	return nil
}

func (r *PostgresRepository) SetProfilePhoto(ctx context.Context, id string, path string) error {
	query := "UPDATE " + string(TABLE) + " SET " + string(AVATAR_URL) + " = $1 WHERE " + string(ID) + " = $2"
	_, err := r.db.ExecContext(ctx, query, path, id)
	return err
}
