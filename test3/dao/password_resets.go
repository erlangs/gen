package dao

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"time"

	"example.com/rest/example/model"
)

var (
	_ = time.Second

	_ = uuid.UUID{}
)

// GetAllPasswordResets is a function to get a slice of record(s) from password_resets table in the test1 database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllPasswordResets(ctx context.Context, page, pagesize int, order string) (results []*model.PasswordResets, totalRows int64, err error) {

	resultOrm := DB.Model(&model.PasswordResets{})
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

// GetPasswordResets is a function to get a single record from the password_resets table in the test1 database
// error - ErrNotFound, db Find error
func GetPasswordResets(ctx context.Context, argEmail string) (record *model.PasswordResets, err error) {
	record = &model.PasswordResets{}
	if err = DB.First(record, argEmail).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddPasswordResets is a function to add a single record to password_resets table in the test1 database
// error - ErrInsertFailed, db save call failed
func AddPasswordResets(ctx context.Context, record *model.PasswordResets) (result *model.PasswordResets, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdatePasswordResets is a function to update a single record from password_resets table in the test1 database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdatePasswordResets(ctx context.Context, argEmail string, updated *model.PasswordResets) (result *model.PasswordResets, RowsAffected int64, err error) {

	result = &model.PasswordResets{}
	db := DB.First(result, "email = ?", argEmail)
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

// DeletePasswordResets is a function to delete a single record from password_resets table in the test1 database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeletePasswordResets(ctx context.Context, argEmail string) (rowsAffected int64, err error) {

	record := &model.PasswordResets{}
	db := DB.First(record, "email = ?", argEmail)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
