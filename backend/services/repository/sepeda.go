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
	query2 := `SELECT id, brand, size, type, quantity 
	           FROM sepeda 
	           WHERE brand = $1 AND type = $2 AND deleted_at IS NULL`

	var existingSepeda domain.Sepeda

	err := r.db.QueryRow(query2, sepeda.Brand, sepeda.Type).Scan(&existingSepeda.ID, &existingSepeda.Brand, &existingSepeda.Size, &existingSepeda.Type, &existingSepeda.Quantity)

	if err == sql.ErrNoRows {
		query := `
			INSERT INTO sepeda (brand, size, type, quantity, created_at, updated_at, deleted_at) 
			VALUES ($1, $2, $3, $4, $5, $6, $7) 
			RETURNING id
		`
		err := r.db.QueryRow(query, sepeda.Brand, sepeda.Size, sepeda.Type, sepeda.Quantity, time.Now(), time.Now(), nil).Scan(&sepeda.ID)
		if err != nil {
			return fmt.Errorf("error inserting new sepeda: %w", err)
		}
		return nil
	} else if err != nil {
		return fmt.Errorf("error querying sepeda: %w", err)
	}

	if existingSepeda.Brand == sepeda.Brand && existingSepeda.Type == sepeda.Type {
		penambahanValue := sepeda.Quantity + existingSepeda.Quantity

		updateQuery := `
		UPDATE sepeda 
		SET size = $1, quantity = $2, updated_at = $3 
		WHERE id = $4
		`
		_, err = r.db.Exec(updateQuery, sepeda.Size, penambahanValue, time.Now(), existingSepeda.ID)
		if err != nil {
			return fmt.Errorf("error updating sepeda: %w", err)
		}

		historyQuery := `
			INSERT INTO update_history (sepeda_id, old_size, old_type, old_quantity, new_size, new_type, new_quantity, updated_at) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		`
		_, err = r.db.Exec(historyQuery, existingSepeda.ID, existingSepeda.Size, existingSepeda.Type, existingSepeda.Quantity, sepeda.Size, sepeda.Type, sepeda.Quantity, time.Now())
		if err != nil {
			return fmt.Errorf("error logging update history: %w", err)
		}
	}

	return nil
}

func (r *sqlSepedaRepo) GetByID(id uint) (*domain.Sepeda, error) {
	sepeda := &domain.Sepeda{}
	query := `
		SELECT id, brand, size, type, quantity, created_at, updated_at, deleted_at 
		FROM sepeda 
		WHERE id = $1 AND deleted_at IS NULL
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
	if err != nil {
		return err
	}
	return nil
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

	return &sepedaList, nil
}

func (r *sqlSepedaRepo) GetAllUpdateHistory() (*[]domain.UpdateHistory, error) {
	query := `
		SELECT id, sepeda_id, old_size, old_type, old_quantity, new_size, new_type, new_quantity, updated_at 
		FROM update_history
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var historyList []domain.UpdateHistory
	for rows.Next() {
		var history domain.UpdateHistory
		err := rows.Scan(&history.ID, &history.SepedaID, &history.OldSize, &history.OldType, &history.OldQuantity, &history.NewSize, &history.NewType, &history.NewQuantity, &history.UpdatedAt)
		if err != nil {
			return nil, err
		}
		historyList = append(historyList, history)
	}

	return &historyList, nil
}
