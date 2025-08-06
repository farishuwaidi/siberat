package service

import (
	"Siberat/dto"
	"Siberat/helper"
	"Siberat/repository"
	"time"
)

type ObjekService interface{
	GetObjekByNab(NoAb1 string, NoAb2 string, NoAb3 string) (*dto.ObjekPabResponse, error)
	GetObjekSubjek(NoAb1 string, NoAb2 string, NoAb3 string) (dto.GetSubjekObjekPajak, error)
	GetInfoPAB(req dto.ObjekPabRequest) (*dto.InfoPabFlat, error)
}

type objekService struct {
	repository repository.ObjekRepository
	hitungPenetapanTahunan PenetapanService
}

func NewObjeKService(r repository.ObjekRepository, hitungPenetapanTahunan PenetapanService) *objekService {
	return &objekService{
		repository: r,hitungPenetapanTahunan: hitungPenetapanTahunan,
	}
}

func (s *objekService) GetObjekByNab(NoAb1 string, NoAb2 string, NoAb3 string) (*dto.ObjekPabResponse, error) {
	return s.repository.GetObjekByNab(NoAb1, NoAb2, NoAb3)
}

func (s *objekService) GetObjekSubjek(NoAb1 string, NoAb2 string, NoAb3 string) (dto.GetSubjekObjekPajak, error) {
	return s.repository.GetObjekSubjekByNab(NoAb1, NoAb2, NoAb3)
}

