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

// GetAllFedUserConsentClScope is a function to get a slice of record(s) from fed_user_consent_cl_scope table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllFedUserConsentClScope(ctx context.Context, page, pagesize int, order string) (results []*model.FedUserConsentClScope, totalRows int64, err error) {

	resultOrm := DB.Model(&model.FedUserConsentClScope{})
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

// GetFedUserConsentClScope is a function to get a single record from the fed_user_consent_cl_scope table in the keycloak database
// error - ErrNotFound, db Find error
func GetFedUserConsentClScope(ctx context.Context, argUserConsentID string, argScopeID string) (record *model.FedUserConsentClScope, err error) {
	record = &model.FedUserConsentClScope{}
	if err = DB.First(record, argUserConsentID, argScopeID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddFedUserConsentClScope is a function to add a single record to fed_user_consent_cl_scope table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddFedUserConsentClScope(ctx context.Context, record *model.FedUserConsentClScope) (result *model.FedUserConsentClScope, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateFedUserConsentClScope is a function to update a single record from fed_user_consent_cl_scope table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateFedUserConsentClScope(ctx context.Context, argUserConsentID string, argScopeID string, updated *model.FedUserConsentClScope) (result *model.FedUserConsentClScope, RowsAffected int64, err error) {

	result = &model.FedUserConsentClScope{}
	db := DB.First(result, "user_consent_id = ?", argUserConsentID, "scope_id = ?", argScopeID)
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

// DeleteFedUserConsentClScope is a function to delete a single record from fed_user_consent_cl_scope table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteFedUserConsentClScope(ctx context.Context, argUserConsentID string, argScopeID string) (rowsAffected int64, err error) {

	record := &model.FedUserConsentClScope{}
	db := DB.First(record, "user_consent_id = ?", argUserConsentID, "scope_id = ?", argScopeID)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
