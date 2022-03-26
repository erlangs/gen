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

// GetAllPolicyConfig is a function to get a slice of record(s) from policy_config table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllPolicyConfig(ctx context.Context, page, pagesize int, order string) (results []*model.PolicyConfig, totalRows int64, err error) {

	resultOrm := DB.Model(&model.PolicyConfig{})
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

// GetPolicyConfig is a function to get a single record from the policy_config table in the keycloak database
// error - ErrNotFound, db Find error
func GetPolicyConfig(ctx context.Context, argPolicyID string, argName string) (record *model.PolicyConfig, err error) {
	record = &model.PolicyConfig{}
	if err = DB.First(record, argPolicyID, argName).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddPolicyConfig is a function to add a single record to policy_config table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddPolicyConfig(ctx context.Context, record *model.PolicyConfig) (result *model.PolicyConfig, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdatePolicyConfig is a function to update a single record from policy_config table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdatePolicyConfig(ctx context.Context, argPolicyID string, argName string, updated *model.PolicyConfig) (result *model.PolicyConfig, RowsAffected int64, err error) {

	result = &model.PolicyConfig{}
	db := DB.First(result, "policy_id = ?", argPolicyID, "name = ?", argName)
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

// DeletePolicyConfig is a function to delete a single record from policy_config table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeletePolicyConfig(ctx context.Context, argPolicyID string, argName string) (rowsAffected int64, err error) {

	record := &model.PolicyConfig{}
	db := DB.First(record, "policy_id = ?", argPolicyID, "name = ?", argName)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
