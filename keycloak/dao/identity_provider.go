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

// GetAllIdentityProvider is a function to get a slice of record(s) from identity_provider table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllIdentityProvider(ctx context.Context, page, pagesize int, order string) (results []*model.IdentityProvider, totalRows int64, err error) {

	resultOrm := DB.Model(&model.IdentityProvider{})
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

// GetIdentityProvider is a function to get a single record from the identity_provider table in the keycloak database
// error - ErrNotFound, db Find error
func GetIdentityProvider(ctx context.Context, argInternalID string) (record *model.IdentityProvider, err error) {
	record = &model.IdentityProvider{}
	if err = DB.First(record, argInternalID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddIdentityProvider is a function to add a single record to identity_provider table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddIdentityProvider(ctx context.Context, record *model.IdentityProvider) (result *model.IdentityProvider, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateIdentityProvider is a function to update a single record from identity_provider table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateIdentityProvider(ctx context.Context, argInternalID string, updated *model.IdentityProvider) (result *model.IdentityProvider, RowsAffected int64, err error) {

	result = &model.IdentityProvider{}
	db := DB.First(result, "internal_id = ?", argInternalID)
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

// DeleteIdentityProvider is a function to delete a single record from identity_provider table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteIdentityProvider(ctx context.Context, argInternalID string) (rowsAffected int64, err error) {

	record := &model.IdentityProvider{}
	db := DB.First(record, "internal_id = ?", argInternalID)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
