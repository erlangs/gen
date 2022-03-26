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

// GetAllProtocolMapperConfig is a function to get a slice of record(s) from protocol_mapper_config table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllProtocolMapperConfig(ctx context.Context, page, pagesize int, order string) (results []*model.ProtocolMapperConfig, totalRows int64, err error) {

	resultOrm := DB.Model(&model.ProtocolMapperConfig{})
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

// GetProtocolMapperConfig is a function to get a single record from the protocol_mapper_config table in the keycloak database
// error - ErrNotFound, db Find error
func GetProtocolMapperConfig(ctx context.Context, argProtocolMapperID string, argName string) (record *model.ProtocolMapperConfig, err error) {
	record = &model.ProtocolMapperConfig{}
	if err = DB.First(record, argProtocolMapperID, argName).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddProtocolMapperConfig is a function to add a single record to protocol_mapper_config table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddProtocolMapperConfig(ctx context.Context, record *model.ProtocolMapperConfig) (result *model.ProtocolMapperConfig, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateProtocolMapperConfig is a function to update a single record from protocol_mapper_config table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateProtocolMapperConfig(ctx context.Context, argProtocolMapperID string, argName string, updated *model.ProtocolMapperConfig) (result *model.ProtocolMapperConfig, RowsAffected int64, err error) {

	result = &model.ProtocolMapperConfig{}
	db := DB.First(result, "protocol_mapper_id = ?", argProtocolMapperID, "name = ?", argName)
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

// DeleteProtocolMapperConfig is a function to delete a single record from protocol_mapper_config table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteProtocolMapperConfig(ctx context.Context, argProtocolMapperID string, argName string) (rowsAffected int64, err error) {

	record := &model.ProtocolMapperConfig{}
	db := DB.First(record, "protocol_mapper_id = ?", argProtocolMapperID, "name = ?", argName)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
