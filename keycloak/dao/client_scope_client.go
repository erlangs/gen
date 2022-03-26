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

// GetAllClientScopeClient is a function to get a slice of record(s) from client_scope_client table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllClientScopeClient(ctx context.Context, page, pagesize int, order string) (results []*model.ClientScopeClient, totalRows int64, err error) {

	resultOrm := DB.Model(&model.ClientScopeClient{})
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

// GetClientScopeClient is a function to get a single record from the client_scope_client table in the keycloak database
// error - ErrNotFound, db Find error
func GetClientScopeClient(ctx context.Context, argClientID string, argScopeID string) (record *model.ClientScopeClient, err error) {
	record = &model.ClientScopeClient{}
	if err = DB.First(record, argClientID, argScopeID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddClientScopeClient is a function to add a single record to client_scope_client table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddClientScopeClient(ctx context.Context, record *model.ClientScopeClient) (result *model.ClientScopeClient, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateClientScopeClient is a function to update a single record from client_scope_client table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateClientScopeClient(ctx context.Context, argClientID string, argScopeID string, updated *model.ClientScopeClient) (result *model.ClientScopeClient, RowsAffected int64, err error) {

	result = &model.ClientScopeClient{}
	db := DB.First(result, "client_id = ?", argClientID, "scope_id = ?", argScopeID)
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

// DeleteClientScopeClient is a function to delete a single record from client_scope_client table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteClientScopeClient(ctx context.Context, argClientID string, argScopeID string) (rowsAffected int64, err error) {

	record := &model.ClientScopeClient{}
	db := DB.First(record, "client_id = ?", argClientID, "scope_id = ?", argScopeID)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
