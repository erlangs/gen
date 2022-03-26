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

// GetAllClientAttributes is a function to get a slice of record(s) from client_attributes table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllClientAttributes(ctx context.Context, page, pagesize int, order string) (results []*model.ClientAttributes, totalRows int64, err error) {

	resultOrm := DB.Model(&model.ClientAttributes{})
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

// GetClientAttributes is a function to get a single record from the client_attributes table in the keycloak database
// error - ErrNotFound, db Find error
func GetClientAttributes(ctx context.Context, argClientID string, argName string) (record *model.ClientAttributes, err error) {
	record = &model.ClientAttributes{}
	if err = DB.First(record, argClientID, argName).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddClientAttributes is a function to add a single record to client_attributes table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddClientAttributes(ctx context.Context, record *model.ClientAttributes) (result *model.ClientAttributes, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateClientAttributes is a function to update a single record from client_attributes table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateClientAttributes(ctx context.Context, argClientID string, argName string, updated *model.ClientAttributes) (result *model.ClientAttributes, RowsAffected int64, err error) {

	result = &model.ClientAttributes{}
	db := DB.First(result, "client_id = ?", argClientID, "name = ?", argName)
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

// DeleteClientAttributes is a function to delete a single record from client_attributes table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteClientAttributes(ctx context.Context, argClientID string, argName string) (rowsAffected int64, err error) {

	record := &model.ClientAttributes{}
	db := DB.First(record, "client_id = ?", argClientID, "name = ?", argName)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
