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

// GetAllGroupRoleMapping is a function to get a slice of record(s) from group_role_mapping table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllGroupRoleMapping(ctx context.Context, page, pagesize int, order string) (results []*model.GroupRoleMapping, totalRows int64, err error) {

	resultOrm := DB.Model(&model.GroupRoleMapping{})
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

// GetGroupRoleMapping is a function to get a single record from the group_role_mapping table in the keycloak database
// error - ErrNotFound, db Find error
func GetGroupRoleMapping(ctx context.Context, argRoleID string, argGroupID string) (record *model.GroupRoleMapping, err error) {
	record = &model.GroupRoleMapping{}
	if err = DB.First(record, argRoleID, argGroupID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddGroupRoleMapping is a function to add a single record to group_role_mapping table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddGroupRoleMapping(ctx context.Context, record *model.GroupRoleMapping) (result *model.GroupRoleMapping, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateGroupRoleMapping is a function to update a single record from group_role_mapping table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateGroupRoleMapping(ctx context.Context, argRoleID string, argGroupID string, updated *model.GroupRoleMapping) (result *model.GroupRoleMapping, RowsAffected int64, err error) {

	result = &model.GroupRoleMapping{}
	db := DB.First(result, "role_id = ?", argRoleID, "group_id = ?", argGroupID)
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

// DeleteGroupRoleMapping is a function to delete a single record from group_role_mapping table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteGroupRoleMapping(ctx context.Context, argRoleID string, argGroupID string) (rowsAffected int64, err error) {

	record := &model.GroupRoleMapping{}
	db := DB.First(record, "role_id = ?", argRoleID, "group_id = ?", argGroupID)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
