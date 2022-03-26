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

// GetAllRealmEnabledEventTypes is a function to get a slice of record(s) from realm_enabled_event_types table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllRealmEnabledEventTypes(ctx context.Context, page, pagesize int, order string) (results []*model.RealmEnabledEventTypes, totalRows int64, err error) {

	resultOrm := DB.Model(&model.RealmEnabledEventTypes{})
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

// GetRealmEnabledEventTypes is a function to get a single record from the realm_enabled_event_types table in the keycloak database
// error - ErrNotFound, db Find error
func GetRealmEnabledEventTypes(ctx context.Context, argRealmID string, argValue string) (record *model.RealmEnabledEventTypes, err error) {
	record = &model.RealmEnabledEventTypes{}
	if err = DB.First(record, argRealmID, argValue).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddRealmEnabledEventTypes is a function to add a single record to realm_enabled_event_types table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddRealmEnabledEventTypes(ctx context.Context, record *model.RealmEnabledEventTypes) (result *model.RealmEnabledEventTypes, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateRealmEnabledEventTypes is a function to update a single record from realm_enabled_event_types table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateRealmEnabledEventTypes(ctx context.Context, argRealmID string, argValue string, updated *model.RealmEnabledEventTypes) (result *model.RealmEnabledEventTypes, RowsAffected int64, err error) {

	result = &model.RealmEnabledEventTypes{}
	db := DB.First(result, "realm_id = ?", argRealmID, "value = ?", argValue)
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

// DeleteRealmEnabledEventTypes is a function to delete a single record from realm_enabled_event_types table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteRealmEnabledEventTypes(ctx context.Context, argRealmID string, argValue string) (rowsAffected int64, err error) {

	record := &model.RealmEnabledEventTypes{}
	db := DB.First(record, "realm_id = ?", argRealmID, "value = ?", argValue)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
