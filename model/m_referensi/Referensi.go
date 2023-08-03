package m_referensi

type Referensi struct {
	IDref  int    `column:"id_referensi" gorm:"primarykey"`
	GrpID  string `column:"grp_id" gorm:"type:varchar(50)"`
	Ref    string `column:"ref" gorm:"type:varchar(50)"`
	Ket    string `column:"ket" gorm:"type:varchar(50)"`
	GrpID2 string `column:"grp_id2" gorm:"type:varchar(50)"`
	Ref2   string `column:"ref2" gorm:"type:varchar(50)"`
	Ket2   string `column:"ket2" gorm:"type:varchar(50)"`
}
