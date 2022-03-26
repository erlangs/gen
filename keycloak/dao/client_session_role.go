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

// GetAllClientSessionRole is a function to get a slice of record(s) from client_session_role table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllClientSessionRole(ctx context.Context, page, pagesize int, order string) (results []*model.ClientSessionRole, totalRows int64, err error) {

	resultOrm := DB.Model(&model.ClientSessionRole{})
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

// GetClientSessionRole is a function to get a single record from the client_session_role table in the keycloak database
// error - ErrNotFound, db Find error
func GetClientSessionRole(ctx context.Context, argRoleID string, argClientSession string) (record *model.ClientSessionRole, err error) {
	record = &model.ClientSessionRole{}
	if err = DB.First(record, argRoleID, argClientSession).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddClientSessionRole is a function to add a single record to client_session_role table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddClientSessionRole(ctx context.Context, record *model.ClientSessionRole) (result *model.ClientSessionRole, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateClientSessionRole is a function to update a single record from client_session_role table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateClientSessionRole(ctx context.Context, argRoleID string, argClientSession string, updated *model.ClientSessionRole) (result *model.ClientSessionRole, RowsAffected int64, err error) {

	result = &model.ClientSessionRole{}
	db := DB.First(result, "role_id = ?", argRoleID, "client_session = ?", argClientSession)
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

// DeleteClientSessionRole is a function to delete a single record from client_session_role table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteClientSessionRole(ctx context.Context, argRoleID string, argClientSession string) (rowsAffected int64, err error) {

	record := &model.ClientSessionRole{}
	db := DB.First(record, "role_id = ?", argRoleID, "client_session = ?", argClientSession)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
