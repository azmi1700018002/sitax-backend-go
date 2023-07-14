package m_group

import "sitax/model/m_user"

type Group struct {
	GroupID        string `column:"group_id" gorm:"type:char(4);primarykey"`
	GroupNama      string `column:"group_nama" gorm:"type:varchar(100)"`
	GroupDeskripsi string `column:"group_deskripsi" gorm:"type:varchar(100)"`

	GroupIDfk []m_user.User `gorm:"foreignKey:GroupID"`
}
