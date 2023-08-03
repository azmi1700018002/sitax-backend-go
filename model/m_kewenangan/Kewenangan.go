package m_kewenangan

// import "sitax/model/m_user"

type Kewenangan struct {
	GroupID   string `column:"group_id" gorm:"type:varchar(4)"`
	MenuID    string `column:"menu_id" gorm:"type:varchar(4)"`
	IsCreated string `column:"is_created" gorm:"type:varchar(1)"`
	IsUpdated string `column:"is_updated" gorm:"type:varchar(1)"`
	IsDeleted string `column:"is_deleted" gorm:"type:varchar(1)"`

	// GroupIDfk []m_user.User `gorm:"foreignKey:GroupID"`
}
