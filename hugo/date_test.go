package hugo

import (
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected time.Time
		err      bool
	}{
		{
			name:     "RFC3339",
			input:    "2006-01-02T15:04:05Z",
			expected: time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
			err:      false,
		},
		{
			name:     "ShortDate",
			input:    "2010-03-28",
			expected: time.Date(2010, 3, 28, 0, 0, 0, 0, time.UTC),
			err:      false,
		},
		{
			name:     "ShortDate 2012-11-15",
			input:    "2012-11-15",
			expected: time.Date(2012, 11, 15, 0, 0, 0, 0, time.UTC),
			err:      false,
		},
		{
			name:     "RFC3339 with timezone",
			input:    "2023-03-25T02:40:05+09:00",
			expected: time.Date(2023, 3, 25, 2, 40, 5, 0, time.FixedZone("UTC+9", 9*60*60)),
			err:      false,
		},
		{
			name:  "InvalidDate",
			input: "invalid date",
			err:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Parse(tc.input)

			if tc.err {
				if err == nil {
					t.Error("expected error, but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if !result.Equal(tc.expected) {
					t.Errorf("expected %v, but got %v", tc.expected, result)
				}
			}
		})
	}
}
