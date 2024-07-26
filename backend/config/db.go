package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var sqlDB *sql.DB

func GetDatabaseURL() string {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"))
	return dsn
}

func BootDB() (*sql.DB, error) {
	url := GetDatabaseURL()
	fmt.Println("Connecting to database with URL:", url)
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	if sqlDB == nil {
		sqlDB = db
	}

	err = Migrate(sqlDB)
	if err != nil {
		return sqlDB, err
	}

	return sqlDB, nil
}

func Migrate(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS sepeda (
        id SERIAL PRIMARY KEY,
        brand VARCHAR(50),
        size INTEGER,
        type VARCHAR(100),
        quantity INTEGER,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        deleted_at TIMESTAMP WITH TIME ZONE
    );
    `
	_, err := db.Exec(query)
	if err != nil {
		fmt.Printf("Error executing migration query: %v\n", err)
		return fmt.Errorf("failed to run migrations: %w", err)
	}
	return nil
}
