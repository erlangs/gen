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

// GetAllAssociatedPolicy is a function to get a slice of record(s) from associated_policy table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllAssociatedPolicy(ctx context.Context, page, pagesize int, order string) (results []*model.AssociatedPolicy, totalRows int64, err error) {

	resultOrm := DB.Model(&model.AssociatedPolicy{})
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

// GetAssociatedPolicy is a function to get a single record from the associated_policy table in the keycloak database
// error - ErrNotFound, db Find error
func GetAssociatedPolicy(ctx context.Context, argPolicyID string, argAssociatedPolicyID string) (record *model.AssociatedPolicy, err error) {
	record = &model.AssociatedPolicy{}
	if err = DB.First(record, argPolicyID, argAssociatedPolicyID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddAssociatedPolicy is a function to add a single record to associated_policy table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddAssociatedPolicy(ctx context.Context, record *model.AssociatedPolicy) (result *model.AssociatedPolicy, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateAssociatedPolicy is a function to update a single record from associated_policy table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateAssociatedPolicy(ctx context.Context, argPolicyID string, argAssociatedPolicyID string, updated *model.AssociatedPolicy) (result *model.AssociatedPolicy, RowsAffected int64, err error) {

	result = &model.AssociatedPolicy{}
	db := DB.First(result, "policy_id = ?", argPolicyID, "associated_policy_id = ?", argAssociatedPolicyID)
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

// DeleteAssociatedPolicy is a function to delete a single record from associated_policy table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteAssociatedPolicy(ctx context.Context, argPolicyID string, argAssociatedPolicyID string) (rowsAffected int64, err error) {

	record := &model.AssociatedPolicy{}
	db := DB.First(record, "policy_id = ?", argPolicyID, "associated_policy_id = ?", argAssociatedPolicyID)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
