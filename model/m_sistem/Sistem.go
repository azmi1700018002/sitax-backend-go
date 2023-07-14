package m_sistem

type Sistem struct {
	SistemID    int    `column:"sistem_id"`
	SistemNama  string `column:"sistem_nama" gorm:"type:varchar(100)"`
	SistemDesc  string `column:"sistem_desc" gorm:"type:varchar(100)"`
	SistemVersi string `column:"sistem_versi" gorm:"type:varchar(30)"`
}
