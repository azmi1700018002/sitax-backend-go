package s_pajak_detail

import (
	"sitax/repository/r_pajak_detail"
)

type DeletePajakDetailService interface {
	DeletePajakDetailByID(PajakDetailID int) error
}

type deletePajakDetailService struct {
	deletePajakDetailRepo r_pajak_detail.DeletePajakDetailRepository
}

func NewDeletePajakDetailService(deletePajakDetailRepo r_pajak_detail.DeletePajakDetailRepository) DeletePajakDetailService {
	return &deletePajakDetailService{deletePajakDetailRepo}
}

func (s *deletePajakDetailService) DeletePajakDetailByID(PajakDetailID int) error {
	err := s.deletePajakDetailRepo.DeletePajakDetailByID(PajakDetailID)
	if err != nil {
		return err
	}
	return nil
}
