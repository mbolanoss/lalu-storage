package handlers

import (
	"net/http"

	"lalu-storage/helpers"

	"github.com/gofiber/fiber/v2"
)

func UploadFile(ctx *fiber.Ctx, path string) error {

	fileHeader, err := ctx.FormFile("file")

	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString("Error while parsing file")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString("Error while opening file")
	}
	defer file.Close()
	
	err = helpers.Uploader.UploadFile(file, fileHeader.Filename, path)
	
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString("Error while uploading file to GCS")
		
	}

	return ctx.SendStatus(http.StatusOK)
}

func FetchFile(ctx *fiber.Ctx, path string) error {

	fileName := ctx.Query("fileName")

	if fileName == "" {
		return ctx.Status(http.StatusBadRequest).SendString("No file name query param found")
	}
	
	file, err := helpers.Uploader.FetchFile(fileName, path)
	
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString("Error fetching file from GCS")
	}

	return ctx.Send(file)
}

func DeleteFile(ctx *fiber.Ctx, path string) error {
	
	var err error

	fileName := ctx.Query("fileName")
	if fileName == "" {
		return ctx.Status(http.StatusBadRequest).SendString("File name not found in query params")
	}

	err = helpers.Uploader.DeleteFile(fileName, path)
	
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString("Error while deleting file in GCS")
		
	}

	return ctx.SendStatus(http.StatusOK)
}