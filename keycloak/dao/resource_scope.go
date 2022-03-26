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

// GetAllResourceScope is a function to get a slice of record(s) from resource_scope table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllResourceScope(ctx context.Context, page, pagesize int, order string) (results []*model.ResourceScope, totalRows int64, err error) {

	resultOrm := DB.Model(&model.ResourceScope{})
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

// GetResourceScope is a function to get a single record from the resource_scope table in the keycloak database
// error - ErrNotFound, db Find error
func GetResourceScope(ctx context.Context, argResourceID string, argScopeID string) (record *model.ResourceScope, err error) {
	record = &model.ResourceScope{}
	if err = DB.First(record, argResourceID, argScopeID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddResourceScope is a function to add a single record to resource_scope table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddResourceScope(ctx context.Context, record *model.ResourceScope) (result *model.ResourceScope, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateResourceScope is a function to update a single record from resource_scope table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateResourceScope(ctx context.Context, argResourceID string, argScopeID string, updated *model.ResourceScope) (result *model.ResourceScope, RowsAffected int64, err error) {

	result = &model.ResourceScope{}
	db := DB.First(result, "resource_id = ?", argResourceID, "scope_id = ?", argScopeID)
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

// DeleteResourceScope is a function to delete a single record from resource_scope table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteResourceScope(ctx context.Context, argResourceID string, argScopeID string) (rowsAffected int64, err error) {

	record := &model.ResourceScope{}
	db := DB.First(record, "resource_id = ?", argResourceID, "scope_id = ?", argScopeID)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
