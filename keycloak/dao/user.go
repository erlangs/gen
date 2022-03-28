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

// GetAllUser is a function to get a slice of record(s) from user table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllUser(ctx context.Context, page, pagesize int, order string) (results []*model.User, totalRows int64, err error) {

	resultOrm := DB.Model(&model.User{})
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

// GetUser is a function to get a single record from the user table in the keycloak database
// error - ErrNotFound, db Find error
func GetUser(ctx context.Context, argID int64) (record *model.User, err error) {
	record = &model.User{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddUser is a function to add a single record to user table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddUser(ctx context.Context, record *model.User) (result *model.User, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateUser is a function to update a single record from user table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateUser(ctx context.Context, argID int64, updated *model.User) (result *model.User, RowsAffected int64, err error) {

	result = &model.User{}
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

// DeleteUser is a function to delete a single record from user table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteUser(ctx context.Context, argID int64) (rowsAffected int64, err error) {

	record := &model.User{}
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
