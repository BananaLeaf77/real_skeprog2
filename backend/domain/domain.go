package domain

import "time"

type Sepeda struct {
	ID        uint       `json:"id"`
	Brand     string     `json:"brand"`
	Size      int        `json:"size"`
	Type      string     `json:"type"`
	Quantity  int        `json:"quantity"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type SepedaRepository interface {
	Create(sepeda *Sepeda) error
	GetByID(id uint) (*Sepeda, error)
	Update(sepeda *Sepeda) error
	Delete(id uint) error
	GetAll() (*[]Sepeda, error)
}

type SepedaUseCase interface {
	CreateUC(sepeda *Sepeda) error
	GetByIDUC(id uint) (*Sepeda, error)
	UpdateUC(sepeda *Sepeda) error
	DeleteUC(id uint) error
	GetAllUC() (*[]Sepeda, error)
}
