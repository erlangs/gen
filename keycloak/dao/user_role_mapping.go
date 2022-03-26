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

// GetAllUserRoleMapping is a function to get a slice of record(s) from user_role_mapping table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllUserRoleMapping(ctx context.Context, page, pagesize int, order string) (results []*model.UserRoleMapping, totalRows int64, err error) {

	resultOrm := DB.Model(&model.UserRoleMapping{})
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

// GetUserRoleMapping is a function to get a single record from the user_role_mapping table in the keycloak database
// error - ErrNotFound, db Find error
func GetUserRoleMapping(ctx context.Context, argRoleID string, argUserID string) (record *model.UserRoleMapping, err error) {
	record = &model.UserRoleMapping{}
	if err = DB.First(record, argRoleID, argUserID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddUserRoleMapping is a function to add a single record to user_role_mapping table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddUserRoleMapping(ctx context.Context, record *model.UserRoleMapping) (result *model.UserRoleMapping, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateUserRoleMapping is a function to update a single record from user_role_mapping table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateUserRoleMapping(ctx context.Context, argRoleID string, argUserID string, updated *model.UserRoleMapping) (result *model.UserRoleMapping, RowsAffected int64, err error) {

	result = &model.UserRoleMapping{}
	db := DB.First(result, "role_id = ?", argRoleID, "user_id = ?", argUserID)
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

// DeleteUserRoleMapping is a function to delete a single record from user_role_mapping table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteUserRoleMapping(ctx context.Context, argRoleID string, argUserID string) (rowsAffected int64, err error) {

	record := &model.UserRoleMapping{}
	db := DB.First(record, "role_id = ?", argRoleID, "user_id = ?", argUserID)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
