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

// GetAllPersonalAccessTokens is a function to get a slice of record(s) from personal_access_tokens table in the test1 database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllPersonalAccessTokens(ctx context.Context, page, pagesize int, order string) (results []*model.PersonalAccessTokens, totalRows int64, err error) {

	resultOrm := DB.Model(&model.PersonalAccessTokens{})
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

// GetPersonalAccessTokens is a function to get a single record from the personal_access_tokens table in the test1 database
// error - ErrNotFound, db Find error
func GetPersonalAccessTokens(ctx context.Context, argID uint64) (record *model.PersonalAccessTokens, err error) {
	record = &model.PersonalAccessTokens{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddPersonalAccessTokens is a function to add a single record to personal_access_tokens table in the test1 database
// error - ErrInsertFailed, db save call failed
func AddPersonalAccessTokens(ctx context.Context, record *model.PersonalAccessTokens) (result *model.PersonalAccessTokens, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdatePersonalAccessTokens is a function to update a single record from personal_access_tokens table in the test1 database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdatePersonalAccessTokens(ctx context.Context, argID uint64, updated *model.PersonalAccessTokens) (result *model.PersonalAccessTokens, RowsAffected int64, err error) {

	result = &model.PersonalAccessTokens{}
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

// DeletePersonalAccessTokens is a function to delete a single record from personal_access_tokens table in the test1 database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeletePersonalAccessTokens(ctx context.Context, argID uint64) (rowsAffected int64, err error) {

	record := &model.PersonalAccessTokens{}
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
