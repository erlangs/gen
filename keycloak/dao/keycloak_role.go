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

// GetAllKeycloakRole is a function to get a slice of record(s) from keycloak_role table in the keycloak database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllKeycloakRole(ctx context.Context, page, pagesize int, order string) (results []*model.KeycloakRole, totalRows int64, err error) {

	resultOrm := DB.Model(&model.KeycloakRole{})
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

// GetKeycloakRole is a function to get a single record from the keycloak_role table in the keycloak database
// error - ErrNotFound, db Find error
func GetKeycloakRole(ctx context.Context, argID string) (record *model.KeycloakRole, err error) {
	record = &model.KeycloakRole{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddKeycloakRole is a function to add a single record to keycloak_role table in the keycloak database
// error - ErrInsertFailed, db save call failed
func AddKeycloakRole(ctx context.Context, record *model.KeycloakRole) (result *model.KeycloakRole, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateKeycloakRole is a function to update a single record from keycloak_role table in the keycloak database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateKeycloakRole(ctx context.Context, argID string, updated *model.KeycloakRole) (result *model.KeycloakRole, RowsAffected int64, err error) {

	result = &model.KeycloakRole{}
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

// DeleteKeycloakRole is a function to delete a single record from keycloak_role table in the keycloak database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteKeycloakRole(ctx context.Context, argID string) (rowsAffected int64, err error) {

	record := &model.KeycloakRole{}
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
