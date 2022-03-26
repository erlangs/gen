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

// GetAllClientScopeRoleMapping is a function to get a slice of record(s) from client_scope_role_mapping table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllClientScopeRoleMapping(ctx context.Context, page, pagesize int, order string) (results []*model.ClientScopeRoleMapping, totalRows int64, err error) {

	resultOrm := DB.Model(&model.ClientScopeRoleMapping{})
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

// GetClientScopeRoleMapping is a function to get a single record from the client_scope_role_mapping table in the keycloak database
// error - ErrNotFound, db Find error
func GetClientScopeRoleMapping(ctx context.Context, argScopeID string, argRoleID string) (record *model.ClientScopeRoleMapping, err error) {
	record = &model.ClientScopeRoleMapping{}
	if err = DB.First(record, argScopeID, argRoleID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddClientScopeRoleMapping is a function to add a single record to client_scope_role_mapping table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddClientScopeRoleMapping(ctx context.Context, record *model.ClientScopeRoleMapping) (result *model.ClientScopeRoleMapping, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateClientScopeRoleMapping is a function to update a single record from client_scope_role_mapping table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateClientScopeRoleMapping(ctx context.Context, argScopeID string, argRoleID string, updated *model.ClientScopeRoleMapping) (result *model.ClientScopeRoleMapping, RowsAffected int64, err error) {

	result = &model.ClientScopeRoleMapping{}
	db := DB.First(result, "scope_id = ?", argScopeID, "role_id = ?", argRoleID)
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

// DeleteClientScopeRoleMapping is a function to delete a single record from client_scope_role_mapping table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteClientScopeRoleMapping(ctx context.Context, argScopeID string, argRoleID string) (rowsAffected int64, err error) {

	record := &model.ClientScopeRoleMapping{}
	db := DB.First(record, "scope_id = ?", argScopeID, "role_id = ?", argRoleID)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
