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

// GetAllUserRequiredAction is a function to get a slice of record(s) from user_required_action table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllUserRequiredAction(ctx context.Context, page, pagesize int, order string) (results []*model.UserRequiredAction, totalRows int64, err error) {

	resultOrm := DB.Model(&model.UserRequiredAction{})
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

// GetUserRequiredAction is a function to get a single record from the user_required_action table in the keycloak database
// error - ErrNotFound, db Find error
func GetUserRequiredAction(ctx context.Context, argUserID string, argRequiredAction string) (record *model.UserRequiredAction, err error) {
	record = &model.UserRequiredAction{}
	if err = DB.First(record, argUserID, argRequiredAction).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddUserRequiredAction is a function to add a single record to user_required_action table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddUserRequiredAction(ctx context.Context, record *model.UserRequiredAction) (result *model.UserRequiredAction, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateUserRequiredAction is a function to update a single record from user_required_action table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateUserRequiredAction(ctx context.Context, argUserID string, argRequiredAction string, updated *model.UserRequiredAction) (result *model.UserRequiredAction, RowsAffected int64, err error) {

	result = &model.UserRequiredAction{}
	db := DB.First(result, "user_id = ?", argUserID, "required_action = ?", argRequiredAction)
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

// DeleteUserRequiredAction is a function to delete a single record from user_required_action table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteUserRequiredAction(ctx context.Context, argUserID string, argRequiredAction string) (rowsAffected int64, err error) {

	record := &model.UserRequiredAction{}
	db := DB.First(record, "user_id = ?", argUserID, "required_action = ?", argRequiredAction)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
