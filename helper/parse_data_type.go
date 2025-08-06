package helper

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
)

// StructToMap mengubah struct menjadi map[string]interface{}
func StructToMap(data interface{}) map[string]interface{} {
	var result map[string]interface{}
	temp, _ := json.Marshal(data)
	json.Unmarshal(temp, &result)
	return result
}

func StringPtr(s string) *string {
	return &s
}

func FormatObjekResponse(input map[string]interface{}) map[string]interface{} {
	formatted := make(map[string]interface{})

	for key, value := range input {
		switch v := value.(type) {
		case time.Time:
			formatted[key] = v.Format("2006-01-02") // format tanggal
		case float64:
			if key == "bobot" {
				formatted[key] = fmt.Sprintf("%.3f", v)
			} else {
				formatted[key] = int64(v)
			}
		case int:
			formatted[key] = fmt.Sprintf("%d", v)
		case int64:
			formatted[key] = fmt.Sprintf("%d", v)
		case nil:
			formatted[key] = nil
		default:
			// untuk data lain, pastikan string
			rv := reflect.ValueOf(v)
			if rv.Kind() == reflect.String {
				formatted[key] = v
			} else {
				formatted[key] = v
			}
		}
	}

	// Hardcoded tambahan formatting opsional
	if tgBaru, ok := formatted["tg_akhir_pajak_baru"]; !ok || tgBaru == "" {
		formatted["tg_akhir_pajak_baru"] = time.Now().AddDate(1, 0, 0).Format("2006-01-02")
	}

	return formatted
}

func InterfaceToString(val interface{}) string {
	if val == nil {
		return ""
	}
	return strings.TrimSpace(fmt.Sprintf("%v", val))
}

func CeilByPecahan(value float64) int64 {
	pecahan := 1000.0
	return int64(pecahan * float64(int((value+pecahan-1)/pecahan)))
}