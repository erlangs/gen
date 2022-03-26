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

// GetAllIdpMapperConfig is a function to get a slice of record(s) from idp_mapper_config table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllIdpMapperConfig(ctx context.Context, page, pagesize int, order string) (results []*model.IdpMapperConfig, totalRows int64, err error) {

	resultOrm := DB.Model(&model.IdpMapperConfig{})
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

// GetIdpMapperConfig is a function to get a single record from the idp_mapper_config table in the keycloak database
// error - ErrNotFound, db Find error
func GetIdpMapperConfig(ctx context.Context, argIdpMapperID string, argName string) (record *model.IdpMapperConfig, err error) {
	record = &model.IdpMapperConfig{}
	if err = DB.First(record, argIdpMapperID, argName).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddIdpMapperConfig is a function to add a single record to idp_mapper_config table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddIdpMapperConfig(ctx context.Context, record *model.IdpMapperConfig) (result *model.IdpMapperConfig, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateIdpMapperConfig is a function to update a single record from idp_mapper_config table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateIdpMapperConfig(ctx context.Context, argIdpMapperID string, argName string, updated *model.IdpMapperConfig) (result *model.IdpMapperConfig, RowsAffected int64, err error) {

	result = &model.IdpMapperConfig{}
	db := DB.First(result, "idp_mapper_id = ?", argIdpMapperID, "name = ?", argName)
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

// DeleteIdpMapperConfig is a function to delete a single record from idp_mapper_config table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteIdpMapperConfig(ctx context.Context, argIdpMapperID string, argName string) (rowsAffected int64, err error) {

	record := &model.IdpMapperConfig{}
	db := DB.First(record, "idp_mapper_id = ?", argIdpMapperID, "name = ?", argName)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
