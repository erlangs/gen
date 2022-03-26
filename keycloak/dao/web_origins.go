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

// GetAllWebOrigins is a function to get a slice of record(s) from web_origins table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllWebOrigins(ctx context.Context, page, pagesize int, order string) (results []*model.WebOrigins, totalRows int64, err error) {

	resultOrm := DB.Model(&model.WebOrigins{})
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

// GetWebOrigins is a function to get a single record from the web_origins table in the keycloak database
// error - ErrNotFound, db Find error
func GetWebOrigins(ctx context.Context, argClientID string, argValue string) (record *model.WebOrigins, err error) {
	record = &model.WebOrigins{}
	if err = DB.First(record, argClientID, argValue).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddWebOrigins is a function to add a single record to web_origins table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddWebOrigins(ctx context.Context, record *model.WebOrigins) (result *model.WebOrigins, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateWebOrigins is a function to update a single record from web_origins table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateWebOrigins(ctx context.Context, argClientID string, argValue string, updated *model.WebOrigins) (result *model.WebOrigins, RowsAffected int64, err error) {

	result = &model.WebOrigins{}
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

// DeleteWebOrigins is a function to delete a single record from web_origins table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteWebOrigins(ctx context.Context, argClientID string, argValue string) (rowsAffected int64, err error) {

	record := &model.WebOrigins{}
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
