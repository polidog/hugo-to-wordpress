package hugo

import (
	"errors"
	"time"
)

func parseDate(dateStr string) (time.Time, error) {
	var dateFormats = []string{
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05",
		"2006-01-02",
	}

	var parsedDate time.Time
	var err error
	for _, format := range dateFormats {
		parsedDate, err = time.Parse(format, dateStr)
		if err == nil {
			return parsedDate, nil
		}
	}

	return time.Time{}, errors.New("failed to parse date")
}
