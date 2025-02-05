package convert

import (
	"time"

	"github.com/erickxeno/clib/errors"
)

// ToTime convert interface to time.Time
func ToTime(i interface{}) time.Time {
	v, _ := ToTimeE(i)
	return v
}

// ToDuration convert interface to time.Duration
func ToDuration(i interface{}) time.Duration {
	v, _ := ToDurationE(i)
	return v
}

// StringToDateE attempts to parse a string into a time.Time type using a
// predefined list of formats.  If no suitable format is found, an error is
// returned.
func StringToDateE(s string) (time.Time, error) {
	return parseDateWith(s, []string{
		time.RFC3339,
		"2006-01-02T15:04:05", // iso8601 without timezone
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822Z,
		time.RFC822,
		time.RFC850,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		"2006-01-02 15:04:05.999999999 -0700 MST", // Time.String()
		"2006-01-02",
		"02 Jan 2006",
		"2006-01-02T15:04:05-0700", // RFC3339 without timezone hh:mm colon
		"2006-01-02 15:04:05 -07:00",
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05Z07:00", // RFC3339 without T
		"2006-01-02 15:04:05Z0700",  // RFC3339 without T or timezone hh:mm colon
		"2006-01-02 15:04:05",
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	})
}

func parseDateWith(s string, dates []string) (d time.Time, e error) {
	for _, dateType := range dates {
		if d, e = time.Parse(dateType, s); e == nil {
			return
		}
	}
	return d, errors.Errorf("unable to parse date: %s", s)
}

// SecToTime convert time second to time.Time
func SecToTime(sec int64) time.Time {
	return time.Unix(sec, 0)
}

// MsToTime convert time millisecond to time.Time
func MsToTime(ms int64) time.Time {
	return time.Unix(ms/1e3, 0)
}

// UsToTime convert time microsecond to time.Time
func UsToTime(us int64) time.Time {
	return time.Unix(us/1e6, us%1e6*1e3)
}

// NsToTime convert time nanosecond to time.Time
func NsToTime(nsec int64) time.Time {
	return time.Unix(nsec/1e9, nsec%1e9)
}

// TimeToSec convert time.Time to second
func TimeToSec(t time.Time) int64 {
	if t.IsZero() {
		return 0
	} else {
		return t.UnixNano() / 1e9
	}
}

// TimeToMs convert time.Time to millisecond
func TimeToMs(t time.Time) int64 {
	if t.IsZero() {
		return 0
	} else {
		return t.UnixNano() / 1e6
	}
}

// TimeToUs convert time.Time to microsecond
func TimeToUs(t time.Time) int64 {
	if t.IsZero() {
		return 0
	} else {
		return t.UnixNano() / 1e3
	}
}

// TimeToNs convert time.Time to nanosecond
func TimeToNs(t time.Time) int64 {
	if t.IsZero() {
		return 0
	} else {
		return t.UnixNano()
	}
}
