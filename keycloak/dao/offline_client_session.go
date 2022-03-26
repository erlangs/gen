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

// GetAllOfflineClientSession is a function to get a slice of record(s) from offline_client_session table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllOfflineClientSession(ctx context.Context, page, pagesize int, order string) (results []*model.OfflineClientSession, totalRows int64, err error) {

	resultOrm := DB.Model(&model.OfflineClientSession{})
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

// GetOfflineClientSession is a function to get a single record from the offline_client_session table in the keycloak database
// error - ErrNotFound, db Find error
func GetOfflineClientSession(ctx context.Context, argUserSessionID string, argClientID string, argOfflineFlag string, argClientStorageProvider string, argExternalClientID string) (record *model.OfflineClientSession, err error) {
	record = &model.OfflineClientSession{}
	if err = DB.First(record, argUserSessionID, argClientID, argOfflineFlag, argClientStorageProvider, argExternalClientID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddOfflineClientSession is a function to add a single record to offline_client_session table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddOfflineClientSession(ctx context.Context, record *model.OfflineClientSession) (result *model.OfflineClientSession, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateOfflineClientSession is a function to update a single record from offline_client_session table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateOfflineClientSession(ctx context.Context, argUserSessionID string, argClientID string, argOfflineFlag string, argClientStorageProvider string, argExternalClientID string, updated *model.OfflineClientSession) (result *model.OfflineClientSession, RowsAffected int64, err error) {

	result = &model.OfflineClientSession{}
	db := DB.First(result, "user_session_id = ?", argUserSessionID, "client_id = ?", argClientID, "offline_flag = ?", argOfflineFlag, "client_storage_provider = ?", argClientStorageProvider, "external_client_id = ?", argExternalClientID)
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

// DeleteOfflineClientSession is a function to delete a single record from offline_client_session table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteOfflineClientSession(ctx context.Context, argUserSessionID string, argClientID string, argOfflineFlag string, argClientStorageProvider string, argExternalClientID string) (rowsAffected int64, err error) {

	record := &model.OfflineClientSession{}
	db := DB.First(record, "user_session_id = ?", argUserSessionID, "client_id = ?", argClientID, "offline_flag = ?", argOfflineFlag, "client_storage_provider = ?", argClientStorageProvider, "external_client_id = ?", argExternalClientID)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
