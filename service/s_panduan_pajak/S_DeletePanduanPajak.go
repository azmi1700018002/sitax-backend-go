package s_panduan_pajak

import "sitax/repository/r_panduan_pajak"

type DeletePanduanPajakService interface {
	DeletePanduanPajakByID(PanduanPajakID int) error
}

type deletePanduanPajakService struct {
	deletePanduanPajakRepo r_panduan_pajak.DeletePanduanPajakRepository
}

func NewDeletePanduanPajakService(deletePanduanPajakRepo r_panduan_pajak.DeletePanduanPajakRepository) DeletePanduanPajakService {
	return &deletePanduanPajakService{deletePanduanPajakRepo}
}

func (s *deletePanduanPajakService) DeletePanduanPajakByID(PanduanPajakID int) error {
	err := s.deletePanduanPajakRepo.DeletePanduanPajakByID(PanduanPajakID)
	if err != nil {
		return err
	}
	return nil
}
