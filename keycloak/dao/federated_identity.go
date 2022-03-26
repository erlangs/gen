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

// GetAllFederatedIdentity is a function to get a slice of record(s) from federated_identity table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllFederatedIdentity(ctx context.Context, page, pagesize int, order string) (results []*model.FederatedIdentity, totalRows int64, err error) {

	resultOrm := DB.Model(&model.FederatedIdentity{})
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

// GetFederatedIdentity is a function to get a single record from the federated_identity table in the keycloak database
// error - ErrNotFound, db Find error
func GetFederatedIdentity(ctx context.Context, argIdentityProvider string, argUserID string) (record *model.FederatedIdentity, err error) {
	record = &model.FederatedIdentity{}
	if err = DB.First(record, argIdentityProvider, argUserID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddFederatedIdentity is a function to add a single record to federated_identity table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddFederatedIdentity(ctx context.Context, record *model.FederatedIdentity) (result *model.FederatedIdentity, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateFederatedIdentity is a function to update a single record from federated_identity table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateFederatedIdentity(ctx context.Context, argIdentityProvider string, argUserID string, updated *model.FederatedIdentity) (result *model.FederatedIdentity, RowsAffected int64, err error) {

	result = &model.FederatedIdentity{}
	db := DB.First(result, "identity_provider = ?", argIdentityProvider, "user_id = ?", argUserID)
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

// DeleteFederatedIdentity is a function to delete a single record from federated_identity table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteFederatedIdentity(ctx context.Context, argIdentityProvider string, argUserID string) (rowsAffected int64, err error) {

	record := &model.FederatedIdentity{}
	db := DB.First(record, "identity_provider = ?", argIdentityProvider, "user_id = ?", argUserID)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
