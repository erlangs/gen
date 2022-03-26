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

// GetAllClientAuthFlowBindings is a function to get a slice of record(s) from client_auth_flow_bindings table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllClientAuthFlowBindings(ctx context.Context, page, pagesize int, order string) (results []*model.ClientAuthFlowBindings, totalRows int64, err error) {

	resultOrm := DB.Model(&model.ClientAuthFlowBindings{})
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

// GetClientAuthFlowBindings is a function to get a single record from the client_auth_flow_bindings table in the keycloak database
// error - ErrNotFound, db Find error
func GetClientAuthFlowBindings(ctx context.Context, argClientID string, argBindingName string) (record *model.ClientAuthFlowBindings, err error) {
	record = &model.ClientAuthFlowBindings{}
	if err = DB.First(record, argClientID, argBindingName).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddClientAuthFlowBindings is a function to add a single record to client_auth_flow_bindings table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddClientAuthFlowBindings(ctx context.Context, record *model.ClientAuthFlowBindings) (result *model.ClientAuthFlowBindings, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateClientAuthFlowBindings is a function to update a single record from client_auth_flow_bindings table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateClientAuthFlowBindings(ctx context.Context, argClientID string, argBindingName string, updated *model.ClientAuthFlowBindings) (result *model.ClientAuthFlowBindings, RowsAffected int64, err error) {

	result = &model.ClientAuthFlowBindings{}
	db := DB.First(result, "client_id = ?", argClientID, "binding_name = ?", argBindingName)
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

// DeleteClientAuthFlowBindings is a function to delete a single record from client_auth_flow_bindings table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteClientAuthFlowBindings(ctx context.Context, argClientID string, argBindingName string) (rowsAffected int64, err error) {

	record := &model.ClientAuthFlowBindings{}
	db := DB.First(record, "client_id = ?", argClientID, "binding_name = ?", argBindingName)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
