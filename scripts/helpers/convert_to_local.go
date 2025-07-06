// package convert_to_local

// import (
// 	"fmt"
// 	"regexp"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// // ConvertToLocalTimezone converts a UTC time to the provided timezone (IANA or offset like 'UTC+5').
// func ConvertToLocalTimezone(scheduledDate time.Time, timezone string) (time.Time, error) {
// 	// Try to load as IANA location
// 	loc, err := time.LoadLocation(timezone)
// 	if err == nil {
// 		return scheduledDate.In(loc), nil
// 	}

// 	// Try to parse as offset (e.g., "UTC+5", "GMT-3")
// 	re := regexp.MustCompile(`^(UTC|GMT)([+-]\d{1,2})$`)
// 	matches := re.FindStringSubmatch(strings.ToUpper(timezone))
// 	if len(matches) == 3 {
// 		offsetHours, err := strconv.Atoi(matches[2])
// 		if err != nil {
// 			return time.Time{}, fmt.Errorf("invalid offset in timezone: %v", err)
// 		}
// 		loc = time.FixedZone(timezone, offsetHours*3600)
// 		return scheduledDate.In(loc), nil
// 	}

// 	return time.Time{}, fmt.Errorf("invalid timezone: %s", timezone)
// }
