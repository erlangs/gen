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

// GetAllRealmDefaultGroups is a function to get a slice of record(s) from realm_default_groups table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllRealmDefaultGroups(ctx context.Context, page, pagesize int, order string) (results []*model.RealmDefaultGroups, totalRows int64, err error) {

	resultOrm := DB.Model(&model.RealmDefaultGroups{})
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

// GetRealmDefaultGroups is a function to get a single record from the realm_default_groups table in the keycloak database
// error - ErrNotFound, db Find error
func GetRealmDefaultGroups(ctx context.Context, argRealmID string, argGroupID string) (record *model.RealmDefaultGroups, err error) {
	record = &model.RealmDefaultGroups{}
	if err = DB.First(record, argRealmID, argGroupID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddRealmDefaultGroups is a function to add a single record to realm_default_groups table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddRealmDefaultGroups(ctx context.Context, record *model.RealmDefaultGroups) (result *model.RealmDefaultGroups, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateRealmDefaultGroups is a function to update a single record from realm_default_groups table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateRealmDefaultGroups(ctx context.Context, argRealmID string, argGroupID string, updated *model.RealmDefaultGroups) (result *model.RealmDefaultGroups, RowsAffected int64, err error) {

	result = &model.RealmDefaultGroups{}
	db := DB.First(result, "realm_id = ?", argRealmID, "group_id = ?", argGroupID)
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

// DeleteRealmDefaultGroups is a function to delete a single record from realm_default_groups table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteRealmDefaultGroups(ctx context.Context, argRealmID string, argGroupID string) (rowsAffected int64, err error) {

	record := &model.RealmDefaultGroups{}
	db := DB.First(record, "realm_id = ?", argRealmID, "group_id = ?", argGroupID)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
