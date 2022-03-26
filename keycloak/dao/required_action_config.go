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

// GetAllRequiredActionConfig is a function to get a slice of record(s) from required_action_config table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllRequiredActionConfig(ctx context.Context, page, pagesize int, order string) (results []*model.RequiredActionConfig, totalRows int64, err error) {

	resultOrm := DB.Model(&model.RequiredActionConfig{})
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

// GetRequiredActionConfig is a function to get a single record from the required_action_config table in the keycloak database
// error - ErrNotFound, db Find error
func GetRequiredActionConfig(ctx context.Context, argRequiredActionID string, argName string) (record *model.RequiredActionConfig, err error) {
	record = &model.RequiredActionConfig{}
	if err = DB.First(record, argRequiredActionID, argName).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddRequiredActionConfig is a function to add a single record to required_action_config table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddRequiredActionConfig(ctx context.Context, record *model.RequiredActionConfig) (result *model.RequiredActionConfig, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateRequiredActionConfig is a function to update a single record from required_action_config table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateRequiredActionConfig(ctx context.Context, argRequiredActionID string, argName string, updated *model.RequiredActionConfig) (result *model.RequiredActionConfig, RowsAffected int64, err error) {

	result = &model.RequiredActionConfig{}
	db := DB.First(result, "required_action_id = ?", argRequiredActionID, "name = ?", argName)
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

// DeleteRequiredActionConfig is a function to delete a single record from required_action_config table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteRequiredActionConfig(ctx context.Context, argRequiredActionID string, argName string) (rowsAffected int64, err error) {

	record := &model.RequiredActionConfig{}
	db := DB.First(record, "required_action_id = ?", argRequiredActionID, "name = ?", argName)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
