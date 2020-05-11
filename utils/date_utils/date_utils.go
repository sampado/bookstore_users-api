package dateutils

import "time"

const (
	apiDateFormat = "2006-01-02T15:04:05Z"
	apiDBFormat   = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	// Placeholder used for the date formatter
	// Mon Jan 2 15:04:05 -0700 MST 2006
	return GetNow().Format(apiDateFormat)
}

func GetNowDBFormat() string {
	return GetNow().Format(apiDBFormat)
}
