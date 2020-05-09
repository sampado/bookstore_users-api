package dateutils

import "time"

const (
	apiDateFormat = "2006-01-02:04:05Z"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	// Placeholder used for the date formatter
	// Mon Jan 2 15:04:05 -0700 MST 2006
	return GetNow().Format(apiDateFormat)
}
