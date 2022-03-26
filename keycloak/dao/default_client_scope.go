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

// GetAllDefaultClientScope is a function to get a slice of record(s) from default_client_scope table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllDefaultClientScope(ctx context.Context, page, pagesize int, order string) (results []*model.DefaultClientScope, totalRows int64, err error) {

	resultOrm := DB.Model(&model.DefaultClientScope{})
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

// GetDefaultClientScope is a function to get a single record from the default_client_scope table in the keycloak database
// error - ErrNotFound, db Find error
func GetDefaultClientScope(ctx context.Context, argRealmID string, argScopeID string) (record *model.DefaultClientScope, err error) {
	record = &model.DefaultClientScope{}
	if err = DB.First(record, argRealmID, argScopeID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddDefaultClientScope is a function to add a single record to default_client_scope table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddDefaultClientScope(ctx context.Context, record *model.DefaultClientScope) (result *model.DefaultClientScope, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateDefaultClientScope is a function to update a single record from default_client_scope table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateDefaultClientScope(ctx context.Context, argRealmID string, argScopeID string, updated *model.DefaultClientScope) (result *model.DefaultClientScope, RowsAffected int64, err error) {

	result = &model.DefaultClientScope{}
	db := DB.First(result, "realm_id = ?", argRealmID, "scope_id = ?", argScopeID)
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

// DeleteDefaultClientScope is a function to delete a single record from default_client_scope table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteDefaultClientScope(ctx context.Context, argRealmID string, argScopeID string) (rowsAffected int64, err error) {

	record := &model.DefaultClientScope{}
	db := DB.First(record, "realm_id = ?", argRealmID, "scope_id = ?", argScopeID)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
