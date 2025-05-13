package data

import (
	"errors"
	"fmt"
	"github.com/AlexJarrah/go-ods"
	"gorm.io/gorm"
	"log"
	"rcjv-app/backend/database"
	"rcjv-app/backend/util"
)

// GenerateODS generates the appropriate ods for creating the timetable.
// It automatically enters all teams and fields, so you can just pick.
// Generated via admin panel
func GenerateODS(league string, db *gorm.DB) (string, error) {
	// Read the file into the ods library
	data, files, err := ods.Read("ods/template.ods")
	if err != nil {
		log.Printf("Error opening file: %v\n", err)
	}
	defer files.Close()

	// Read the content of the file
	data.Content = ods.Uncompress(data.Content, 20)

	// Open the first sheet
	mainSheet := data.Content.Body.Spreadsheet.Table[0]

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

	// Put the fields and the teams into the sheet
	for index, field := range fields {
		mainSheet.TableRow[1].TableCell[index+3].P = field.Name
	}

	// Write a new temp file
	var tmpPath = fmt.Sprintf("./tmp/%s_template_%s.ods", league, util.GenerateFileTempToken())

	data.Content.Body.Spreadsheet.Table[0] = mainSheet
	if err := ods.Write(tmpPath, data, files); err != nil {
		return "", errors.New(fmt.Sprintf("Error writing file: %v", err))
	}

	log.Printf("Successfully generated ODS for league: %s\n", league)
	log.Println(tmpPath)
	return tmpPath, nil
}
