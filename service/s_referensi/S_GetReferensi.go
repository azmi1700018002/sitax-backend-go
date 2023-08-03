package s_referensi

import (
	"sitax/model/m_referensi"
	"sitax/repository/r_referensi"
)

type GetReferensiService interface {
	GetAllReferensi(grpID, ref, ket string) ([]m_referensi.Referensi, error)
	GetReferensiByID(referensiID int) (*m_referensi.Referensi, error)
}

type getReferensiService struct {
	getReferensiRepository r_referensi.GetReferensiRepository
}

func NewGetReferensiService(getReferensiRepository r_referensi.GetReferensiRepository) GetReferensiService {
	return &getReferensiService{
		getReferensiRepository: getReferensiRepository,
	}
}

func (s *getReferensiService) GetAllReferensi(grpID, ref, ket string) ([]m_referensi.Referensi, error) {
	return s.getReferensiRepository.GetAllReferensi(grpID, ref, ket)
}

func (s *getReferensiService) GetReferensiByID(referensiID int) (*m_referensi.Referensi, error) {
	return s.getReferensiRepository.GetReferensiByID(referensiID)
}
