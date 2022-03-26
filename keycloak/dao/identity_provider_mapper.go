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

// GetAllIdentityProviderMapper is a function to get a slice of record(s) from identity_provider_mapper table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllIdentityProviderMapper(ctx context.Context, page, pagesize int, order string) (results []*model.IdentityProviderMapper, totalRows int64, err error) {

	resultOrm := DB.Model(&model.IdentityProviderMapper{})
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

// GetIdentityProviderMapper is a function to get a single record from the identity_provider_mapper table in the keycloak database
// error - ErrNotFound, db Find error
func GetIdentityProviderMapper(ctx context.Context, argID string) (record *model.IdentityProviderMapper, err error) {
	record = &model.IdentityProviderMapper{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddIdentityProviderMapper is a function to add a single record to identity_provider_mapper table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddIdentityProviderMapper(ctx context.Context, record *model.IdentityProviderMapper) (result *model.IdentityProviderMapper, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateIdentityProviderMapper is a function to update a single record from identity_provider_mapper table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateIdentityProviderMapper(ctx context.Context, argID string, updated *model.IdentityProviderMapper) (result *model.IdentityProviderMapper, RowsAffected int64, err error) {

	result = &model.IdentityProviderMapper{}
	db := DB.First(result, "id = ?", argID)
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

// DeleteIdentityProviderMapper is a function to delete a single record from identity_provider_mapper table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteIdentityProviderMapper(ctx context.Context, argID string) (rowsAffected int64, err error) {

	record := &model.IdentityProviderMapper{}
	db := DB.First(record, "id = ?", argID)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
