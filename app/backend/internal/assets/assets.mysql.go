package assets

import (
	"context"
	"database/sql"
	"profile-template/internal/db"
	"strings"
)

type MysqlRepo struct {
	DB db.DBTX
}

func NewMysqlRepo(db db.DBTX) AssetRepository {
	return &MysqlRepo{DB: db}
}

func (r *MysqlRepo) Record(ctx context.Context, payload RecordPayload) error {
	collom := db.MakeColm(ID, FILENAME, FILE_KEY, MIME, AUTHOR_ID, CATEGORY, STATUS)
	placeHolder := db.MakePlaceHolder(len(collom))
	input := []any{payload.ID, payload.Filename, payload.FileKey, payload.Mime, payload.AuthorID, payload.Category, payload.Status}

	query := `INSERT INTO ` + string(TABLE) + `(` + strings.Join(collom, ",") + `) VALUES (` + strings.Join(placeHolder, ",") + `)`
	_, err := r.DB.ExecContext(ctx, query, input...)
	if err != nil {
		return err
	}
	return nil
}

func (r *MysqlRepo) MarkAsDeleted(ctx context.Context, payload DeletePayload) error {
	if len(payload.IDs) == 0 {
		return ErrAssetNotFound
	}
	inputPH := db.MakePlaceHolder(len(payload.IDs))
	query := `UPDATE ` + string(TABLE) + ` SET ` + string(STATUS) + ` = ? WHERE ` + string(ID) + ` IN (` + strings.Join(inputPH, ",") + `)`
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
	if rowsAffected < int64(len(payload.IDs)) {
		return nil
	}
	return nil

}

func (r *MysqlRepo) Update(ctx context.Context, payload UpdatePayload) error {
	collom := []string{string(SIZE), string(STATUS)}
	where := db.MakeColm(ID, STATUS, AUTHOR_ID)
	input := []any{payload.Size, payload.Status, payload.Id, payload.OldStatus, payload.AuthorID}
	query := `UPDATE ` + string(TABLE) + ` SET ` + strings.Join(collom, "= ?, ") + "= ? WHERE " + strings.Join(where, "= ? AND ") + " = ?"
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

func (r *MysqlRepo) GetByIDs(ctx context.Context, ids []string) ([]AssetMetaData, error) {
	if len(ids) == 0 {
		return []AssetMetaData{}, ErrAssetNotFound
	}
	collom := db.MakeColm(ID, FILENAME, FILE_KEY, MIME, AUTHOR_ID, CATEGORY, STATUS, PARENT_ID, SIZE, CREATED_AT, UPDATED_AT)
	ph := db.MakePlaceHolder(len(ids))
	query := `SELECT ` + strings.Join(collom, ",") + ` FROM ` + string(TABLE) + ` WHERE ` + string(ID) + ` IN (` + strings.Join(ph, ",") + `)`
	rows, err := r.DB.QueryContext(ctx, query, db.ToArgs(ids)...)
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
func (r *MysqlRepo) GetByFileKeys(ctx context.Context, fileKeys []string) ([]AssetMetaData, error) {
	if len(fileKeys) == 0 {
		return []AssetMetaData{}, ErrAssetNotFound
	}
	collom := db.MakeColm(ID, FILENAME, FILE_KEY, MIME, AUTHOR_ID, CATEGORY, STATUS, PARENT_ID, SIZE, CREATED_AT, UPDATED_AT)
	ph := db.MakePlaceHolder(len(fileKeys))
	query := `SELECT ` + strings.Join(collom, ",") + ` FROM ` + string(TABLE) + ` WHERE ` + string(FILE_KEY) + ` IN (` + strings.Join(ph, ",") + `)`
	rows, err := r.DB.QueryContext(ctx, query, db.ToArgs(fileKeys)...)
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

func (r *MysqlRepo) ReadByParentID(ctx context.Context, parentIDs []string) ([]AssetMetaData, error) {
	if len(parentIDs) == 0 {
		return []AssetMetaData{}, ErrAssetNotFound
	}
	collom := db.MakeColm(ID, FILENAME, FILE_KEY, MIME, AUTHOR_ID, CATEGORY, STATUS, PARENT_ID, SIZE, CREATED_AT, UPDATED_AT)
	ph := db.MakePlaceHolder(len(parentIDs))
	query := `SELECT ` + strings.Join(collom, ",") + ` FROM ` + string(TABLE) + ` WHERE ` + string(PARENT_ID) + ` IN (` + strings.Join(ph, ",") + `)`
	rows, err := r.DB.QueryContext(ctx, query, db.ToArgs(parentIDs)...)
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

func (r *MysqlRepo) LinkingAsset(ctx context.Context, assetIDs []string, parentID string) error {
	if parentID == "" || len(assetIDs) == 0 {
		return ErrAssetNotFound
	}
	inPH := db.MakePlaceHolder(len(assetIDs))
	query := `UPDATE ` + string(TABLE) + ` SET ` + string(PARENT_ID) + ` = ? WHERE ` + string(ID) + ` IN (` + strings.Join(inPH, ",") + `) AND ` + string(PARENT_ID) + ` IS NULL`
	args := []any{parentID}
	args = append(args, db.ToArgs(assetIDs)...)
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected < 1 {
		return ErrAssetNotFound
	}
	if rowsAffected < int64(len(assetIDs)) {
		return nil
	}
	return nil
}
func (r *MysqlRepo) LinkingAssetByFileKeys(ctx context.Context, fileKeys []string, parentID string) error {
	if parentID == "" || len(fileKeys) == 0 {
		return ErrAssetNotFound
	}
	inPH := db.MakePlaceHolder(len(fileKeys))
	query := `UPDATE ` + string(TABLE) + ` SET ` + string(PARENT_ID) + ` = ? WHERE ` + string(FILE_KEY) + ` IN (` + strings.Join(inPH, ",") + `) AND ` + string(PARENT_ID) + ` IS NULL`
	args := []any{parentID}
	args = append(args, db.ToArgs(fileKeys)...)
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected < 1 {
		return ErrAssetNotFound
	}
	if rowsAffected < int64(len(fileKeys)) {
		return nil
	}
	return nil
}

func (r *MysqlRepo) ReleaseByParentID(ctx context.Context, parentIDs []string) error {
	if len(parentIDs) == 0 {
		return ErrAssetNotFound
	}
	inPH := db.MakePlaceHolder(len(parentIDs))
	query := `UPDATE ` + string(TABLE) + ` SET ` + string(PARENT_ID) + ` = NULL WHERE ` + string(PARENT_ID) + ` IN (` + strings.Join(inPH, ",") + `)`
	res, err := r.DB.ExecContext(ctx, query, db.ToArgs(parentIDs)...)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected < 1 {
		return ErrAssetNotFound
	}
	return nil
}
func (r *MysqlRepo) ReleaseByFileKeys(ctx context.Context, fileKeys []string) error {
	if len(fileKeys) == 0 {
		return ErrAssetNotFound
	}
	inPH := db.MakePlaceHolder(len(fileKeys))
	query := `UPDATE ` + string(TABLE) + ` SET ` + string(PARENT_ID) + ` = NULL WHERE ` + string(FILE_KEY) + ` IN (` + strings.Join(inPH, ",") + `)`
	res, err := r.DB.ExecContext(ctx, query, db.ToArgs(fileKeys)...)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected < 1 {
		return ErrAssetNotFound
	}
	return nil
}
