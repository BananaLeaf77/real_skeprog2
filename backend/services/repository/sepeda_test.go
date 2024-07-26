package repository

import (
	"regexp"
	"skeprogz/domain"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSqlSepedaRepo_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewSqlSepedaRepo(db)

	sepeda := &domain.Sepeda{
		Brand:    "Polygon",
		Size:     26,
		Type:     "Road",
		Quantity: 10,
	}

	mock.ExpectQuery(regexp.QuoteMeta(`
		INSERT INTO sepeda (brand, size, type, quantity, created_at, updated_at, deleted_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7) 
		RETURNING id
	`)).
		WithArgs(sepeda.Brand, sepeda.Size, sepeda.Type, sepeda.Quantity, sqlmock.AnyArg(), sqlmock.AnyArg(), nil).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	err = repo.Create(sepeda)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), sepeda.ID)
}

func TestSqlSepedaRepo_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewSqlSepedaRepo(db)

	sepeda := &domain.Sepeda{
		ID:        1,
		Brand:     "Polygon",
		Size:      26,
		Type:      "Road",
		Quantity:  10,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT id, brand, size, type, quantity, created_at, updated_at, deleted_at 
		FROM sepeda 
		WHERE id = $1
	`)).
		WithArgs(sepeda.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "brand", "size", "type", "quantity", "created_at", "updated_at", "deleted_at"}).
			AddRow(sepeda.ID, sepeda.Brand, sepeda.Size, sepeda.Type, sepeda.Quantity, sepeda.CreatedAt, sepeda.UpdatedAt, nil))

	result, err := repo.GetByID(sepeda.ID)
	assert.NoError(t, err)
	assert.Equal(t, sepeda, result)
}

func TestSqlSepedaRepo_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewSqlSepedaRepo(db)

	sepeda := &domain.Sepeda{
		ID:        1,
		Brand:     "Polygon",
		Size:      26,
		Type:      "Road",
		Quantity:  10,
		UpdatedAt: time.Now(),
	}

	mock.ExpectExec(regexp.QuoteMeta(`
		UPDATE sepeda 
		SET brand = $1, size = $2, type = $3, quantity = $4, updated_at = $5 
		WHERE id = $6
	`)).
		WithArgs(sepeda.Brand, sepeda.Size, sepeda.Type, sepeda.Quantity, sqlmock.AnyArg(), sepeda.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Update(sepeda)
	assert.NoError(t, err)
}

func TestSqlSepedaRepo_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewSqlSepedaRepo(db)

	sepedaID := uint(1)

	mock.ExpectExec(regexp.QuoteMeta(`
		UPDATE sepeda 
		SET deleted_at = CURRENT_TIMESTAMP 
		WHERE id = $1
	`)).
		WithArgs(sepedaID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Delete(sepedaID)
	assert.NoError(t, err)
}

func TestSqlSepedaRepo_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewSqlSepedaRepo(db)

	sepeda := &domain.Sepeda{
		ID:        1,
		Brand:     "Polygon",
		Size:      26,
		Type:      "Road",
		Quantity:  10,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT id, brand, size, type, quantity, created_at, updated_at, deleted_at 
		FROM sepeda
		WHERE deleted_at IS NULL
	`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "brand", "size", "type", "quantity", "created_at", "updated_at", "deleted_at"}).
			AddRow(sepeda.ID, sepeda.Brand, sepeda.Size, sepeda.Type, sepeda.Quantity, sepeda.CreatedAt, sepeda.UpdatedAt, nil))

	result, err := repo.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, len(*result))
	assert.Equal(t, sepeda, &(*result)[0])
}
