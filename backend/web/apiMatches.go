package web

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"log"
	"mime/multipart"
	"os"
	"rcjv-app/backend/data"
	"rcjv-app/backend/database"
	"rcjv-app/backend/util"
	"strconv"
)

func (a *API) uploadMatches(c *fiber.Ctx) error {
	if !util.CheckAuth(c.GetReqHeaders(), a.PSQL) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{})
	}
	// Get the league the games should be uploaded to
	var league = c.Params("league")
	// Parse the file header
	fileHeader, err := c.FormFile("matches")
	if err != nil {
		log.Printf("Error getting file header: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON("Invalid file header")
	}

	// Open the file
	file, err := fileHeader.Open()
	if err != nil {
		log.Printf("Error opening file: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error opening file")
	}

	// Create a temp file, so the ods library can open it
	path, err := storeTmpFile(file)
	if err != nil {
		log.Printf("Error storing file: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error storing file")
	}

	// Call the function for storing the matches async
	// If there is an error, there will be sent an error message to the frontend via the websocket
	// with an identifier key for the frontend to detect for whom the error message is
	go data.StoreMatches(path, league, a.PSQL)

	// ToDo(Send update signal to all clients)

	return c.Status(fiber.StatusOK).JSON("Upload Successfully")
}

func storeTmpFile(file multipart.File) (string, error) {
	// So multiple uploads work, I give each file a random name.
	// The tmp directory is writable inside the docker container, which is very important
	var path = fmt.Sprintf("./tmp/matches_tmp_%s.ods", util.GenerateFileTempToken())
	out, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return "", err
	}

	return path, nil
}

func (a *API) generateODS(c *fiber.Ctx) error {
	if !util.CheckAuth(c.GetReqHeaders(), a.PSQL) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{})
	}
	league := c.Params("league")
	if league == "" {
		return c.Status(fiber.StatusBadRequest).JSON("invalid league")
	}

	ods, err := data.GenerateXLSX(league, a.PSQL)
	if err != nil {
		log.Print(err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error generating ODS")
	}

	log.Println(ods)
	return c.SendStatus(fiber.StatusOK)
}

func (a *API) updateMatch(c *fiber.Ctx) error {
	if !util.CheckAuth(c.GetReqHeaders(), a.PSQL) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{})
	}
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil || id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON("invalid id")
	}

	// ToDo(Update match)
	// ToDo(Send update signal to all clients)

	return c.Status(fiber.StatusOK).JSON("Update Successfully")
}

func (a *API) deleteMatch(c *fiber.Ctx) error {
	if !util.CheckAuth(c.GetReqHeaders(), a.PSQL) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{})
	}
	// Parse and check the id
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil || id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON("invalid id")
	}

	// Delete the match
	err = a.PSQL.Delete(database.Match{}, id).Error
	if err != nil {
		log.Printf("Error deleting match: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON("Error deleting match")
	}

	// ToDo(Send update signal to all clients)

	return c.Status(fiber.StatusOK).JSON("Delete Successfully")
}
