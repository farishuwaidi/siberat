package helper

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

const dateFormat = "2006-01-02"
type DateDiff struct {
	Years  int
	Months int
	Days   int
	TotalDays int
}

func GetDateDiffParts(t1, t2 time.Time) DateDiff {
	if t1.After(t2) {
		t1, t2 = t2, t1
	}
	years := t2.Year() - t1.Year()
	months := int(t2.Month()) - int(t1.Month())
	days := t2.Day() - t1.Day()

	if days < 0 {
		months--
		prevMonth := t2.AddDate(0, 0, -t2.Day())
		days += prevMonth.Day()
	}

	if months < 0 {
		years--
		months += 12
	}
	totalDays := int(t2.Sub(t1).Hours()/24)+1

	return DateDiff{
		Years: years,
		Months: months,
		Days: days,
		TotalDays: totalDays,
	}
}

func daysInMonth(year int, month time.Month) int {
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

func CreateDataSelangHari(tgAwal, tgAkhir time.Time) DateDiff {
	diff := GetDateDiffParts(tgAkhir, tgAwal)

	return DateDiff {
		Years: diff.Years,
		Months: diff.Months,
		Days: diff.Days,
	}
}

func AtoiDefault(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}

func IsAbleBayar(tgDaftarStr, tgAkhirStr string) bool {
	tgDaftar, err1 := time.Parse("2006-01-02", tgDaftarStr)
	tgAkhir, err2 := time.Parse("2006-01-02", tgAkhirStr)
	if err1 != nil || err2 != nil {
		return false
	}

	if tgDaftar.After(tgAkhir) {
		return true
	}

	selisihHari := GetDateDiffParts(tgDaftar, tgAkhir)
	totalHari := (selisihHari.Years*365) + (selisihHari.Months*30) + selisihHari.Days
	return totalHari < 180
}


func FormatDate(s string) string {
	if s == "" {
		return ""
	}
	parsed, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return s
	}
	return parsed.Format("2006-01-02")
}

func ParseTanggal(str string) (time.Time, error) {
	layouts := []string{
		"2006-01-02",
		"2006-01-02 15:04:05",
		"2006-01-02 15:04:05 -0700 MST",
	}
	for _, layout := range layouts {
		if t, err := time.Parse(layout, str); err == nil {
			return t, nil
		}
	}
	return time.Time{}, errors.New("format tanggal tidak dikenali")
}

func GetTanggalAkhirPajakBaru(tgDaftarstr, tgAkhirstr, byrkdpn string) string {
	tgDaftar,_ := time.Parse(dateFormat, tgDaftarstr)
	tgAkhir,_ := time.Parse(dateFormat, tgAkhirstr)
	tgAkhirBaru := tgAkhir

	// jika tgakhir >= tgDaftar
	if !tgAkhir.Before(tgDaftar) {
		tgAkhirBaru = tgAkhir.AddDate(1,0,0)
	} else {
		// ambil tahun berjalan dari akhir pajak
		thSekarang := time.Now().Year()
		monthDay := tgAkhir.Format("-01-02")
		tgAkhirTahunBerjalan,_ := time.Parse(dateFormat, fmt.Sprintf("%d%s", thSekarang, monthDay))

		if tgAkhirTahunBerjalan.Before(tgDaftar) {
			tgAkhirBaru = tgAkhirTahunBerjalan.AddDate(1,0,0)
		} else {
			tgAkhirBaru = tgAkhirTahunBerjalan
		}
		if tgAkhirBaru.Before(tgDaftar) {
			tgAkhirBaru = tgAkhirBaru.AddDate(1,0,0)
		}
	}

	diff := GetDateDiffParts(tgAkhirBaru, tgDaftar)
	addBaru := false
	if diff.TotalDays < 30 {
		addBaru = true
	} else {
		if diff.TotalDays < 180 && byrkdpn == "Y" {
			addBaru = true
		}
	}

	if addBaru {
		tgAkhirBaru = tgAkhirBaru.AddDate(1,0,0)
	}

	return tgAkhirBaru.Format(dateFormat)
}
