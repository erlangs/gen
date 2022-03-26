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

// GetAllClientNodeRegistrations is a function to get a slice of record(s) from client_node_registrations table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllClientNodeRegistrations(ctx context.Context, page, pagesize int, order string) (results []*model.ClientNodeRegistrations, totalRows int64, err error) {

	resultOrm := DB.Model(&model.ClientNodeRegistrations{})
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

// GetClientNodeRegistrations is a function to get a single record from the client_node_registrations table in the keycloak database
// error - ErrNotFound, db Find error
func GetClientNodeRegistrations(ctx context.Context, argClientID string, argName string) (record *model.ClientNodeRegistrations, err error) {
	record = &model.ClientNodeRegistrations{}
	if err = DB.First(record, argClientID, argName).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddClientNodeRegistrations is a function to add a single record to client_node_registrations table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddClientNodeRegistrations(ctx context.Context, record *model.ClientNodeRegistrations) (result *model.ClientNodeRegistrations, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateClientNodeRegistrations is a function to update a single record from client_node_registrations table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateClientNodeRegistrations(ctx context.Context, argClientID string, argName string, updated *model.ClientNodeRegistrations) (result *model.ClientNodeRegistrations, RowsAffected int64, err error) {

	result = &model.ClientNodeRegistrations{}
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

// DeleteClientNodeRegistrations is a function to delete a single record from client_node_registrations table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteClientNodeRegistrations(ctx context.Context, argClientID string, argName string) (rowsAffected int64, err error) {

	record := &model.ClientNodeRegistrations{}
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
