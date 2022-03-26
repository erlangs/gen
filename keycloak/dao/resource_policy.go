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

// GetAllResourcePolicy is a function to get a slice of record(s) from resource_policy table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllResourcePolicy(ctx context.Context, page, pagesize int, order string) (results []*model.ResourcePolicy, totalRows int64, err error) {

	resultOrm := DB.Model(&model.ResourcePolicy{})
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

// GetResourcePolicy is a function to get a single record from the resource_policy table in the keycloak database
// error - ErrNotFound, db Find error
func GetResourcePolicy(ctx context.Context, argResourceID string, argPolicyID string) (record *model.ResourcePolicy, err error) {
	record = &model.ResourcePolicy{}
	if err = DB.First(record, argResourceID, argPolicyID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddResourcePolicy is a function to add a single record to resource_policy table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddResourcePolicy(ctx context.Context, record *model.ResourcePolicy) (result *model.ResourcePolicy, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateResourcePolicy is a function to update a single record from resource_policy table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateResourcePolicy(ctx context.Context, argResourceID string, argPolicyID string, updated *model.ResourcePolicy) (result *model.ResourcePolicy, RowsAffected int64, err error) {

	result = &model.ResourcePolicy{}
	db := DB.First(result, "resource_id = ?", argResourceID, "policy_id = ?", argPolicyID)
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

// DeleteResourcePolicy is a function to delete a single record from resource_policy table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteResourcePolicy(ctx context.Context, argResourceID string, argPolicyID string) (rowsAffected int64, err error) {

	record := &model.ResourcePolicy{}
	db := DB.First(record, "resource_id = ?", argResourceID, "policy_id = ?", argPolicyID)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
