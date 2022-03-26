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

// GetAllScopePolicy is a function to get a slice of record(s) from scope_policy table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllScopePolicy(ctx context.Context, page, pagesize int, order string) (results []*model.ScopePolicy, totalRows int64, err error) {

	resultOrm := DB.Model(&model.ScopePolicy{})
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

// GetScopePolicy is a function to get a single record from the scope_policy table in the keycloak database
// error - ErrNotFound, db Find error
func GetScopePolicy(ctx context.Context, argScopeID string, argPolicyID string) (record *model.ScopePolicy, err error) {
	record = &model.ScopePolicy{}
	if err = DB.First(record, argScopeID, argPolicyID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddScopePolicy is a function to add a single record to scope_policy table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddScopePolicy(ctx context.Context, record *model.ScopePolicy) (result *model.ScopePolicy, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateScopePolicy is a function to update a single record from scope_policy table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateScopePolicy(ctx context.Context, argScopeID string, argPolicyID string, updated *model.ScopePolicy) (result *model.ScopePolicy, RowsAffected int64, err error) {

	result = &model.ScopePolicy{}
	db := DB.First(result, "scope_id = ?", argScopeID, "policy_id = ?", argPolicyID)
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

// DeleteScopePolicy is a function to delete a single record from scope_policy table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteScopePolicy(ctx context.Context, argScopeID string, argPolicyID string) (rowsAffected int64, err error) {

	record := &model.ScopePolicy{}
	db := DB.First(record, "scope_id = ?", argScopeID, "policy_id = ?", argPolicyID)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
