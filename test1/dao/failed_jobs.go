package dao

import (
	"context"
	"time"

	"example.com/rest/example/model"

	uuid "github.com/satori/go.uuid"
)

var (
	_ = time.Second

	_ = uuid.UUID{}
)

// GetAllFailedJobs is a function to get a slice of record(s) from failed_jobs table in the test1 database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllFailedJobs(ctx context.Context, page, pagesize int, order string) (results []*model.FailedJobs, totalRows int64, err error) {

	resultOrm := DB.Model(&model.FailedJobs{})
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

// GetFailedJobs is a function to get a single record from the failed_jobs table in the test1 database
// error - ErrNotFound, db Find error
func GetFailedJobs(ctx context.Context, argID uint64) (record *model.FailedJobs, err error) {
	record = &model.FailedJobs{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddFailedJobs is a function to add a single record to failed_jobs table in the test1 database
// error - ErrInsertFailed, db save call failed
func AddFailedJobs(ctx context.Context, record *model.FailedJobs) (result *model.FailedJobs, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateFailedJobs is a function to update a single record from failed_jobs table in the test1 database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateFailedJobs(ctx context.Context, argID uint64, updated *model.FailedJobs) (result *model.FailedJobs, RowsAffected int64, err error) {

	result = &model.FailedJobs{}
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

// DeleteFailedJobs is a function to delete a single record from failed_jobs table in the test1 database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteFailedJobs(ctx context.Context, argID uint64) (rowsAffected int64, err error) {

	record := &model.FailedJobs{}
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
