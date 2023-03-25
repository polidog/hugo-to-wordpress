package hugo

import (
	"errors"
	"time"
)

var dateFormats = []string{
	"2006-01-02T15:04:05",
	"2006-01-02",
	"2006-01-02T15:04:05Z07:00",
}

func parseDate(dateStr string) (time.Time, error) {
	for _, format := range dateFormats {
		t, err := time.Parse(format, dateStr)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, errors.New("could not parse date with any known format")
}
