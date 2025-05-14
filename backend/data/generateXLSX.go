package data

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"log"
	"rcjv-app/backend/database"
	"rcjv-app/backend/util"
)

// GenerateXLSX generates the appropriate ods for creating the timetable.
// It automatically enters all teams and fields, so you can just pick.
// Generated via admin panel
func GenerateXLSX(league string, db *gorm.DB) (string, error) {
	var (
		characters = []string{"D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

		err error
	)
	// Read the file into the excelize library
	file, err := excelize.OpenFile("excel/template.xlsx")
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error reading template file: %v\n", err))
	}
	defer func(file *excelize.File) {
		err := file.Close()
		if err != nil {
			log.Printf("Error closing file: %v\n", err)
		}
	}(file)

	// Receive all the necessary data from the database
	var fields []database.Field
	err = db.Where("league = ?", league).Find(&fields).Error
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error finding fields: %v", err))
	}

	var teams []database.Team
	err = db.Where("league = ?", league).Find(&teams).Error
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error finding teams: %v", err))
	}

	// Get all sheets
	sheets := file.GetSheetList()
	// Put the fields and the teams into the sheet
	for index, field := range fields {
		err = file.SetCellValue(sheets[0], fmt.Sprintf("%s1", characters[index]), field.Name)
		if err != nil {
			return "", errors.New(fmt.Sprintf("Error setting cell value (field): %v\n", err))
		}
	}

	for index, team := range teams {
		err = file.SetCellValue(sheets[0], fmt.Sprintf("AC%d", index+2), team.Name)
		if err != nil {
			return "", errors.New(fmt.Sprintf("Error setting cell value (team): %v\n", err))
		}
	}

	// Write a new temp file
	var tmpPath = fmt.Sprintf("tmp/%s_template_%s.xlsx", league, util.GenerateFileTempToken())

	if err := file.SaveAs(tmpPath); err != nil {
		return "", errors.New(fmt.Sprintf("Error saving file: %v\n", err))
	}

	log.Printf("Successfully generated XLSX for league: %s\n", league)
	return tmpPath, nil
}
