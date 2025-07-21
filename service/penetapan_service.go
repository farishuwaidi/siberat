package service

import (
	"Siberat/dto"
	"Siberat/helper"
	"Siberat/repository"
	"time"
)

type PenetapanService interface{
	HitungUlang(req dto.HitungUlangRequest) (dto.HitungUlangResponse, error)
	GetPenetapan(IdTrnab string) (*dto.GetTrnabResponse, error)
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

	selangTahun, selangBulan, selangHari := helper.GetDateDiffParts(tglAkhir, tglBaru)

	// konversi string to int
	pokok1 := helper.AtoiDefault(req.BeaPabPok1)
	pokok := helper.AtoiDefault(req.BeaPabPok)

	den := (pokok * selangBulan) / 100
	den1 := (pokok1 * selangBulan) / 100

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
		SelangHari: selangHari,
		SelangBulan: selangBulan,
		SelangTahun: selangTahun,
	}

	return data, nil
}

func (s *penetapanService) GetPenetapan(IdTrnab string) (*dto.GetTrnabResponse, error) {
	return s.repository.GetPenetapan(IdTrnab)
}