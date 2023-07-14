package r_kantor

import (
	"sitax/config/db"
	"sitax/model/m_kantor"
)

type UpdateKantorRepository interface {
	UpdateKantor(kantor *m_kantor.Kantor) (*m_kantor.Kantor, error)
}

type updateKantorRepository struct{}

func NewUpdateKantorRepository() UpdateKantorRepository {
	return &updateKantorRepository{}
}

func (r *updateKantorRepository) UpdateKantor(kantor *m_kantor.Kantor) (*m_kantor.Kantor, error) {
	// Check if the user exists
	var existingKantor m_kantor.Kantor
	if err := db.Server().Where("kd_kantor = ?", kantor.KdKantor).First(&existingKantor).Error; err != nil {
		return nil, err
	}

	// Update user data in the database
	if err := db.Server().Model(&m_kantor.Kantor{}).
		Where("kd_kantor = ?", kantor.KdKantor).
		Updates(map[string]interface{}{
			"kd_induk":       kantor.KdInduk,
			"alamat_kantor":  kantor.AlamatKantor,
			"kota_alamat":    kantor.KotaAlamat,
			"no_tlp":         kantor.NoTlp,
			"no_faks":        kantor.NoFaks,
			"kd_sts_kantor":  kantor.KdStsKantor,
			"kd_jns_kantor":  kantor.KdJnsKantor,
			"kd_bank1":       kantor.KdBank1,
			"kd_bank2":       kantor.KdBank2,
			"npwp":           kantor.NPWP,
			"kd_pejabat1":    kantor.KdPejabat1,
			"kd_pejabat2":    kantor.KdPejabat2,
			"flg_data":       kantor.FlgData,
			"kd_kantor_lama": kantor.KdKantorLama,
			"kd_lokasi":      kantor.KdLokasi,
		}).Error; err != nil {
		return nil, err
	}
	return kantor, nil
}
