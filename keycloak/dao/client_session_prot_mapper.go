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

// GetAllClientSessionProtMapper is a function to get a slice of record(s) from client_session_prot_mapper table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllClientSessionProtMapper(ctx context.Context, page, pagesize int, order string) (results []*model.ClientSessionProtMapper, totalRows int64, err error) {

	resultOrm := DB.Model(&model.ClientSessionProtMapper{})
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

// GetClientSessionProtMapper is a function to get a single record from the client_session_prot_mapper table in the keycloak database
// error - ErrNotFound, db Find error
func GetClientSessionProtMapper(ctx context.Context, argProtocolMapperID string, argClientSession string) (record *model.ClientSessionProtMapper, err error) {
	record = &model.ClientSessionProtMapper{}
	if err = DB.First(record, argProtocolMapperID, argClientSession).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddClientSessionProtMapper is a function to add a single record to client_session_prot_mapper table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddClientSessionProtMapper(ctx context.Context, record *model.ClientSessionProtMapper) (result *model.ClientSessionProtMapper, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateClientSessionProtMapper is a function to update a single record from client_session_prot_mapper table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateClientSessionProtMapper(ctx context.Context, argProtocolMapperID string, argClientSession string, updated *model.ClientSessionProtMapper) (result *model.ClientSessionProtMapper, RowsAffected int64, err error) {

	result = &model.ClientSessionProtMapper{}
	db := DB.First(result, "protocol_mapper_id = ?", argProtocolMapperID, "client_session = ?", argClientSession)
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

// DeleteClientSessionProtMapper is a function to delete a single record from client_session_prot_mapper table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteClientSessionProtMapper(ctx context.Context, argProtocolMapperID string, argClientSession string) (rowsAffected int64, err error) {

	record := &model.ClientSessionProtMapper{}
	db := DB.First(record, "protocol_mapper_id = ?", argProtocolMapperID, "client_session = ?", argClientSession)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
