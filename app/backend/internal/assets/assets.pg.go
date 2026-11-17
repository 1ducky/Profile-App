package assets

import (
	"context"
	"database/sql"
	"profile-template/internal/db"
	"strconv"
	"strings"
)

type PostgresRepo struct {
	DB db.DBTX
}

func NewPostgresRepo(db db.DBTX) AssetRepository {
	return &PostgresRepo{DB: db}
}

// placeholders menghasilkan $start, $start+1, ... sebanyak n.
func placeholders(start, n int) []string {
	out := make([]string, n)
	for i := range out {
		out[i] = "$" + strconv.Itoa(start+i)
	}
	return out
}

var assetColumns = db.MakeColm(ID, FILENAME, FILE_KEY, MIME, AUTHOR_ID, CATEGORY, STATUS, PARENT_ID, SIZE, CREATED_AT, UPDATED_AT)

func (r *PostgresRepo) selectAssets(ctx context.Context, whereCol string, values []string) ([]AssetMetaData, error) {
	if len(values) == 0 {
		return []AssetMetaData{}, ErrAssetNotFound
	}
	query := `SELECT ` + strings.Join(assetColumns, ",") + ` FROM ` + string(TABLE) +
		` WHERE ` + whereCol + ` IN (` + strings.Join(placeholders(1, len(values)), ",") + `)`
	rows, err := r.DB.QueryContext(ctx, query, db.ToArgs(values)...)
	if err != nil {
		return []AssetMetaData{}, err
	}
	defer rows.Close()

	var assets []AssetMetaData
	for rows.Next() {
		var nullParentID sql.NullString
		var asset AssetMetaData
		if err := rows.Scan(&asset.ID, &asset.Filename, &asset.FileKey, &asset.Mime, &asset.AuthorID, &asset.Category, &asset.Status, &nullParentID, &asset.Size, &asset.CreatedAt, &asset.UpdatedAt); err != nil {
			return []AssetMetaData{}, err
		}
		if nullParentID.Valid {
			asset.ParentID = nullParentID.String
		}
		assets = append(assets, asset)
	}
	if err := rows.Err(); err != nil {
		return []AssetMetaData{}, err
	}
	if len(assets) == 0 {
		return []AssetMetaData{}, ErrAssetNotFound
	}
	return assets, nil
}

func (r *PostgresRepo) Record(ctx context.Context, payload RecordPayload) error {
	cols := db.MakeColm(ID, FILENAME, FILE_KEY, MIME, AUTHOR_ID, CATEGORY, STATUS)
	ph := placeholders(1, len(cols))
	input := []any{payload.ID, payload.Filename, payload.FileKey, payload.Mime, payload.AuthorID, payload.Category, payload.Status}

	query := `INSERT INTO ` + string(TABLE) + `(` + strings.Join(cols, ",") + `) VALUES (` + strings.Join(ph, ",") + `)`
	_, err := r.DB.ExecContext(ctx, query, input...)
	return err
}

func (r *PostgresRepo) MarkAsDeleted(ctx context.Context, payload DeletePayload) error {
	if len(payload.IDs) == 0 {
		return ErrAssetNotFound
	}
	// $1 = status, $2.. = ids
	inPH := placeholders(2, len(payload.IDs))
	query := `UPDATE ` + string(TABLE) + ` SET ` + string(STATUS) + ` = $1 WHERE ` + string(ID) + ` IN (` + strings.Join(inPH, ",") + `)`
	args := []any{AssetStatusDeleted}
	args = append(args, db.ToArgs(payload.IDs)...)
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected < 1 {
		return err
	}
	return nil
}

func (r *PostgresRepo) Update(ctx context.Context, payload UpdatePayload) error {
	query := `UPDATE ` + string(TABLE) +
		` SET ` + string(SIZE) + ` = $1, ` + string(STATUS) + ` = $2` +
		` WHERE ` + string(ID) + ` = $3 AND ` + string(STATUS) + ` = $4 AND ` + string(AUTHOR_ID) + ` = $5`
	input := []any{payload.Size, payload.Status, payload.Id, payload.OldStatus, payload.AuthorID}
	res, err := r.DB.ExecContext(ctx, query, db.ToArgs(input)...)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected != 1 {
		return err
	}
	return nil
}

func (r *PostgresRepo) GetByIDs(ctx context.Context, ids []string) ([]AssetMetaData, error) {
	return r.selectAssets(ctx, string(ID), ids)
}

func (r *PostgresRepo) GetByFileKeys(ctx context.Context, fileKeys []string) ([]AssetMetaData, error) {
	return r.selectAssets(ctx, string(FILE_KEY), fileKeys)
}

func (r *PostgresRepo) ReadByParentID(ctx context.Context, parentIDs []string) ([]AssetMetaData, error) {
	return r.selectAssets(ctx, string(PARENT_ID), parentIDs)
}

func (r *PostgresRepo) linkBy(ctx context.Context, whereCol string, values []string, parentID string) error {
	if parentID == "" || len(values) == 0 {
		return ErrAssetNotFound
	}
	// $1 = parent_id, $2.. = values
	inPH := placeholders(2, len(values))
	query := `UPDATE ` + string(TABLE) + ` SET ` + string(PARENT_ID) + ` = $1 WHERE ` + whereCol +
		` IN (` + strings.Join(inPH, ",") + `) AND ` + string(PARENT_ID) + ` IS NULL`
	args := []any{parentID}
	args = append(args, db.ToArgs(values)...)
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected < 1 {
		return ErrAssetNotFound
	}
	return nil
}

func (r *PostgresRepo) LinkingAsset(ctx context.Context, assetIDs []string, parentID string) error {
	return r.linkBy(ctx, string(ID), assetIDs, parentID)
}

func (r *PostgresRepo) LinkingAssetByFileKeys(ctx context.Context, fileKeys []string, parentID string) error {
	return r.linkBy(ctx, string(FILE_KEY), fileKeys, parentID)
}

func (r *PostgresRepo) releaseBy(ctx context.Context, whereCol string, values []string) error {
	if len(values) == 0 {
		return ErrAssetNotFound
	}
	inPH := placeholders(1, len(values))
	query := `UPDATE ` + string(TABLE) + ` SET ` + string(PARENT_ID) + ` = NULL WHERE ` + whereCol +
		` IN (` + strings.Join(inPH, ",") + `)`
	res, err := r.DB.ExecContext(ctx, query, db.ToArgs(values)...)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected < 1 {
		return ErrAssetNotFound
	}
	return nil
}

func (r *PostgresRepo) ReleaseByParentID(ctx context.Context, parentIDs []string) error {
	return r.releaseBy(ctx, string(PARENT_ID), parentIDs)
}

func (r *PostgresRepo) ReleaseByFileKeys(ctx context.Context, fileKeys []string) error {
	return r.releaseBy(ctx, string(FILE_KEY), fileKeys)
}
