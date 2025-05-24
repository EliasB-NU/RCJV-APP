package data

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"log"
	"rcjv-app/backend/database"
	"strconv"
	"sync"
	"time"
)

// StoreMatches parses the csv file and stores the matches in the database.
// Sends a message to all connected clients to reload their data
// A example of the timetable can be found in root of the directory
func StoreMatches(tmpPath string, league string, db *gorm.DB) {
	var (
		newMatches []database.Match

		teams            []database.Team
		teamsInstitution map[string]database.Institution

		err error
	)
	teamsInstitution = make(map[string]database.Institution)

	// Get all teams and Institution
	err = db.Preload("Institution").Find(&teams).Error
	if err != nil {
		log.Printf("Error fetching teams & preloading institution: %v\n", err)
	}
	for _, team := range teams {
		teamsInstitution[team.Name] = team.Institution
	}

	// Open xlsx
	file, err := excelize.OpenFile(tmpPath)
	if err != nil {
		log.Printf("Error opening file: %v\n", err)
	}

	// Parse each sheet
	sheets := file.GetSheetList()
	var wg sync.WaitGroup

	for _, sheet := range sheets {
		wg.Add(1)
		go func(sheet string) {
			parsedMatches := parseSheet(league, sheet, teamsInstitution, teams, file)
			newMatches = append(newMatches, parsedMatches...)
			defer wg.Done()
		}(sheet)
		wg.Wait()
	}

	// Delete all current matches and load the new ones
	err = db.Where("league = ?", league).Delete(&database.Match{}).Error
	if err != nil {
		log.Printf("Error deleting matches: %v\n", err)
	}

	if len(newMatches) == 0 {
		log.Printf("No matches found for %s\n", league)
		return
	}
	err = db.Create(&newMatches).Error
	if err != nil {
		log.Printf("Error creating matches: %v\n", err)
	}

	log.Println("Created matches")
}

func parseSheet(league string, sheet string, teamsInstitution map[string]database.Institution, teams []database.Team, file *excelize.File) []database.Match {
	var matches []database.Match
	// Get all rows
	rows, err := file.GetRows(sheet)
	if err != nil {
		log.Printf("Error fetching rows: %v\n", err)
		return nil
	}

	// Get the fields
	fieldHeaders := rows[0]

	// Parse from the third row onward
	for rowIdx := 2; rowIdx < len(rows); rowIdx++ {
		row := rows[rowIdx]
		if row[0] == "" {
			continue
		}

		dateStr := row[0]
		timeStr := row[1]
		durationStr := row[2]

		// Parse the time
		startTime, err := parseExcelDateTime(dateStr, timeStr)
		if err != nil {
			log.Printf("Error parsing date time: %v\n", err)
			continue
		}

		// Parse the duration
		durationMinutes, err := strconv.Atoi(durationStr)
		if err != nil {
			log.Printf("invalid duration at row %d: %v", rowIdx+1, err)
			continue
		}
		duration := time.Duration(durationMinutes) * time.Minute

		// Iterate teams, starting from the 4 column
		for colIdx := 3; colIdx < len(row); colIdx++ {
			teamName := row[colIdx]
			if teamName == "" {
				continue
			}

			// Get the overlaying field name
			var fieldName string
			if colIdx > len(fieldHeaders) {
				continue
			}
			fieldName = fieldHeaders[colIdx]
			if fieldName == "Teams" {
				continue
			}

			// Get the team from all prefetched teams
			var team database.Team
			for _, t := range teams {
				if t.Name == teamName {
					team = t
				}
			}

			// Create the match and append it
			var match = database.Match{
				Name:          sheet,
				StartTime:     startTime,
				Duration:      duration,
				Field:         fieldName,
				League:        league,
				Institution:   teamsInstitution[teamName],
				InstitutionID: teamsInstitution[teamName].ID,
				Team:          team,
				TeamID:        team.ID,
			}

			matches = append(matches, match)
		}
	}

	return matches
}

func parseExcelDateTime(dateStr, timeStr string) (time.Time, error) {
	combined := fmt.Sprintf("%s %s", dateStr, timeStr)
	layout := "02-01-06 15:04"
	return time.Parse(layout, combined)
}
