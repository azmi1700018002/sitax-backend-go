package m_menu

import "sitax/model/m_kewenangan"

type Menu struct {
	MenuID        string `column:"menu_id"  gorm:"type:char(4);primarykey"`
	MenuNama      string `column:"menu_nama" gorm:"type:varchar(100)"`
	MenuLink      string `column:"menu_link" gorm:"type:varchar(100)"`
	MenuDeskripsi string `column:"menu_deskripsi" gorm:"type:varchar(200)"`
	MenuStatus    string `column:"menu_status" gorm:"type:varchar(1)"`
	MenuIcon      string `column:"menu_icon" gorm:"type:varchar(50)"`
	MenuKategori  string `column:"menu_kategori" gorm:"type:varchar(50)"`
	ParentID      string `column:"parent_id" gorm:"type:char(4);default:null"`
	ParentSts     string `column:"parent_sts" gorm:"type:char(1);default:null"`
	NoUrut        int    `column:"no_urut"`

	MenuIDfk []m_kewenangan.Kewenangan `gorm:"foreignKey:MenuID"`
}
