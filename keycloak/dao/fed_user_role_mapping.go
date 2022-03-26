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

// GetAllFedUserRoleMapping is a function to get a slice of record(s) from fed_user_role_mapping table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllFedUserRoleMapping(ctx context.Context, page, pagesize int, order string) (results []*model.FedUserRoleMapping, totalRows int64, err error) {

	resultOrm := DB.Model(&model.FedUserRoleMapping{})
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

// GetFedUserRoleMapping is a function to get a single record from the fed_user_role_mapping table in the keycloak database
// error - ErrNotFound, db Find error
func GetFedUserRoleMapping(ctx context.Context, argRoleID string, argUserID string) (record *model.FedUserRoleMapping, err error) {
	record = &model.FedUserRoleMapping{}
	if err = DB.First(record, argRoleID, argUserID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddFedUserRoleMapping is a function to add a single record to fed_user_role_mapping table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddFedUserRoleMapping(ctx context.Context, record *model.FedUserRoleMapping) (result *model.FedUserRoleMapping, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateFedUserRoleMapping is a function to update a single record from fed_user_role_mapping table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateFedUserRoleMapping(ctx context.Context, argRoleID string, argUserID string, updated *model.FedUserRoleMapping) (result *model.FedUserRoleMapping, RowsAffected int64, err error) {

	result = &model.FedUserRoleMapping{}
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

// DeleteFedUserRoleMapping is a function to delete a single record from fed_user_role_mapping table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteFedUserRoleMapping(ctx context.Context, argRoleID string, argUserID string) (rowsAffected int64, err error) {

	record := &model.FedUserRoleMapping{}
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
