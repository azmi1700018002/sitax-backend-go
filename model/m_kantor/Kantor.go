package m_kantor

import "sitax/model/m_user"

type Kantor struct {
	KdKantor     string `column:"kd_kantor" gorm:"type:char(3);primarykey"`
	KdInduk      string `column:"kd_induk" gorm:"type:char(3)"`
	AlamatKantor string `column:"alamat_kantor" gorm:"type:varchar(100)"`
	KotaAlamat   string `column:"kota_alamat" gorm:"type:varchar(50)"`
	NoTlp        string `column:"no_tlp" gorm:"type:varchar(20)"`
	NoFaks       string `column:"no_faks" gorm:"type:varchar(20)"`
	KdStsKantor  string `column:"kd_sts_kantor" gorm:"type:char(1)"`
	KdJnsKantor  string `column:"kd_jns_kantor" gorm:"type:char(1)"`
	KdBank1      string `column:"kd_bank1" gorm:"type:char(3)"`
	KdBank2      string `column:"kd_bank2" gorm:"type:char(3)"`
	NPWP         string `column:"npwp" gorm:"type:varchar(15)"`
	KdPejabat1   string `column:"kd_pejabat1" gorm:"type:varchar(10)"`
	KdPejabat2   string `column:"kd_pejabat2" gorm:"type:varchar(10)"`
	FlgData      string `column:"flg_data" gorm:"type:varchar(1)"`
	KdKantorLama string `column:"kd_kantor_lama" gorm:"type:varchar(2)"`
	KdLokasi     string `column:"kd_lokasi" gorm:"type:varchar(4)"`

	KdKantorfk []m_user.User `gorm:"foreignKey:KdKantor"`
}
