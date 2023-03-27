package hugo

import (
	"errors"
	"time"
)

const (
	layout1 = "2006-01-02T15:04:05"
	layout2 = "2006-01-02"
	layout3 = "2006-01-02T15:04:05Z07:00"
)

func Parse(dateStr string) (time.Time, error) {
	layouts := []string{layout1, layout2, layout3}

	for _, layout := range layouts {
		parsedDate, err := time.Parse(layout, dateStr)
		if err == nil {
			return parsedDate, nil
		}
	}

	return time.Time{}, errors.New("failed to parse date")
}
