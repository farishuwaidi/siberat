package entity

type RoleUser struct {
	IdRole                int    `gorm:"column:id_role;primarykey"`
	NameRole              string `gorm:"column:nm_role"`
	StsDaftar             int    `gorm:"column:sts_daftar"`
	StsTambahPotensi      int    `gorm:"column:sts_tambah_potensi"`
	StsInfoPab            int    `gorm:"column:sts_info_pab"`
	StsKirimUlangSKKP     int    `gorm:"column:sts_kirim_ulg_skkp"`
	StsCetakNAB           int    `gorm:"column:sts_cetak_nab"`
	StsCetakUlangNAB      int    `gorm:"column:sts_cetak_ulg_nab"`
	StsInfoSubjekPab      int    `gorm:"column:sts_info_subjek_pab"`
	StsDaftarUlangPotBaru int    `gorm:"column:sts_daftar_ulg_pot_baru"`
	StsUserMngmt          int    `gorm:"column:sts_user_mngmt"`
	StsNabMngmt           int    `gorm:"column:sts_njab_mngmt"`
	StsInfoNjab           int    `gorm:"column:sts_info_njab"`
	StsInfoUser           int    `gorm:"column:sts_info_user"`
	StsTable              int    `gorm:"column:sts_table"`
	StsTableMngmt         int    `gorm:"column:sts_table_mngmt"`
	StsVerif              int    `gorm:"column:sts_verif"`
}
