package gen

import (
	"fmt"
	"strings"
	"time"

	"github.com/teris-io/shortid"
)

func GenerateShipmentCode() (string, error) {
	currentDate := time.Now()

	month := strings.ToUpper(currentDate.Format("Jan")) // Tháng viết tắt (Uppercase)
	day := currentDate.Format("02")                     // Ngày
	year := currentDate.Format("06")                    // Năm 2 chữ số

	prefix := fmt.Sprintf("%s%s%s", month, day, year)

	code, err := shortid.Generate()
	if err != nil {
		return "", err
	}

	shortCode := strings.ToUpper(code[:14-len(prefix)])

	return prefix + shortCode, nil
}
