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

// GetAllUsernameLoginFailure is a function to get a slice of record(s) from username_login_failure table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllUsernameLoginFailure(ctx context.Context, page, pagesize int, order string) (results []*model.UsernameLoginFailure, totalRows int64, err error) {

	resultOrm := DB.Model(&model.UsernameLoginFailure{})
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

// GetUsernameLoginFailure is a function to get a single record from the username_login_failure table in the keycloak database
// error - ErrNotFound, db Find error
func GetUsernameLoginFailure(ctx context.Context, argRealmID string, argUsername string) (record *model.UsernameLoginFailure, err error) {
	record = &model.UsernameLoginFailure{}
	if err = DB.First(record, argRealmID, argUsername).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddUsernameLoginFailure is a function to add a single record to username_login_failure table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddUsernameLoginFailure(ctx context.Context, record *model.UsernameLoginFailure) (result *model.UsernameLoginFailure, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateUsernameLoginFailure is a function to update a single record from username_login_failure table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateUsernameLoginFailure(ctx context.Context, argRealmID string, argUsername string, updated *model.UsernameLoginFailure) (result *model.UsernameLoginFailure, RowsAffected int64, err error) {

	result = &model.UsernameLoginFailure{}
	db := DB.First(result, "realm_id = ?", argRealmID, "username = ?", argUsername)
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

// DeleteUsernameLoginFailure is a function to delete a single record from username_login_failure table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteUsernameLoginFailure(ctx context.Context, argRealmID string, argUsername string) (rowsAffected int64, err error) {

	record := &model.UsernameLoginFailure{}
	db := DB.First(record, "realm_id = ?", argRealmID, "username = ?", argUsername)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
