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

// GetAllOfflineUserSession is a function to get a slice of record(s) from offline_user_session table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllOfflineUserSession(ctx context.Context, page, pagesize int, order string) (results []*model.OfflineUserSession, totalRows int64, err error) {

	resultOrm := DB.Model(&model.OfflineUserSession{})
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

// GetOfflineUserSession is a function to get a single record from the offline_user_session table in the keycloak database
// error - ErrNotFound, db Find error
func GetOfflineUserSession(ctx context.Context, argUserSessionID string, argOfflineFlag string) (record *model.OfflineUserSession, err error) {
	record = &model.OfflineUserSession{}
	if err = DB.First(record, argUserSessionID, argOfflineFlag).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddOfflineUserSession is a function to add a single record to offline_user_session table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddOfflineUserSession(ctx context.Context, record *model.OfflineUserSession) (result *model.OfflineUserSession, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateOfflineUserSession is a function to update a single record from offline_user_session table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateOfflineUserSession(ctx context.Context, argUserSessionID string, argOfflineFlag string, updated *model.OfflineUserSession) (result *model.OfflineUserSession, RowsAffected int64, err error) {

	result = &model.OfflineUserSession{}
	db := DB.First(result, "user_session_id = ?", argUserSessionID, "offline_flag = ?", argOfflineFlag)
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

// DeleteOfflineUserSession is a function to delete a single record from offline_user_session table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteOfflineUserSession(ctx context.Context, argUserSessionID string, argOfflineFlag string) (rowsAffected int64, err error) {

	record := &model.OfflineUserSession{}
	db := DB.First(record, "user_session_id = ?", argUserSessionID, "offline_flag = ?", argOfflineFlag)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
