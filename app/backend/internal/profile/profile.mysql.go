package profile

import (
	"context"
	"database/sql"
	"profile-template/internal/db"
	"strings"
)

type MySQLRepository struct {
	db db.DBTX
	// errMapper mapper.ErrorMapping
}

func NewMySQLRepository(db db.DBTX) RepoContract {
	return &MySQLRepository{db: db}
}

func (r *MySQLRepository) CreateUser(ctx context.Context, req CreateData) error {
	var sql []string
	colloms := "(" + strings.Join(db.MakeColm(ID, USERNAME, TAG), ",") + ")"
	command := "INSERT INTO"
	values := "(" + strings.Join(db.MakePlaceHolder(3), ",") + ")"
	sql = append(sql, command, string(TABLE), colloms, "VALUES", values)
	query := strings.Join(sql, " ")
	_, err := r.db.ExecContext(ctx, query, req.ID, req.Username, req.Tag)
	if err != nil {
		return err
	}

	return nil
}

func (r *MySQLRepository) UpdateProfile(ctx context.Context, req UpdateRequest) error {
	var sqlQuery []string

	command := "UPDATE"
	setStr := "SET"
	where := "WHERE"
	sqlQuery = append(sqlQuery, command, string(TABLE), setStr, strings.Join(db.MakeSet([]db.Collom{USERNAME, TAG, BIO}), ","), where, string(ID), "=?")
	query := strings.Join(sqlQuery, " ")
	_, err := r.db.ExecContext(ctx, query, req.Username, req.Tag, req.Bio, req.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *MySQLRepository) GetUserByID(ctx context.Context, id string) (Profile, error) {
	var sqlQuery []string
	command := "SELECT"
	colloms := strings.Join(db.MakeColm(ID, USERNAME, TAG, BIO, AVATAR_URL, COVER_URL), ",")
	sqlQuery = append(sqlQuery, command, colloms, "FROM", string(TABLE), "WHERE", string(ID), "=?")
	query := strings.Join(sqlQuery, " ")
	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return Profile{}, err
	}
	defer rows.Close()
	var profile Profile
	var bio, avatar, cover sql.NullString
	found := false
	for rows.Next() {
		if err := rows.Scan(&profile.ID, &profile.Username, &profile.Tag, &bio, &avatar, &cover); err != nil {
			return Profile{}, err
		}
		found = true
	}
	if !found {
		return Profile{}, ErrNotFound
	}
	if err := rows.Err(); err != nil {
		return Profile{}, err
	}
	if bio.Valid {
		profile.Bio = bio.String
	}
	if avatar.Valid {
		profile.AvatarURL = avatar.String
	}
	if cover.Valid {
		profile.CoverURL = cover.String
	}
	return profile, nil

}
func (r *MySQLRepository) GetUsersByIDs(ctx context.Context, ids []string) ([]Profile, error) {
	if len(ids) == 0 {
		return []Profile{}, nil
	}

	// Hapus duplikat agar query lebih ringan
	seen := make(map[string]struct{}, len(ids))
	uniqueIDs := make([]string, 0, len(ids))
	for _, id := range ids {
		if _, ok := seen[id]; ok {
			continue
		}
		seen[id] = struct{}{}
		uniqueIDs = append(uniqueIDs, id)
	}

	// Buat placeholder "?,?,?" dan args
	placeholders := strings.TrimSuffix(strings.Repeat("?,", len(uniqueIDs)), ",")
	args := make([]any, len(uniqueIDs))
	for i, id := range uniqueIDs {
		args[i] = id
	}

	colloms := strings.Join(db.MakeColm(ID, USERNAME, TAG, BIO, AVATAR_URL, COVER_URL), ",")
	query := strings.Join([]string{
		"SELECT", colloms,
		"FROM", string(TABLE),
		"WHERE", string(ID), "IN (" + placeholders + ")",
	}, " ")

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	profiles := make([]Profile, 0, len(uniqueIDs))
	for rows.Next() {
		var profile Profile
		var bio, avatar, cover sql.NullString

		if err := rows.Scan(&profile.ID, &profile.Username, &profile.Tag, &bio, &avatar, &cover); err != nil {
			return nil, err
		}
		if bio.Valid {
			profile.Bio = bio.String
		}
		if avatar.Valid {
			profile.AvatarURL = avatar.String
		}
		if cover.Valid {
			profile.CoverURL = cover.String
		}
		profiles = append(profiles, profile)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return profiles, nil
}
func (r *MySQLRepository) DeleteUser(ctx context.Context, id string) error {
	return nil
}

func (r *MySQLRepository) SetProfilePhoto(ctx context.Context, id string, path string) error {
	var sqlQuery []string

	command := "UPDATE"
	setStr := "SET"
	where := "WHERE"
	sqlQuery = append(sqlQuery, command, string(TABLE), setStr, strings.Join(db.MakeSet([]db.Collom{AVATAR_URL}), ","), where, string(ID), "=?")
	query := strings.Join(sqlQuery, " ")
	_, err := r.db.ExecContext(ctx, query, path, id)
	if err != nil {
		return err
	}

	return nil
}
