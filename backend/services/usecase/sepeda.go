package usecase

import (
	"skeprogz/domain"
)

type sepedaUseCase struct {
	sepedaRepo domain.SepedaRepository
}

func NewSepedaUseCase(sepedaRepo domain.SepedaRepository) domain.SepedaUseCase {
	return &sepedaUseCase{
		sepedaRepo: sepedaRepo,
	}
}

func (uc *sepedaUseCase) CreateUC(sepeda *domain.Sepeda) error {
	return uc.sepedaRepo.Create(sepeda)
}

func (uc *sepedaUseCase) GetByIDUC(id uint) (*domain.Sepeda, error) {
	return uc.sepedaRepo.GetByID(id)
}

func (uc *sepedaUseCase) UpdateUC(sepeda *domain.Sepeda) error {
	return uc.sepedaRepo.Update(sepeda)
}

func (uc *sepedaUseCase) DeleteUC(id uint) error {
	return uc.sepedaRepo.Delete(id)
}

func (uc *sepedaUseCase) GetAllUC() (*[]domain.Sepeda, error) {
	return uc.sepedaRepo.GetAll()
}
