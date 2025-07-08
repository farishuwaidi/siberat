package entity

type RoleUser struct {
	IdRole              string `gorm:"column:id_role;primarykey" json:"id_role"`
	NmRole              string `gorm:"column:nm_role" json:"nm_role"`
	StsDaftar           int    `gorm:"column:sts_daftar" json:"sts_daftar"`
	StsTambahPotensi    int    `gorm:"column:sts_tambah_potensi" json:"sts_tambah_potensi"`
	StsInfoPab          int    `gorm:"column:sts_info_pab" json:"sts_info_pab"`
	StsKirimUlgSkkp     int    `gorm:"column:sts_kirim_ulg_skkp" json:"sts_kirim_ulg_skkp"`
	StsCetakNab         int    `gorm:"column:sts_cetak_nab" json:"sts_cetak_nab"`
	StsCetakUlgNab      int    `gorm:"column:sts_cetak_ulg_nab" json:"sts_cetak_ulg_nab"`
	StsInfoSubjekPab    int    `gorm:"column:sts_info_subjek_pab" json:"sts_info_subjek_pab"`
	StsDaftarUlgPotBaru int    `gorm:"column:sts_daftar_ulg_pot_baru" json:"sts_daftar_ulg_pot_baru"`
	StsUserMngmt        int    `gorm:"column:sts_user_mngmt" json:"sts_user_mngmt"`
	StsNjabMngmt        int    `gorm:"column:sts_njab_mngmt" json:"sts_njab_mngmt"`
	StsInfoNjab         int    `gorm:"column:sts_info_njab" json:"sts_info_njab"`
	StsInfoUser         int    `gorm:"column:sts_info_user" json:"sts_info_user"`
	StsTable            int    `gorm:"column:sts_table" json:"sts_table"`
	StsTableMngmt       int    `gorm:"column:sts_table_mngmt" json:"sts_table_mngmt"`
	StsVerif            int    `gorm:"column:sts_verif" json:"sts_verif"`
}
