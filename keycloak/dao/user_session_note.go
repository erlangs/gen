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

// GetAllUserSessionNote is a function to get a slice of record(s) from user_session_note table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllUserSessionNote(ctx context.Context, page, pagesize int, order string) (results []*model.UserSessionNote, totalRows int64, err error) {

	resultOrm := DB.Model(&model.UserSessionNote{})
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

// GetUserSessionNote is a function to get a single record from the user_session_note table in the keycloak database
// error - ErrNotFound, db Find error
func GetUserSessionNote(ctx context.Context, argUserSession string, argName string) (record *model.UserSessionNote, err error) {
	record = &model.UserSessionNote{}
	if err = DB.First(record, argUserSession, argName).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddUserSessionNote is a function to add a single record to user_session_note table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddUserSessionNote(ctx context.Context, record *model.UserSessionNote) (result *model.UserSessionNote, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateUserSessionNote is a function to update a single record from user_session_note table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateUserSessionNote(ctx context.Context, argUserSession string, argName string, updated *model.UserSessionNote) (result *model.UserSessionNote, RowsAffected int64, err error) {

	result = &model.UserSessionNote{}
	db := DB.First(result, "user_session = ?", argUserSession, "name = ?", argName)
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

// DeleteUserSessionNote is a function to delete a single record from user_session_note table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteUserSessionNote(ctx context.Context, argUserSession string, argName string) (rowsAffected int64, err error) {

	record := &model.UserSessionNote{}
	db := DB.First(record, "user_session = ?", argUserSession, "name = ?", argName)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
