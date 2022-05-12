package main

import (
	"context"
	"os"

	"cloud.google.com/go/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"lalu-storage/config"
	"lalu-storage/handlers"
	"lalu-storage/helpers"
)

const (
	projectId = "lalu-storage"
	bucketName = "lalu-data-storage"

	// GCS directories paths
	songsDirPath = "songs/"
	albumCoversDirPath = "album-covers/"
	playlistCoversDirPath = "playlist-covers/"
	profilePicsDirPath = "profile-pics/"
	eventPicsDirPath = "event-pics/"
)

func init(){
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	GCSCredentials := os.Getenv("GCS_CREDENTIALS")

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", GCSCredentials) // FILL IN WITH YOUR FILE PATH
	client, err := storage.NewClient(context.Background())

	if err != nil {
		panic("Failed to create GCS client")
	}

	helpers.Uploader = &helpers.ClientUploader{
		Cl:         client,
		BucketName: bucketName,
		ProjectId:  projectId,
		WorkingPath: "",
	}

}

func setupRoutes(app *fiber.App){
	// Songs routes
	app.Get("/songs", func(ctx *fiber.Ctx) error {
		return handlers.FetchFile(ctx, songsDirPath)
	})

	app.Put("/songs", func(ctx *fiber.Ctx) error {
		return handlers.UploadFile(ctx, songsDirPath)
	})

	app.Delete("/songs", func(ctx *fiber.Ctx) error {
		return handlers.DeleteFile(ctx, songsDirPath)
	})

	// Album covers routes
	app.Get("/album-covers", func(ctx *fiber.Ctx) error {
		return handlers.FetchFile(ctx, albumCoversDirPath)
	})

	app.Post("/album-covers", func(ctx *fiber.Ctx) error {
		return handlers.UploadFile(ctx, albumCoversDirPath)
	})

	app.Put("/album-covers", func(ctx *fiber.Ctx) error {
		return handlers.UploadFile(ctx, albumCoversDirPath)
	})

	app.Delete("/album-covers", func(ctx *fiber.Ctx) error {
		return handlers.DeleteFile(ctx, albumCoversDirPath)
	})

	// Playlists covers routes
	app.Get("/playlists-covers", func(ctx *fiber.Ctx) error {
		return handlers.FetchFile(ctx, playlistCoversDirPath)
	})

	app.Post("/playlists-covers", func(ctx *fiber.Ctx) error {
		return handlers.UploadFile(ctx, playlistCoversDirPath)
	})

	app.Put("/playlists-covers", func(ctx *fiber.Ctx) error {
		return handlers.UploadFile(ctx, playlistCoversDirPath)
	})

	app.Delete("/playlists-covers", func(ctx *fiber.Ctx) error {
		return handlers.DeleteFile(ctx, playlistCoversDirPath)
	})

	// Profile pics routes
	app.Get("/profile-pics", func(ctx *fiber.Ctx) error {
		return handlers.FetchFile(ctx, profilePicsDirPath)
	})

	app.Post("/profile-pics", func(ctx *fiber.Ctx) error {
		return handlers.UploadFile(ctx, profilePicsDirPath)
	})

	app.Put("/profile-pics", func(ctx *fiber.Ctx) error {
		return handlers.UploadFile(ctx, profilePicsDirPath)
	})

	app.Delete("/profile-pics", func(ctx *fiber.Ctx) error {
		return handlers.DeleteFile(ctx, profilePicsDirPath)
	})

	// Event pics routes
	app.Get("/event-pics", func(ctx *fiber.Ctx) error {
		return handlers.FetchFile(ctx, eventPicsDirPath)
	})

	app.Post("/event-pics", func(ctx *fiber.Ctx) error {
		return handlers.UploadFile(ctx, eventPicsDirPath)
	})

	app.Put("/event-pics", func(ctx *fiber.Ctx) error {
		return handlers.UploadFile(ctx, eventPicsDirPath)
	})

	app.Delete("/event-pics", func(ctx *fiber.Ctx) error {
		return handlers.DeleteFile(ctx, eventPicsDirPath)
	})
}

func main(){
	app := fiber.New(fiber.Config{
		BodyLimit: 15 * 1024 * 1024,
	})
	
	setupRoutes(app)
	
	go app.Listen(":3000")

	// RabbitMQ config
	channel := config.RabbitMQConn()
	handlers.DequeueSongs(channel)
	
}