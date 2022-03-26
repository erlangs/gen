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

// GetAllRealmSMTPConfig is a function to get a slice of record(s) from realm_smtp_config table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllRealmSMTPConfig(ctx context.Context, page, pagesize int, order string) (results []*model.RealmSMTPConfig, totalRows int64, err error) {

	resultOrm := DB.Model(&model.RealmSMTPConfig{})
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

// GetRealmSMTPConfig is a function to get a single record from the realm_smtp_config table in the keycloak database
// error - ErrNotFound, db Find error
func GetRealmSMTPConfig(ctx context.Context, argRealmID string, argName string) (record *model.RealmSMTPConfig, err error) {
	record = &model.RealmSMTPConfig{}
	if err = DB.First(record, argRealmID, argName).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddRealmSMTPConfig is a function to add a single record to realm_smtp_config table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddRealmSMTPConfig(ctx context.Context, record *model.RealmSMTPConfig) (result *model.RealmSMTPConfig, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateRealmSMTPConfig is a function to update a single record from realm_smtp_config table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateRealmSMTPConfig(ctx context.Context, argRealmID string, argName string, updated *model.RealmSMTPConfig) (result *model.RealmSMTPConfig, RowsAffected int64, err error) {

	result = &model.RealmSMTPConfig{}
	db := DB.First(result, "realm_id = ?", argRealmID, "name = ?", argName)
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

// DeleteRealmSMTPConfig is a function to delete a single record from realm_smtp_config table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteRealmSMTPConfig(ctx context.Context, argRealmID string, argName string) (rowsAffected int64, err error) {

	record := &model.RealmSMTPConfig{}
	db := DB.First(record, "realm_id = ?", argRealmID, "name = ?", argName)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
