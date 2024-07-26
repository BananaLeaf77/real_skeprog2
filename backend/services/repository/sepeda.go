package repository

import (
	"database/sql"
	"fmt"
	"skeprogz/domain"
	"time"
)

type sqlSepedaRepo struct {
	db *sql.DB
}

func NewSqlSepedaRepo(db *sql.DB) domain.SepedaRepository {
	return &sqlSepedaRepo{
		db: db,
	}
}

func (r *sqlSepedaRepo) Create(sepeda *domain.Sepeda) error {
	query := `
		INSERT INTO sepeda (brand, size, type, quantity, created_at, updated_at, deleted_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7) 
		RETURNING id
	`
	err := r.db.QueryRow(query, sepeda.Brand, sepeda.Size, sepeda.Type, sepeda.Quantity, time.Now(), time.Now(), nil).Scan(&sepeda.ID)
	if err != nil {
		fmt.Println("ada err")
		return err
	} else {
		return nil
	}
}

func (r *sqlSepedaRepo) GetByID(id uint) (*domain.Sepeda, error) {
	sepeda := &domain.Sepeda{}
	query := `
		SELECT id, brand, size, type, quantity, created_at, updated_at, deleted_at 
		FROM sepeda 
		WHERE id = $1
	`
	err := r.db.QueryRow(query, id).Scan(&sepeda.ID, &sepeda.Brand, &sepeda.Size, &sepeda.Type, &sepeda.Quantity, &sepeda.CreatedAt, &sepeda.UpdatedAt, &sepeda.DeletedAt)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("sepeda not found")
	}
	if err != nil {
		return nil, err
	}
	return sepeda, nil
}

func (r *sqlSepedaRepo) Update(sepeda *domain.Sepeda) error {
	query := `
		UPDATE sepeda 
		SET brand = $1, size = $2, type = $3, quantity = $4, updated_at = $5 
		WHERE id = $6
	`
	_, err := r.db.Exec(query, sepeda.Brand, sepeda.Size, sepeda.Type, sepeda.Quantity, time.Now(), sepeda.ID)
	return err
}

func (r *sqlSepedaRepo) Delete(id uint) error {
	query := `
		UPDATE sepeda 
		SET deleted_at = CURRENT_TIMESTAMP 
		WHERE id = $1
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *sqlSepedaRepo) GetAll() (*[]domain.Sepeda, error) {
	query := `
		SELECT id, brand, size, type, quantity, created_at, updated_at, deleted_at 
		FROM sepeda
		WHERE deleted_at IS NULL
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sepedaList []domain.Sepeda
	for rows.Next() {
		var sepeda domain.Sepeda
		err := rows.Scan(&sepeda.ID, &sepeda.Brand, &sepeda.Size, &sepeda.Type, &sepeda.Quantity, &sepeda.CreatedAt, &sepeda.UpdatedAt, &sepeda.DeletedAt)
		if err != nil {
			return nil, err
		}
		sepedaList = append(sepedaList, sepeda)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &sepedaList, nil
}
