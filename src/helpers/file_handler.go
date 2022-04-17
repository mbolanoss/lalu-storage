package helpers

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
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

func (c *ClientUploader) UploadFile(file multipart.File, fileName string, uploadPath string) error {
	c.WorkingPath = uploadPath
	
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	writer := c.Cl.Bucket(c.BucketName).Object(c.WorkingPath + fileName).NewWriter(ctx)

	if _, err := io.Copy(writer, file); err != nil {
		return fmt.Errorf("error while copying file to object handler")
	}
	if err := writer.Close(); err != nil {
		return fmt.Errorf("error closing object handler")
	}

	return nil
}

func (c *ClientUploader) DeleteFile(fileName string, deletePath string) error {
	c.WorkingPath = deletePath

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()
	
	o := c.Cl.Bucket(c.BucketName).Object(c.WorkingPath + fileName)
	
	err := o.Delete(ctx)
	
	if err != nil {
		return fmt.Errorf("error while deleting file")
	}
	
	return nil
}

func (c *ClientUploader) FetchFile(fileName string, fetchPath string) ([]byte, error) {
	c.WorkingPath = fetchPath

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	rc, err := c.Cl.Bucket(c.BucketName).Object(c.WorkingPath + fileName).NewReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("error while creating file reader")
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, fmt.Errorf("error while reading file data")
	}
	return data, nil
}