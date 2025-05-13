package data

import (
	"github.com/AlexJarrah/go-ods"
	"gorm.io/gorm"
	"log"
	"os"
)

// StoreMatches parses the csv file and stores the matches in the database.
// Sends a message to all connected clients to reload their data
// A example of the timetable can be found in root of the directory
func StoreMatches(tmpPath string, league string, db *gorm.DB) {
	log.Println(league, db)

	// Read the file into the ods library
	data, files, err := ods.Read(tmpPath)
	if err != nil {
		log.Printf("Error opening file: %v\n", err)
	}
	// Close and remove the tmp file after the function ran
	defer func() {
		err = files.Close()
		if err != nil {
			log.Printf("Error closing file: %v\n", err)
		}

		err = os.Remove(tmpPath)
		if err != nil {
			log.Printf("Error removing file: %v\n", err)
		}
	}()

	// Read the content of the file
	data.Content = ods.Uncompress(data.Content, 20)

	// Open the first sheet
	var numberOfTables = len(data.Content.Body.Spreadsheet.Table)
	for i := 0; i < numberOfTables; i++ {
		log.Println(data.Content.Body.Spreadsheet.Table[i].Name)
	}
}
