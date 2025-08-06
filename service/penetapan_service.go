package service

import (
	"Siberat/dto"
	"Siberat/helper"
	"Siberat/repository"
	"strconv"
	"time"
)

type PenetapanService interface{
	HitungUlang(req dto.HitungUlangRequest) (dto.HitungUlangResponse, error)
	GetPenetapan(IdTrnab string) (*dto.GetTrnabResponse, error)
	HitungPenetapanPABTahunan(req dto.HitungObjekPAB) (*dto.HasilHitungPab, error)
	HitungPabDu(data dto.HitungObjekPAB) (*dto.HasilHitungPab, error)
}

type penetapanService struct{
	repository repository.PenetapanRepository
}

func NewPenetapanService(r repository.PenetapanRepository) *penetapanService {
	return &penetapanService{
		repository: r,
	}
}

func (s *penetapanService) HitungUlang(req dto.HitungUlangRequest) (dto.HitungUlangResponse, error) {
	tglAkhir,_:= time.Parse("2006-01-02", req.TgAkhirPajak)
	tglBaru,_:= time.Parse("2006-01-02", req.TgAkhirPajakBaru)

	diff := helper.GetDateDiffParts(tglAkhir, tglBaru)

	// konversi string to int
	pokok1 := helper.AtoiDefault(req.BeaPabPok1)
	pokok := helper.AtoiDefault(req.BeaPabPok)

	den := (pokok * diff.Months) / 100
	den1 := (pokok1 * diff.Months) / 100

	// data response
	data := dto.HitungUlangResponse{
		BeaBBNABPok: 0,
		BeaBBNABDen: 0,
		BeaPABPok: int64(pokok),
		BeaPABDen: int64(den),
		BeaPABPok1: int64(pokok1),
		BeaPABDen1: int64(den1),
		BeaPABPok2: 0,
		BeaPABDen2: 0,
		BeaPABPok3: 0,
		BeaPABDen3: 0,
		BeaPABPok4: 0,
		BeaPABDen4: 0,
		BeaPABPok5: 0,
		BeaPABDen5: 0,
		BeaBBNABPokD: 0,
		BeaBBNABDenD: 0,
		BeaPABPokD: int64(pokok),
		BeaPABDenD: int64(den),
		BeaPABPok1D: int64(pokok1),
		BeaPABDen1D: int64(den1),
		BeaPABPok2D: 0,
		BeaPABDen2D: 0,
		BeaPABPok3D: 0,
		BeaPABDen3D: 0,
		BeaPABPok4D: 0,
		BeaPABDen4D: 0,
		BeaPABPok5D: 0,
		BeaPABDen5D: 0,
		Ket1:"",
		Ket2:"",
		Ket3:"",
		TgAkhirPajak: req.TgAkhirPajak,
		TgAkhirPajakBaru: req.TgAkhirPajakBaru,
		IsAbleBayar: true,
		SelangHari: diff.Days,
		SelangBulan: diff.Months,
		SelangTahun: diff.Years,
	}

	return data, nil
}

func (s *penetapanService) GetPenetapan(IdTrnab string) (*dto.GetTrnabResponse, error) {
	return s.repository.GetPenetapan(IdTrnab)
}

func (s *penetapanService) HitungPenetapanPABTahunan(req dto.HitungObjekPAB) (*dto.HasilHitungPab, error) {
	var (

	)
	tgDaftar := req.TgDaftar
	tgAkhirPajak:= req.TgAkhirPajak
	
	isAbleBayar := helper.IsAbleBayar(tgDaftar, tgAkhirPajak)
	tgAkhirPajakBaru := helper.GetTanggalAkhirPajakBaru(tgDaftar, tgAkhirPajak, req.ByrKdpn)

	dataHitung := dto.HitungObjekPAB {
		NoAb1: req.NoAb1,
		NoAb2: req.NoAb2,
		NoAb3: req.NoAb3,
		TgAkhirPajak: tgAkhirPajak,
		TgAkhirPajakBaru: tgAkhirPajakBaru,
		NilaiJual: req.NilaiJual,
		Bobot: req.Bobot,
		KdKepemilikan: req.KdKepemilikan,
		KdJenisAb: req.KdJenisAb,
		KdBlockir: req.KdBlockir,
		ByrKdpn: req.ByrKdpn,
		TgDaftar: tgDaftar,
		KdWil: req.KdWil,
	}

	hitungResult,_ := s.HitungPabDu(dataHitung)

	tgdftr,_ := time.Parse("2006-01-02",tgDaftar)
	tgAkhrPjk,_ := time.Parse("2006-01-02",tgAkhirPajak)
	diff := helper.CreateDataSelangHari(tgdftr,tgAkhrPjk)

	// result gabungan
	res := &dto.HasilHitungPab {
		BeaBBNABPok: hitungResult.BeaBBNABPok,
		BeaBBNABDen: hitungResult.BeaBBNABDen,
		BeaPABPok: hitungResult.BeaPABPok,
		BeaPABDen: hitungResult.BeaPABDen,
		BeaPABPok1: hitungResult.BeaPABPok1,
		BeaPABDen1: hitungResult.BeaPABDen1,
		BeaPABPok2: hitungResult.BeaPABPok2,
		BeaPABDen2: hitungResult.BeaPABDen2,
		BeaPABPok3: hitungResult.BeaPABPok3,
		BeaPABDen3: hitungResult.BeaPABDen3,
		BeaPABPok4: hitungResult.BeaPABPok4,
		BeaPABDen4: hitungResult.BeaPABDen4,
		BeaPABPok5: hitungResult.BeaPABPok5,
		BeaPABDen5: hitungResult.BeaPABDen5,
		BeaBBNABPokD: hitungResult.BeaBBNABPokD,
		BeaBBNABDenD: hitungResult.BeaBBNABDenD,
		BeaPABPokD: hitungResult.BeaPABPokD,
		BeaPABDenD: hitungResult.BeaPABDenD,
		BeaPABPok1D: hitungResult.BeaPABPok1D,
		BeaPABDen1D: hitungResult.BeaPABDen1D,
		BeaPABPok2D: hitungResult.BeaPABPok2D,
		BeaPABDen2D: hitungResult.BeaPABDen2D,
		BeaPABPok3D: hitungResult.BeaPABPok3D,
		BeaPABDen3D: hitungResult.BeaPABDen3D,
		BeaPABPok4D: hitungResult.BeaPABPok4D,
		BeaPABDen4D: hitungResult.BeaPABDen4D,
		BeaPABPok5D: hitungResult.BeaPABPok5D,
		BeaPABDen5D: hitungResult.BeaPABDen5D,
		Ket1:"",
		Ket2:"",
		Ket3:"",
		TgAkhirPajak: tgAkhirPajak,
		TgAkhirPajakBaru: tgAkhirPajakBaru,
		IsAbleBayar: isAbleBayar,
		SelangHari: diff.Days,
		SelangBulan: diff.Months,
		SelangTahun: diff.Years,
	}

	return res, nil
}

