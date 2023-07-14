package s_kantor

import "sitax/repository/r_kantor"

type DeleteKantorService interface {
	DeleteKantorByID(KdKantor string) error
}

type deleteKantorService struct {
	deleteKantorRepo r_kantor.DeleteKantorRepository
}

func NewDeleteKantorService(deleteKantorRepo r_kantor.DeleteKantorRepository) DeleteKantorService {
	return &deleteKantorService{deleteKantorRepo}
}

func (s *deleteKantorService) DeleteKantorByID(KdKantor string) error {
	err := s.deleteKantorRepo.DeleteKantorByID(KdKantor)
	if err != nil {
		return err
	}
	return nil
}
