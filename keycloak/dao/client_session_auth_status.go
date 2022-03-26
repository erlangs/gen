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

// GetAllClientSessionAuthStatus is a function to get a slice of record(s) from client_session_auth_status table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllClientSessionAuthStatus(ctx context.Context, page, pagesize int, order string) (results []*model.ClientSessionAuthStatus, totalRows int64, err error) {

	resultOrm := DB.Model(&model.ClientSessionAuthStatus{})
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

// GetClientSessionAuthStatus is a function to get a single record from the client_session_auth_status table in the keycloak database
// error - ErrNotFound, db Find error
func GetClientSessionAuthStatus(ctx context.Context, argAuthenticator string, argClientSession string) (record *model.ClientSessionAuthStatus, err error) {
	record = &model.ClientSessionAuthStatus{}
	if err = DB.First(record, argAuthenticator, argClientSession).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddClientSessionAuthStatus is a function to add a single record to client_session_auth_status table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddClientSessionAuthStatus(ctx context.Context, record *model.ClientSessionAuthStatus) (result *model.ClientSessionAuthStatus, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateClientSessionAuthStatus is a function to update a single record from client_session_auth_status table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateClientSessionAuthStatus(ctx context.Context, argAuthenticator string, argClientSession string, updated *model.ClientSessionAuthStatus) (result *model.ClientSessionAuthStatus, RowsAffected int64, err error) {

	result = &model.ClientSessionAuthStatus{}
	db := DB.First(result, "authenticator = ?", argAuthenticator, "client_session = ?", argClientSession)
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

// DeleteClientSessionAuthStatus is a function to delete a single record from client_session_auth_status table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteClientSessionAuthStatus(ctx context.Context, argAuthenticator string, argClientSession string) (rowsAffected int64, err error) {

	record := &model.ClientSessionAuthStatus{}
	db := DB.First(record, "authenticator = ?", argAuthenticator, "client_session = ?", argClientSession)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
