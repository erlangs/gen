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

// GetAllUserGroupMembership is a function to get a slice of record(s) from user_group_membership table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllUserGroupMembership(ctx context.Context, page, pagesize int, order string) (results []*model.UserGroupMembership, totalRows int64, err error) {

	resultOrm := DB.Model(&model.UserGroupMembership{})
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

// GetUserGroupMembership is a function to get a single record from the user_group_membership table in the keycloak database
// error - ErrNotFound, db Find error
func GetUserGroupMembership(ctx context.Context, argGroupID string, argUserID string) (record *model.UserGroupMembership, err error) {
	record = &model.UserGroupMembership{}
	if err = DB.First(record, argGroupID, argUserID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddUserGroupMembership is a function to add a single record to user_group_membership table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddUserGroupMembership(ctx context.Context, record *model.UserGroupMembership) (result *model.UserGroupMembership, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateUserGroupMembership is a function to update a single record from user_group_membership table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateUserGroupMembership(ctx context.Context, argGroupID string, argUserID string, updated *model.UserGroupMembership) (result *model.UserGroupMembership, RowsAffected int64, err error) {

	result = &model.UserGroupMembership{}
	db := DB.First(result, "group_id = ?", argGroupID, "user_id = ?", argUserID)
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

// DeleteUserGroupMembership is a function to delete a single record from user_group_membership table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteUserGroupMembership(ctx context.Context, argGroupID string, argUserID string) (rowsAffected int64, err error) {

	record := &model.UserGroupMembership{}
	db := DB.First(record, "group_id = ?", argGroupID, "user_id = ?", argUserID)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
