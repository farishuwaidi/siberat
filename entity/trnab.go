package entity

import "time"

type Trnab struct {
	IdTrnab            string     `gorm:"column:id_trnab"`
	KdMohon1           string     `gorm:"column:kd_mohon1"`
	KdMohon2           string     `gorm:"column:kd_mohon2"`
	KdMohon3           string     `gorm:"column:kd_mohon3"`
	KdMohon4           string     `gorm:"column:kd_mohon4"`
	KdMohon5           string     `gorm:"column:kd_mohon5"`
	KdMohon6           string     `gorm:"column:kd_mohon6"`
	KdStatus           string     `gorm:"column:kd_status"`
	KdWilProses        string     `gorm:"column:kd_wil_proses"`
	TgDaftar           *time.Time `gorm:"column:tg_daftar"`
	JamDaftar          string     `gorm:"column:jam_daftar"`
	IdUserDaftar       string     `gorm:"column:id_user_daftar"`
	NmUserDaftar       string     `gorm:"column:nm_user_daftar"`
	TgTetap            *time.Time `gorm:"column:tg_tetap"`
	JamTetap           string     `gorm:"column:jam_tetap"`
	IdUserTetap        string     `gorm:"column:id_user_tetap"`
	NmUserTetap        string     `gorm:"column:nm_user_tetap"`
	TgBayar            *time.Time `gorm:"column:tg_bayar"`
	JamBayar           string     `gorm:"column:jam_bayar"`
	IdUserBayar        string     `gorm:"column:id_user_bayar"`
	NmUserBayar        string     `gorm:"column:nm_user_bayar"`
	KdChannelBayar     string     `gorm:"column:kd_channel_bayar"`

	BeaBbnabPok        int        `gorm:"column:bea_bbnab_pok"`
	BeaBbnabDen        int        `gorm:"column:bea_bbnab_den"`
	BeaPabPok          int        `gorm:"column:bea_pab_pok"`
	BeaPabPok1         int        `gorm:"column:bea_pab_pok1"`
	BeaPabPok2         int        `gorm:"column:bea_pab_pok2"`
	BeaPabPok3         int        `gorm:"column:bea_pab_pok3"`
	BeaPabPok4         int        `gorm:"column:bea_pab_pok4"`
	BeaPabPok5         int        `gorm:"column:bea_pab_pok5"`
	BeaPabDen          int        `gorm:"column:bea_pab_den"`
	BeaPabDen1         int        `gorm:"column:bea_pab_den1"`
	BeaPabDen2         int        `gorm:"column:bea_pab_den2"`
	BeaPabDen3         int        `gorm:"column:bea_pab_den3"`
	BeaPabDen4         int        `gorm:"column:bea_pab_den4"`
	BeaPabDen5         int        `gorm:"column:bea_pab_den5"`

	BeaBbnabPokD       int        `gorm:"column:bea_bbnab_pok_d"`
	BeaBbnabDenD       int        `gorm:"column:bea_bbnab_den_d"`
	BeaPabPokD         int        `gorm:"column:bea_pab_pok_d"`
	BeaPabPok1D        int        `gorm:"column:bea_pab_pok1_d"`
	BeaPabPok2D        int        `gorm:"column:bea_pab_pok2_d"`
	BeaPabPok3D        int        `gorm:"column:bea_pab_pok3_d"`
	BeaPabPok4D        int        `gorm:"column:bea_pab_pok4_d"`
	BeaPabPok5D        int        `gorm:"column:bea_pab_pok5_d"`
	BeaPabDenD         int        `gorm:"column:bea_pab_den_d"`
	BeaPabDen1D        int        `gorm:"column:bea_pab_den1_d"`
	BeaPabDen2D        int        `gorm:"column:bea_pab_den2_d"`
	BeaPabDen3D        int        `gorm:"column:bea_pab_den3_d"`
	BeaPabDen4D        int        `gorm:"column:bea_pab_den4_d"`
	BeaPabDen5D        int        `gorm:"column:bea_pab_den5_d"`

	TgAkhirPajak       *time.Time `gorm:"column:tg_akhir_pajak"`
	TgAkhirPajakBaru   *time.Time `gorm:"column:tg_akhir_pajak_baru"`
	TgAkhirPajakLama   *time.Time `gorm:"column:tg_akhir_pajak_lama"`

	TgFaktur           *time.Time `gorm:"column:tg_faktur"`
	TgFiskal           *time.Time `gorm:"column:tg_fiskal"`
	TgGntMesin         *time.Time `gorm:"column:tg_gntmesin"`
	TgKwitansi         *time.Time `gorm:"column:tg_kwitansi"`
	TgUben             *time.Time `gorm:"column:tg_uben"`
	TgBlockir          *time.Time `gorm:"column:tg_blockir"`

	NoAb1              string     `gorm:"column:no_ab1"`
	NoAb2              string     `gorm:"column:no_ab2"`
	NoAb3              string     `gorm:"column:no_ab3"`

	IdObjekPajak       string     `gorm:"column:id_objek_pajak"`
	KdMerekAb          string     `gorm:"column:kd_merek_ab"`
	NmMerekAb          string     `gorm:"column:nm_merek_ab"`
	NmModelAb          string     `gorm:"column:nm_model_ab"`
	KdJenisAb          string     `gorm:"column:kd_jenis_ab"`
	ThBuatan           string     `gorm:"column:th_buatan"`
	NoRangka           string     `gorm:"column:no_rangka"`
	NoMesin            string     `gorm:"column:no_mesin"`
	WarnaAb            string     `gorm:"column:warna_ab"`
	IdBbm              string     `gorm:"column:id_bbm"`
	JumlahCc           int        `gorm:"column:jumlah_cc"`

	KdWil              string     `gorm:"column:kd_wil"`
	KdKepemilikan      string     `gorm:"column:kd_kepemilikan"`
	KdBlockir          string     `gorm:"column:kd_blockir"`
	TgKepemilikan      *time.Time `gorm:"column:tg_kepemilikan"`
	JamKepemilikan     string     `gorm:"column:jam_kepemilikan"`

	IdSubjekPajak      string     `gorm:"column:id_subjek_pajak"`
	NoKtp              string     `gorm:"column:no_ktp"`
	Npwp               string     `gorm:"column:npwp"`
	Nib                string     `gorm:"column:nib"`
	NmPemilik          string     `gorm:"column:nm_pemilik"`
	AlPemilik          string     `gorm:"column:al_pemilik"`
	KdWilSubjekPajak   string     `gorm:"column:kd_wil_subjek_pajak"`
	KdKelurahan        string     `gorm:"column:kd_kelurahan"`
	KdKecamatan        string     `gorm:"column:kd_kecamatan"`
	Rt                 string     `gorm:"column:rt"`
	Rw                 string     `gorm:"column:rw"`
	IdSubjekPajakType  string     `gorm:"column:id_subjek_pajak_type"`
	Email              string     `gorm:"column:email"`
	NoWa               string     `gorm:"column:no_wa"`
	NoHp               string     `gorm:"column:no_hp"`
	KdPos              string     `gorm:"column:kd_pos"`

	IsAbleBayar        bool       `gorm:"column:is_able_bayar"`
	NilaiJual          int        `gorm:"column:nilai_jual"`
	Bobot              string     `gorm:"column:bobot"`
	BayarKdpn          string     `gorm:"column:bayar_kdpn"`
	KdCetakTnab        string     `gorm:"column:kd_cetak_tnab"`

	// Koreksi
	TgKoreksi          *time.Time `gorm:"column:tg_koreksi"`
	JamKoreksi         string     `gorm:"column:jam_koreksi"`
	IdUserKoreksi      string     `gorm:"column:id_user_koreksi"`
	NmUserKoreksi      string     `gorm:"column:nm_user_koreksi"`

	// Cetak
	TgCetakTnab        *time.Time `gorm:"column:tg_cetak_tnab"`
	JamCetakTnab       string     `gorm:"column:jam_cetak_tnab"`
	IdUserCetakTnab    string     `gorm:"column:id_user_cetak_tnab"`
	NmUserCetakTnab    string     `gorm:"column:nm_user_cetak_tnab"`

	// Pembayaran
	KodeBayar          string     `gorm:"column:kode_bayar"`
	KodeBayarExpired   *time.Time `gorm:"column:kode_bayar_expired"`
	TagihanPokokTot    int        `gorm:"column:tagihan_pokok_tot"`
	TagihanDendaTot    int        `gorm:"column:tagihan_denda_tot"`
	TagihanNoAb        string     `gorm:"column:tagihan_no_ab"`
	TagihanTot         int        `gorm:"column:tagihan_tot"`

	UnitKode           string     `gorm:"column:unit_kode"`
	IsPabEditable      bool       `gorm:"column:is_pab_editable"`

	// Timestamps
	CreatedAt          *time.Time `gorm:"column:created_at"`
	UpdatedAt          *time.Time `gorm:"column:updated_at"`
}