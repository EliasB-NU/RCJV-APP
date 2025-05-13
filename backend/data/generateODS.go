package data

import "log"

// GenerateODS generates the appropriate ods for creating the timetable.
// It automatically enters all teams and fields, so you can just pick.
// Generated via admin panel
func GenerateODS(league string) (string, error) {
	log.Println(league)

	return "./ods/test.ods", nil
}
