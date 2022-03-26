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

// GetAllRedirectUris is a function to get a slice of record(s) from redirect_uris table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllRedirectUris(ctx context.Context, page, pagesize int, order string) (results []*model.RedirectUris, totalRows int64, err error) {

	resultOrm := DB.Model(&model.RedirectUris{})
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

// GetRedirectUris is a function to get a single record from the redirect_uris table in the keycloak database
// error - ErrNotFound, db Find error
func GetRedirectUris(ctx context.Context, argClientID string, argValue string) (record *model.RedirectUris, err error) {
	record = &model.RedirectUris{}
	if err = DB.First(record, argClientID, argValue).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddRedirectUris is a function to add a single record to redirect_uris table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddRedirectUris(ctx context.Context, record *model.RedirectUris) (result *model.RedirectUris, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateRedirectUris is a function to update a single record from redirect_uris table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateRedirectUris(ctx context.Context, argClientID string, argValue string, updated *model.RedirectUris) (result *model.RedirectUris, RowsAffected int64, err error) {

	result = &model.RedirectUris{}
	db := DB.First(result, "client_id = ?", argClientID, "value = ?", argValue)
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

// DeleteRedirectUris is a function to delete a single record from redirect_uris table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteRedirectUris(ctx context.Context, argClientID string, argValue string) (rowsAffected int64, err error) {

	record := &model.RedirectUris{}
	db := DB.First(record, "client_id = ?", argClientID, "value = ?", argValue)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
