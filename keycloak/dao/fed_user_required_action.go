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

// GetAllFedUserRequiredAction is a function to get a slice of record(s) from fed_user_required_action table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllFedUserRequiredAction(ctx context.Context, page, pagesize int, order string) (results []*model.FedUserRequiredAction, totalRows int64, err error) {

	resultOrm := DB.Model(&model.FedUserRequiredAction{})
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

// GetFedUserRequiredAction is a function to get a single record from the fed_user_required_action table in the keycloak database
// error - ErrNotFound, db Find error
func GetFedUserRequiredAction(ctx context.Context, argRequiredAction string, argUserID string) (record *model.FedUserRequiredAction, err error) {
	record = &model.FedUserRequiredAction{}
	if err = DB.First(record, argRequiredAction, argUserID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddFedUserRequiredAction is a function to add a single record to fed_user_required_action table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddFedUserRequiredAction(ctx context.Context, record *model.FedUserRequiredAction) (result *model.FedUserRequiredAction, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateFedUserRequiredAction is a function to update a single record from fed_user_required_action table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateFedUserRequiredAction(ctx context.Context, argRequiredAction string, argUserID string, updated *model.FedUserRequiredAction) (result *model.FedUserRequiredAction, RowsAffected int64, err error) {

	result = &model.FedUserRequiredAction{}
	db := DB.First(result, "required_action = ?", argRequiredAction, "user_id = ?", argUserID)
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

// DeleteFedUserRequiredAction is a function to delete a single record from fed_user_required_action table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteFedUserRequiredAction(ctx context.Context, argRequiredAction string, argUserID string) (rowsAffected int64, err error) {

	record := &model.FedUserRequiredAction{}
	db := DB.First(record, "required_action = ?", argRequiredAction, "user_id = ?", argUserID)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
