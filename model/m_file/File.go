package m_file

import (
	"sitax/model/m_pajak"
	"time"
)

type File struct {
	FileID    string    `column:"file_id" gorm:"type:char(12);primarykey"`
	FileJudul string    `column:"file_judul" gorm:"type:varchar(100)"`
	FilePath  string    `column:"file_path" gorm:"type:varchar"`
	FileDate  time.Time `column:"file_date"`
	FileJenis string    `column:"file_jenis" gorm:"type:char(1)"`

	FileIDfk []m_pajak.Pajak `gorm:"foreignKey:FileID"`
}
