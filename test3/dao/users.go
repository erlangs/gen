package dao

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"time"

	"example.com/rest/example/model"
	//"github.com/satori/go.uuid"
)

var (
	_ = time.Second

	_ = uuid.UUID{}
)

// GetAllUsers_ is a function to get a slice of record(s) from users table in the test1 database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllUsers_(ctx context.Context, page, pagesize int, order string) (results []*model.Users_, totalRows int64, err error) {

	resultOrm := DB.Model(&model.Users_{})
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

// GetUsers_ is a function to get a single record from the users table in the test1 database
// error - ErrNotFound, db Find error
func GetUsers_(ctx context.Context, argID uint64) (record *model.Users_, err error) {
	record = &model.Users_{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddUsers_ is a function to add a single record to users table in the test1 database
// error - ErrInsertFailed, db save call failed
func AddUsers_(ctx context.Context, record *model.Users_) (result *model.Users_, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateUsers_ is a function to update a single record from users table in the test1 database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateUsers_(ctx context.Context, argID uint64, updated *model.Users_) (result *model.Users_, RowsAffected int64, err error) {

	result = &model.Users_{}
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

// DeleteUsers_ is a function to delete a single record from users table in the test1 database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteUsers_(ctx context.Context, argID uint64) (rowsAffected int64, err error) {

	record := &model.Users_{}
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