func (s *penetapanService) HitungPabDu(data dto.HitungObjekPAB) (*dto.HasilHitungPab, error) {
	const (
		maxPajak = 6
		multiplierDenda = 1
		maxBulanDenda = 24
		tarif = 0.2
	)

	nlJual,_ := strconv.ParseFloat(data.NilaiJual, 64)
	koefbobot,_ := strconv.ParseFloat(data.Bobot, 64)

	besaranPokokPab := nlJual*koefbobot*tarif/100

	tanggalDaftar,_ := helper.ParseTanggal(data.TgDaftar)
	tanggalAkhir,_ := helper.ParseTanggal(data.TgAkhirPajak)
	// tanggalAkhirBaru,_ := helper.ParseTanggal(data.TgAkhirPajakBaru)

	pabPok := make([]int64, maxPajak)
	pabDen := make([]int64, maxPajak)
	pabPokD := make([]int64, maxPajak)
	pabDenD := make([]int64, maxPajak)

	currentTgAkhir := tanggalAkhir.AddDate(1,0,0)
	masaPajak := 0
	for masaPajak < maxPajak {
		if !currentTgAkhir.After(tanggalAkhir) {
			break
		}

		currentMasa := currentTgAkhir.AddDate(-1,0,0)
		
		pokok := besaranPokokPab

		var denda float64 = 0

		if tanggalDaftar.After(currentMasa) {
			interval := helper.GetDateDiffParts(currentMasa, tanggalDaftar)
			jumlahBulan := interval.Years*12 + interval.Months
			if interval.Days > 0 {
				jumlahBulan++
			}

			if jumlahBulan > maxBulanDenda {
				jumlahBulan = maxBulanDenda
			}
			persenDenda := float64(jumlahBulan*multiplierDenda)
			denda = persenDenda/100*pokok
		}
		intPokok := int64(pokok)
		intDenda := int64(denda)
		pabPok[masaPajak] = intPokok
		pabDen[masaPajak] = intDenda
		pabPokD[masaPajak] = intPokok
		pabDenD[masaPajak] = intDenda

		currentTgAkhir = currentMasa
		masaPajak++
	}

	// selisih waktu selang
	diff := helper.GetDateDiffParts(tanggalDaftar, tanggalAkhir)
	isAbleBayar := helper.IsAbleBayar(data.TgDaftar, data.TgAkhirPajak)

	response := &dto.HasilHitungPab{
		BeaBBNABPok: 0,
		BeaBBNABDen: 0,
		BeaPABPok: pabPok[0],
		BeaPABDen: pabDen[0],
		BeaPABPok1: pabPok[1],
		BeaPABDen1: pabDen[1],
		BeaPABPok2: pabPok[2],
		BeaPABDen2: pabDen[2],
		BeaPABPok3: pabPok[3],
		BeaPABDen3: pabDen[3],
		BeaPABPok4: pabPok[4],
		BeaPABDen4: pabDen[4],
		BeaPABPok5: pabPok[5],
		BeaPABDen5: pabDen[5],

		BeaBBNABPokD: 0,
		BeaBBNABDenD: 0,
		BeaPABPokD: pabPokD[0],
		BeaPABDenD: pabDenD[0],
		BeaPABPok1D: pabPokD[1],
		BeaPABDen1D: pabDenD[1],
		BeaPABPok2D: pabPokD[2],
		BeaPABDen2D: pabDenD[2],
		BeaPABPok3D: pabPokD[3],
		BeaPABDen3D: pabDenD[3],
		BeaPABPok4D: pabPokD[4],
		BeaPABDen4D: pabDenD[4],
		BeaPABPok5D: pabPokD[5],
		BeaPABDen5D: pabDenD[5],

		Ket1: "",
		Ket2: "",
		Ket3: "",
		TgAkhirPajak: data.TgAkhirPajak,
		TgAkhirPajakBaru: data.TgAkhirPajakBaru,
		IsAbleBayar: isAbleBayar,
		SelangHari: diff.Days,
		SelangBulan: diff.Months,
		SelangTahun: diff.Years,
	}
	return response, nil
}
