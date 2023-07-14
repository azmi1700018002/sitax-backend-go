package m_pajak

type Pajak struct {
	PajakID     int           `column:"id_pajak" gorm:"primary_key"`
	NamaPajak   string        `column:"nama_pajak" gorm:"type:varchar(500)"`
	ParentPajak int           `column:"parent_pajak"`
	StsPajak    int           `column:"sts_pajak"`
	KetPajak    string        `column:"ket_pajak" gorm:"type:varchar(500)"`
	StsParent   int           `column:"sts_parent"`
	FileID      string        `column:"file_id" gorm:"type:char(12)"`
	PajakIDfk   []PajakDetail `gorm:"foreignKey:PajakID"`
}

type PajakDetail struct {
	PajakDetailID int     `column:"pajak_detail_id" gorm:"primary_key"`
	PajakID       int     `column:"pajak_id"`
	Ppn           float64 `column:"ppn"`
	Pasal23       float64 `column:"pasal23"`
	PphFinal      float64 `column:"pph_final"`
	PajakLain     float64 `column:"pajak_lain"`
	Keterangan    string  `column:"keterangan" gorm:"type:varchar(500)"`
}
