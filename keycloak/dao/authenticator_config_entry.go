package dao

import (
	"context"
	"time"

	"keycloak/rest/api/model"

	uuid "github.com/satori/go.uuid"
)

var (
	_ = time.Second

	_ = uuid.UUID{}
)

// GetAllAuthenticatorConfigEntry is a function to get a slice of record(s) from authenticator_config_entry table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllAuthenticatorConfigEntry(ctx context.Context, page, pagesize int, order string) (results []*model.AuthenticatorConfigEntry, totalRows int64, err error) {

	resultOrm := DB.Model(&model.AuthenticatorConfigEntry{})
	resultOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		resultOrm = resultOrm.Offset(offset).Limit(pagesize)
	} else {
		resultOrm = resultOrm.Limit(pagesize)
	}

	if order != "" {
		resultOrm = resultOrm.Order(order)
	}

	if err = resultOrm.Find(&results).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return results, totalRows, nil
}

// GetAuthenticatorConfigEntry is a function to get a single record from the authenticator_config_entry table in the keycloak database
// error - ErrNotFound, db Find error
func GetAuthenticatorConfigEntry(ctx context.Context, argAuthenticatorID string, argName string) (record *model.AuthenticatorConfigEntry, err error) {
	record = &model.AuthenticatorConfigEntry{}
	if err = DB.First(record, argAuthenticatorID, argName).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddAuthenticatorConfigEntry is a function to add a single record to authenticator_config_entry table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddAuthenticatorConfigEntry(ctx context.Context, record *model.AuthenticatorConfigEntry) (result *model.AuthenticatorConfigEntry, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateAuthenticatorConfigEntry is a function to update a single record from authenticator_config_entry table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateAuthenticatorConfigEntry(ctx context.Context, argAuthenticatorID string, argName string, updated *model.AuthenticatorConfigEntry) (result *model.AuthenticatorConfigEntry, RowsAffected int64, err error) {

	result = &model.AuthenticatorConfigEntry{}
	db := DB.First(result, "authenticator_id = ?", argAuthenticatorID, "name = ?", argName)
	if err = db.Error; err != nil {
		return nil, -1, ErrNotFound
	}

	if err = Copy(result, updated); err != nil {
		return nil, -1, ErrUpdateFailed
	}

	db = db.Save(result)
	if err = db.Error; err != nil {
		return nil, -1, ErrUpdateFailed
	}

	return result, db.RowsAffected, nil
}

// DeleteAuthenticatorConfigEntry is a function to delete a single record from authenticator_config_entry table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteAuthenticatorConfigEntry(ctx context.Context, argAuthenticatorID string, argName string) (rowsAffected int64, err error) {

	record := &model.AuthenticatorConfigEntry{}
	db := DB.First(record, "authenticator_id = ?", argAuthenticatorID, "name = ?", argName)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
