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

// GetAllUserConsent is a function to get a slice of record(s) from user_consent table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllUserConsent(ctx context.Context, page, pagesize int, order string) (results []*model.UserConsent, totalRows int64, err error) {

	resultOrm := DB.Model(&model.UserConsent{})
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

// GetUserConsent is a function to get a single record from the user_consent table in the keycloak database
// error - ErrNotFound, db Find error
func GetUserConsent(ctx context.Context, argID string) (record *model.UserConsent, err error) {
	record = &model.UserConsent{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddUserConsent is a function to add a single record to user_consent table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddUserConsent(ctx context.Context, record *model.UserConsent) (result *model.UserConsent, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateUserConsent is a function to update a single record from user_consent table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateUserConsent(ctx context.Context, argID string, updated *model.UserConsent) (result *model.UserConsent, RowsAffected int64, err error) {

	result = &model.UserConsent{}
	db := DB.First(result, "id = ?", argID)
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

// DeleteUserConsent is a function to delete a single record from user_consent table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteUserConsent(ctx context.Context, argID string) (rowsAffected int64, err error) {

	record := &model.UserConsent{}
	db := DB.First(record, "id = ?", argID)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
