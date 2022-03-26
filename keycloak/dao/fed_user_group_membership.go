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

// GetAllFedUserGroupMembership is a function to get a slice of record(s) from fed_user_group_membership table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllFedUserGroupMembership(ctx context.Context, page, pagesize int, order string) (results []*model.FedUserGroupMembership, totalRows int64, err error) {

	resultOrm := DB.Model(&model.FedUserGroupMembership{})
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

// GetFedUserGroupMembership is a function to get a single record from the fed_user_group_membership table in the keycloak database
// error - ErrNotFound, db Find error
func GetFedUserGroupMembership(ctx context.Context, argGroupID string, argUserID string) (record *model.FedUserGroupMembership, err error) {
	record = &model.FedUserGroupMembership{}
	if err = DB.First(record, argGroupID, argUserID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddFedUserGroupMembership is a function to add a single record to fed_user_group_membership table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddFedUserGroupMembership(ctx context.Context, record *model.FedUserGroupMembership) (result *model.FedUserGroupMembership, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateFedUserGroupMembership is a function to update a single record from fed_user_group_membership table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateFedUserGroupMembership(ctx context.Context, argGroupID string, argUserID string, updated *model.FedUserGroupMembership) (result *model.FedUserGroupMembership, RowsAffected int64, err error) {

	result = &model.FedUserGroupMembership{}
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

// DeleteFedUserGroupMembership is a function to delete a single record from fed_user_group_membership table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteFedUserGroupMembership(ctx context.Context, argGroupID string, argUserID string) (rowsAffected int64, err error) {

	record := &model.FedUserGroupMembership{}
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