func (s *objekService) GetInfoPAB(req dto.ObjekPabRequest) (*dto.InfoPabFlat, error) {
	data, err := s.repository.GetinfoSubjekObjekByNab(req.NoAb1, req.NoAb2, req.NoAb3)
	if err != nil || data == nil{
		return nil, err
	}

	if data.StatusBayarPotensiBaru != "Y" {
		return nil, helper.NewCustomError("gagal memproses info pajak alat berat tahunan, belum membayar potensi baru")
	}
	today := time.Now().Format("2006-01-02")
	tgAkhir := data.TgAkhirPajak
	var tgAkhirFormated time.Time

	tgAkhirFormated,err = time.Parse(time.RFC3339, tgAkhir)
	tgAkhirstr := tgAkhirFormated.Format("2006-01-02")
	
	dataHitungTahunan := dto.HitungObjekPAB {
		NoAb1: req.NoAb1,
		NoAb2: req.NoAb2,
		NoAb3: req.NoAb3,
		TgAkhirPajak: tgAkhirstr,
		NilaiJual: data.NilaiJual,
		Bobot: data.Bobot,
		KdKepemilikan: data.KdKepemilikan,
		KdJenisAb: data.KdJenisAB,
		KdBlockir: data.KdBlockir,
		ByrKdpn: "T",
		TgDaftar: today,
		KdWil: data.KdWil,
	}
	
	hasilHitung, err := s.hitungPenetapanTahunan.HitungPenetapanPABTahunan(dataHitungTahunan)

	if err != nil {
		return nil, err
	}

	result := &dto.InfoPabFlat{
		IdObjekPajak: data.IDObjekPajak,
		NoAb1: data.NoAB1,
		NoAb2: data.NoAB2,
		NoAb3: data.NoAB3,
		KdStatus: data.KdStatus,
		KdMerekAb: data.KdMerekAB,
		ThBuatan: data.ThBuatan,
		KdMohon1: data.KdMohon1,
		KdMohon2: data.KdMohon2,
		KdMohon3: data.KdMohon3,
		KdMohon4: data.KdMohon4,
		KdMohon5: data.KdMohon5,
		KdMohon6: data.KdMohon6,
		NmMerekAb: data.NmMerekAB,
		NmModelAb: data.NmModelAB,
		NoRangka: data.NoRangka,
		NoMesin: data.NoMesin,
		WarnaAb: data.WarnaAB,
		IdBbm: data.IDBBM,
		JumlahCc: data.JumlahCC,
		IdSubjekPajak: data.IDSubjekPajak,
		KdWil: data.KdWil,
		KdWilProses: data.KdWilProses,
		KdJenisAB: data.KdJenisAB,
		KdKepemilikan: data.KdKepemilikan,
		TgKepemilikan: data.TgKepemilikan,
		JamKepemilikan: data.JamKepemilikan,
		KdBlockir: data.KdBlockir,
		TgFaktur: data.TgFaktur,
		TgFiskal: data.TgFiskal,
		TgGntMesin: data.TgGantiMesin,
		TgKwitansi: data.TgKwitansi,
		TgUben: data.TgUben,
		TgBlockir: data.KdBlockir,
		NilaiJual: data.NilaiJual,
		Bobot: data.Bobot,
		PabDasar: data.PabDasar,
		NoKtp: data.NoKTP,
		Npwp: data.NPWP,
		NmPemilik: data.NmPemilik,
		AlPemilik: data.AlPemilik,
		KdWilSubjekPajak: data.KdWilSubjekPajak,
		KdKelurahan: data.KdKelurahan,
		KdKecamatan: data.KdKecamatan,
		Source: data.Source,
		TgUpdate: data.TgUpdate,
		JamUpdate: data.JamUpdate,
		Rt: data.RT,
		Rw: data.RW,
		IdSubjekPajakType: data.IDSubjekPajakType,
		Email: data.Email,
		NoWa: data.NoWA,
		NoHp: data.NoHP,
		KdPos: data.KdPos,
		Nib: data.Nib,
		StatusBayarPotensiBaru: data.StatusBayarPotensiBaru,

		BeaBBNABPok: int(hasilHitung.BeaBBNABPok),
		BeaBBNABDen: int(hasilHitung.BeaBBNABDen),
		BeaPABPok: int(hasilHitung.BeaPABPok),
		BeaPABDen: int(hasilHitung.BeaPABDen),
		BeaPABPok1: int(hasilHitung.BeaPABPok1),
		BeaPABDen1: int(hasilHitung.BeaPABDen1),
		BeaPABPok2: int(hasilHitung.BeaPABPok2),
		BeaPABDen2: int(hasilHitung.BeaPABDen2),
		BeaPABPok3: int(hasilHitung.BeaPABPok3),
		BeaPABDen3: int(hasilHitung.BeaPABDen3),
		BeaPABPok4: int(hasilHitung.BeaPABPok4),
		BeaPABDen4: int(hasilHitung.BeaPABDen4),
		BeaPABPok5: int(hasilHitung.BeaPABPok5),
		BeaPABDen5: int(hasilHitung.BeaPABDen5),

		BeaBBNABPokD: int(hasilHitung.BeaBBNABPokD),
		BeaBBNABDenD: int(hasilHitung.BeaBBNABDenD),
		BeaPABPokD: int(hasilHitung.BeaPABPokD),
		BeaPABDenD: int(hasilHitung.BeaPABDenD),
		BeaPABPok1D: int(hasilHitung.BeaPABPok1D),
		BeaPABDen1D: int(hasilHitung.BeaPABDen1D),
		BeaPABPok2D: int(hasilHitung.BeaPABPok2D),
		BeaPABDen2D: int(hasilHitung.BeaPABDen2D),
		BeaPABPok3D: int(hasilHitung.BeaPABPok3D),
		BeaPABDen3D: int(hasilHitung.BeaPABDen3D),
		BeaPABPok4D: int(hasilHitung.BeaPABPok4D),
		BeaPABDen4D: int(hasilHitung.BeaPABDen4D),
		BeaPABPok5D: int(hasilHitung.BeaPABPok5D),
		BeaPABDen5D: int(hasilHitung.BeaPABDen5D),

		Ket1: hasilHitung.Ket1,
		Ket2: hasilHitung.Ket2,
		Ket3: hasilHitung.Ket3,

		TgAkhirPajak: hasilHitung.TgAkhirPajak,
		TgAkhirPajakBaru: hasilHitung.TgAkhirPajakBaru,
		IsAbleBayar: hasilHitung.IsAbleBayar,
		SelangHari: hasilHitung.SelangHari,
		SelangBulan: hasilHitung.SelangBulan,
		SelangTahun: hasilHitung.SelangTahun,
	}

	return result, nil
}