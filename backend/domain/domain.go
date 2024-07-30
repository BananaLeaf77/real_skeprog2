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

type UpdateHistory struct {
	ID          uint      `json:"id"`
	SepedaID    uint      `json:"sepeda_id"`
	OldSize     int       `json:"old_size"`
	OldType     string    `json:"old_type"`
	OldQuantity int       `json:"old_quantity"`
	NewSize     int       `json:"new_size"`
	NewType     string    `json:"new_type"`
	NewQuantity int       `json:"new_quantity"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (s *Sepeda) GetBrand() string {
	return s.Brand
}

func (s *Sepeda) SetBrand(brand string) {
	s.Brand = brand
}

func (s *Sepeda) GetSize() int {
	return s.Size
}

func (s *Sepeda) SetSize(size int) {
	s.Size = size
}

type SepedaListrik struct {
	Brand    string
	quantity int
}

var sepedaL = SepedaListrik{
	Brand:    "Brodi",
	quantity: 12,
}

func (s *SepedaListrik) GetQuantity() int {
	return s.quantity
}

func (s *SepedaListrik) SetQuantity(quantity int) {
	if quantity >= 0 {
		s.quantity = quantity
	}
}

type SepedaRepository interface {
	Create(sepeda *Sepeda) error
	GetByID(id uint) (*Sepeda, error)
	Update(sepeda *Sepeda) error
	Delete(id uint) error
	GetAll() (*[]Sepeda, error)
	GetAllUpdateHistory() (*[]UpdateHistory, error)
}

type SepedaUseCase interface {
	CreateUC(sepeda *Sepeda) error
	GetByIDUC(id uint) (*Sepeda, error)
	UpdateUC(sepeda *Sepeda) error
	DeleteUC(id uint) error
	GetAllUC() (*[]Sepeda, error)
	GetAllUpdateHistoryUC() (*[]UpdateHistory, error)
}
