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

// GetAllUserFederationMapperConfig is a function to get a slice of record(s) from user_federation_mapper_config table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllUserFederationMapperConfig(ctx context.Context, page, pagesize int, order string) (results []*model.UserFederationMapperConfig, totalRows int64, err error) {

	resultOrm := DB.Model(&model.UserFederationMapperConfig{})
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

// GetUserFederationMapperConfig is a function to get a single record from the user_federation_mapper_config table in the keycloak database
// error - ErrNotFound, db Find error
func GetUserFederationMapperConfig(ctx context.Context, argUserFederationMapperID string, argName string) (record *model.UserFederationMapperConfig, err error) {
	record = &model.UserFederationMapperConfig{}
	if err = DB.First(record, argUserFederationMapperID, argName).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddUserFederationMapperConfig is a function to add a single record to user_federation_mapper_config table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddUserFederationMapperConfig(ctx context.Context, record *model.UserFederationMapperConfig) (result *model.UserFederationMapperConfig, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateUserFederationMapperConfig is a function to update a single record from user_federation_mapper_config table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateUserFederationMapperConfig(ctx context.Context, argUserFederationMapperID string, argName string, updated *model.UserFederationMapperConfig) (result *model.UserFederationMapperConfig, RowsAffected int64, err error) {

	result = &model.UserFederationMapperConfig{}
	db := DB.First(result, "user_federation_mapper_id = ?", argUserFederationMapperID, "name = ?", argName)
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

// DeleteUserFederationMapperConfig is a function to delete a single record from user_federation_mapper_config table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteUserFederationMapperConfig(ctx context.Context, argUserFederationMapperID string, argName string) (rowsAffected int64, err error) {

	record := &model.UserFederationMapperConfig{}
	db := DB.First(record, "user_federation_mapper_id = ?", argUserFederationMapperID, "name = ?", argName)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
