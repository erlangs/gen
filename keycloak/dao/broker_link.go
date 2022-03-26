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

// GetAllBrokerLink is a function to get a slice of record(s) from broker_link table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllBrokerLink(ctx context.Context, page, pagesize int, order string) (results []*model.BrokerLink, totalRows int64, err error) {

	resultOrm := DB.Model(&model.BrokerLink{})
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

// GetBrokerLink is a function to get a single record from the broker_link table in the keycloak database
// error - ErrNotFound, db Find error
func GetBrokerLink(ctx context.Context, argIdentityProvider string, argUserID string) (record *model.BrokerLink, err error) {
	record = &model.BrokerLink{}
	if err = DB.First(record, argIdentityProvider, argUserID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddBrokerLink is a function to add a single record to broker_link table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddBrokerLink(ctx context.Context, record *model.BrokerLink) (result *model.BrokerLink, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateBrokerLink is a function to update a single record from broker_link table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateBrokerLink(ctx context.Context, argIdentityProvider string, argUserID string, updated *model.BrokerLink) (result *model.BrokerLink, RowsAffected int64, err error) {

	result = &model.BrokerLink{}
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

// DeleteBrokerLink is a function to delete a single record from broker_link table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteBrokerLink(ctx context.Context, argIdentityProvider string, argUserID string) (rowsAffected int64, err error) {

	record := &model.BrokerLink{}
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
