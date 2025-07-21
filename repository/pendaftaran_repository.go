package repository

import (
	"Siberat/dto"
	"time"

	"gorm.io/gorm"
)

type PendaftaranRepository interface{
	CekTransaksi(NoAb1 string, NoAb2 string, NoAb3 string) (*dto.CekTransaksiResponse, error)
}

type pendaftaranRepository struct {
	db *gorm.DB
}

func NewPendaftaranRepository(db *gorm.DB) *pendaftaranRepository {
	return &pendaftaranRepository{db}
}

func (r *pendaftaranRepository) CekTransaksi(NoAb1 string, NoAb2 string,NoAb3 string) (*dto.CekTransaksiResponse, error) {
	type Trnab struct {
		IdTrnab string
		KdStatus string
		KdWilProses string
		TgDaftar *time.Time
		JamDaftar string
		TgTetap *time.Time
		JamTetap string
		TgBayar *time.Time
		JamBayar string
	}
	var trnab Trnab
	err := r.db.
		Table("t_trnab").
		Where("no_ab1 = ? AND no_ab2 = ? AND no_ab3 = ?", NoAb1, NoAb2, NoAb3).
		First(&trnab).Error
	
	if err == gorm.ErrRecordNotFound{
		return &dto.CekTransaksiResponse{
			IsExist: false,
		}, nil
	} else if err != nil {
		return nil, err
	}

	formatDate := func(t *time.Time) string {
		if t == nil {
			return ""
		}
		return t.Format("2006-01-02")
	}

	return &dto.CekTransaksiResponse{
		IsExist: true,
		IdTrnab: trnab.IdTrnab,
		KdStatus: trnab.KdStatus,
		KdWilProses: trnab.KdWilProses,
		TgDaftar: formatDate(trnab.TgDaftar),
		JamDaftar: trnab.JamDaftar,
		TgTetap: formatDate(trnab.TgTetap),
		JamTetap: trnab.JamTetap,
		TgBayar: formatDate(trnab.TgBayar),
		JamBayar: trnab.JamBayar,
	}, nil
}