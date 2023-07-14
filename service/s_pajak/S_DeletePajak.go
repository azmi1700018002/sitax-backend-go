package s_pajak

import "sitax/repository/r_pajak"

type DeletePajakService interface {
	DeletePajakByID(PajakID int) error
}

type deletePajakService struct {
	deletePajakRepo r_pajak.DeletePajakRepository
}

func NewDeletePajakService(deletePajakRepo r_pajak.DeletePajakRepository) DeletePajakService {
	return &deletePajakService{deletePajakRepo}
}

func (s *deletePajakService) DeletePajakByID(PajakID int) error {
	err := s.deletePajakRepo.DeletePajakByID(PajakID)
	if err != nil {
		return err
	}
	return nil
}
