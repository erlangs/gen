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

// GetAllRealmRequiredCredential is a function to get a slice of record(s) from realm_required_credential table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllRealmRequiredCredential(ctx context.Context, page, pagesize int, order string) (results []*model.RealmRequiredCredential, totalRows int64, err error) {

	resultOrm := DB.Model(&model.RealmRequiredCredential{})
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

// GetRealmRequiredCredential is a function to get a single record from the realm_required_credential table in the keycloak database
// error - ErrNotFound, db Find error
func GetRealmRequiredCredential(ctx context.Context, argType string, argRealmID string) (record *model.RealmRequiredCredential, err error) {
	record = &model.RealmRequiredCredential{}
	if err = DB.First(record, argType, argRealmID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddRealmRequiredCredential is a function to add a single record to realm_required_credential table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddRealmRequiredCredential(ctx context.Context, record *model.RealmRequiredCredential) (result *model.RealmRequiredCredential, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateRealmRequiredCredential is a function to update a single record from realm_required_credential table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateRealmRequiredCredential(ctx context.Context, argType string, argRealmID string, updated *model.RealmRequiredCredential) (result *model.RealmRequiredCredential, RowsAffected int64, err error) {

	result = &model.RealmRequiredCredential{}
	db := DB.First(result, "type = ?", argType, "realm_id = ?", argRealmID)
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

// DeleteRealmRequiredCredential is a function to delete a single record from realm_required_credential table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteRealmRequiredCredential(ctx context.Context, argType string, argRealmID string) (rowsAffected int64, err error) {

	record := &model.RealmRequiredCredential{}
	db := DB.First(record, "type = ?", argType, "realm_id = ?", argRealmID)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
