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

// GetAllUserFederationConfig is a function to get a slice of record(s) from user_federation_config table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllUserFederationConfig(ctx context.Context, page, pagesize int, order string) (results []*model.UserFederationConfig, totalRows int64, err error) {

	resultOrm := DB.Model(&model.UserFederationConfig{})
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

// GetUserFederationConfig is a function to get a single record from the user_federation_config table in the keycloak database
// error - ErrNotFound, db Find error
func GetUserFederationConfig(ctx context.Context, argUserFederationProviderID string, argName string) (record *model.UserFederationConfig, err error) {
	record = &model.UserFederationConfig{}
	if err = DB.First(record, argUserFederationProviderID, argName).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddUserFederationConfig is a function to add a single record to user_federation_config table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddUserFederationConfig(ctx context.Context, record *model.UserFederationConfig) (result *model.UserFederationConfig, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateUserFederationConfig is a function to update a single record from user_federation_config table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateUserFederationConfig(ctx context.Context, argUserFederationProviderID string, argName string, updated *model.UserFederationConfig) (result *model.UserFederationConfig, RowsAffected int64, err error) {

	result = &model.UserFederationConfig{}
	db := DB.First(result, "user_federation_provider_id = ?", argUserFederationProviderID, "name = ?", argName)
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

// DeleteUserFederationConfig is a function to delete a single record from user_federation_config table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteUserFederationConfig(ctx context.Context, argUserFederationProviderID string, argName string) (rowsAffected int64, err error) {

	record := &model.UserFederationConfig{}
	db := DB.First(record, "user_federation_provider_id = ?", argUserFederationProviderID, "name = ?", argName)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
