package m_user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Username       string         `column:"username" gorm:"type:varchar(50);primary_key"`
	Password       string         `column:"password" gorm:"type:varchar(200)"`
	NamaLengkap    string         `column:"nama_lengkap" gorm:"type:varchar(100)"`
	Email          string         `column:"email" gorm:"type:varchar(50)"`
	HostIP         string         `column:"host_ip" gorm:"type:varchar(50)"`
	StsUser        string         `column:"sts_user" gorm:"type:char(1)"`
	GroupID        string         `column:"group_id" gorm:"type:int"`
	KdKantor       *string        `column:"kd_kantor" gorm:"type:char(3)"`
	ProfilePicture string         `column:"profile_picture"`
	CreatedAt      time.Time      `column:"created_at"`
	UpdatedAt      time.Time      `column:"updated_at"`
	DeletedAt      gorm.DeletedAt `column:"deleted_at"`
	LastLogin      *time.Time     `column:"last_login"`
}
