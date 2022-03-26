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

// GetAllFedUserConsent is a function to get a slice of record(s) from fed_user_consent table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllFedUserConsent(ctx context.Context, page, pagesize int, order string) (results []*model.FedUserConsent, totalRows int64, err error) {

	resultOrm := DB.Model(&model.FedUserConsent{})
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

// GetFedUserConsent is a function to get a single record from the fed_user_consent table in the keycloak database
// error - ErrNotFound, db Find error
func GetFedUserConsent(ctx context.Context, argID string) (record *model.FedUserConsent, err error) {
	record = &model.FedUserConsent{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddFedUserConsent is a function to add a single record to fed_user_consent table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddFedUserConsent(ctx context.Context, record *model.FedUserConsent) (result *model.FedUserConsent, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateFedUserConsent is a function to update a single record from fed_user_consent table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateFedUserConsent(ctx context.Context, argID string, updated *model.FedUserConsent) (result *model.FedUserConsent, RowsAffected int64, err error) {

	result = &model.FedUserConsent{}
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

// DeleteFedUserConsent is a function to delete a single record from fed_user_consent table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteFedUserConsent(ctx context.Context, argID string) (rowsAffected int64, err error) {

	record := &model.FedUserConsent{}
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
