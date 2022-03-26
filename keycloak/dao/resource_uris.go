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

// GetAllResourceUris is a function to get a slice of record(s) from resource_uris table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllResourceUris(ctx context.Context, page, pagesize int, order string) (results []*model.ResourceUris, totalRows int64, err error) {

	resultOrm := DB.Model(&model.ResourceUris{})
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

// GetResourceUris is a function to get a single record from the resource_uris table in the keycloak database
// error - ErrNotFound, db Find error
func GetResourceUris(ctx context.Context, argResourceID string, argValue string) (record *model.ResourceUris, err error) {
	record = &model.ResourceUris{}
	if err = DB.First(record, argResourceID, argValue).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddResourceUris is a function to add a single record to resource_uris table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddResourceUris(ctx context.Context, record *model.ResourceUris) (result *model.ResourceUris, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateResourceUris is a function to update a single record from resource_uris table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateResourceUris(ctx context.Context, argResourceID string, argValue string, updated *model.ResourceUris) (result *model.ResourceUris, RowsAffected int64, err error) {

	result = &model.ResourceUris{}
	db := DB.First(result, "resource_id = ?", argResourceID, "value = ?", argValue)
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

// DeleteResourceUris is a function to delete a single record from resource_uris table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteResourceUris(ctx context.Context, argResourceID string, argValue string) (rowsAffected int64, err error) {

	record := &model.ResourceUris{}
	db := DB.First(record, "resource_id = ?", argResourceID, "value = ?", argValue)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
