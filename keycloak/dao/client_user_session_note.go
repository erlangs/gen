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

// GetAllClientUserSessionNote is a function to get a slice of record(s) from client_user_session_note table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllClientUserSessionNote(ctx context.Context, page, pagesize int, order string) (results []*model.ClientUserSessionNote, totalRows int64, err error) {

	resultOrm := DB.Model(&model.ClientUserSessionNote{})
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

// GetClientUserSessionNote is a function to get a single record from the client_user_session_note table in the keycloak database
// error - ErrNotFound, db Find error
func GetClientUserSessionNote(ctx context.Context, argName string, argClientSession string) (record *model.ClientUserSessionNote, err error) {
	record = &model.ClientUserSessionNote{}
	if err = DB.First(record, argName, argClientSession).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddClientUserSessionNote is a function to add a single record to client_user_session_note table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddClientUserSessionNote(ctx context.Context, record *model.ClientUserSessionNote) (result *model.ClientUserSessionNote, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateClientUserSessionNote is a function to update a single record from client_user_session_note table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateClientUserSessionNote(ctx context.Context, argName string, argClientSession string, updated *model.ClientUserSessionNote) (result *model.ClientUserSessionNote, RowsAffected int64, err error) {

	result = &model.ClientUserSessionNote{}
	db := DB.First(result, "name = ?", argName, "client_session = ?", argClientSession)
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

// DeleteClientUserSessionNote is a function to delete a single record from client_user_session_note table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteClientUserSessionNote(ctx context.Context, argName string, argClientSession string) (rowsAffected int64, err error) {

	record := &model.ClientUserSessionNote{}
	db := DB.First(record, "name = ?", argName, "client_session = ?", argClientSession)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
