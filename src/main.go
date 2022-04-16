package main

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/gofiber/fiber/v2"
)

const (
	projectId = "lalu-storage"
	bucketName = "lalu-data-storage"
)

type ClientUploader struct {
	cl *storage.Client
	projectId string
	bucketName string
	uploadPath string
}

var uploader *ClientUploader

func init(){
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/home/miguel/Universidad/Arquisoft/Proyecto/lalu-storage/src/lalu-storage-b7d06ebc57b0.json") // FILL IN WITH YOUR FILE PATH
	client, err := storage.NewClient(context.Background())

	if err != nil {
		panic("Failed to create GCS client")
	}

	uploader = &ClientUploader{
		cl:         client,
		bucketName: bucketName,
		projectId:  projectId,
		uploadPath: "songs/",
	}
}

func main(){
	
	app := fiber.New()

	app.Post("/upload", uploadFile)

	app.Listen(":3000")
}

func uploadFile(ctx *fiber.Ctx) error {

	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString("Error while parsing file")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString("Error while opening file")
	}
	defer file.Close()
	
	err = uploader.UploadFile(file, fileHeader.Filename)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString("Error while uploading file to GCS")
		
	}

	return ctx.SendStatus(http.StatusOK)
}

func (c *ClientUploader) UploadFile(file multipart.File, object string) error {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	// Upload an object with storage.Writer.

	wc := c.cl.Bucket(c.bucketName).Object(c.uploadPath + object).NewWriter(ctx)

	if _, err := io.Copy(wc, file); err != nil {
		return fmt.Errorf("error while copying file to object handler")
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("error closing object handler")
	}

	return nil
}