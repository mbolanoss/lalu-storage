package helpers

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"time"

	"cloud.google.com/go/storage"
)

type ClientUploader struct {
	Cl *storage.Client
	ProjectId string
	BucketName string
	WorkingPath string
}

var Uploader *ClientUploader

func (c *ClientUploader) UploadFile(file multipart.File, object string, uploadPath string) error {
	c.WorkingPath = uploadPath
	
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	writer := c.Cl.Bucket(c.BucketName).Object(c.WorkingPath + object).NewWriter(ctx)

	if _, err := io.Copy(writer, file); err != nil {
		return fmt.Errorf("error while copying file to object handler")
	}
	if err := writer.Close(); err != nil {
		return fmt.Errorf("error closing object handler")
	}

	return nil
}