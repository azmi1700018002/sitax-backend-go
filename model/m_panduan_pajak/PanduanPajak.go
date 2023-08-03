package m_panduan_pajak

type PanduanPajak struct {
	PanduanPajakID     int    `column:"id_panduan_pajak" gorm:"primarykey"`
	NamaPanduanPajak   string `column:"nama_panduan_pajak" gorm:"type=varchar(500)"`
	ParentPanduanPajak string `column:"parent_panduan_pajak" gorm:"type=char(4)"`
	StsPanduanPajak    string `column:"sts_panduan_pajak" gorm:"type=char(1)"`
	StsParent          string `column:"sts_parent" gorm:"type=char(1)"`
	FileID             string `column:"file_id" gorm:"type:char(12)"`
	
}
