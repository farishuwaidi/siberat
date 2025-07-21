package helper

import (
	"strconv"
	"time"
)

func GetDateDiffParts(t1, t2 time.Time) (tahun int, bulan int, hari int) {
	tahun = int(t2.Year()-t1.Year())
	bulan = int(t2.Month()-t1.Month())
	hari = t2.Day()-t1.Day()

	if hari < 0 {
		lastMonth := t2.AddDate(0, -1, 0)
		hari += daysIn(lastMonth)
		bulan--
	}

	if bulan < 0 {
		bulan += 12
		tahun--
	}
	return
}

func daysIn(t time.Time) int {
	year, month := t.Year(), t.Month()
	loc := t.Location()
	firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, loc)
	firstOfNextMonth := firstOfMonth.AddDate(0, 1, 0)
	return int(firstOfNextMonth.Sub(firstOfMonth).Hours()/24)
}

func AtoiDefault(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}

func IsAbleBayar(tgDaftar, TgAkhirPajak time.Time) bool {
	if tgDaftar.After(TgAkhirPajak) {
		return true
	}
	diff := TgAkhirPajak.Sub(tgDaftar)
	if int(diff.Hours()/24) < 180 {
		return true
	}
	return false
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
