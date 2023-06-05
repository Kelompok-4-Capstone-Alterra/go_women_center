package helper

import (
	"context"
	"mime/multipart"
	"path/filepath"
	"regexp"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/h2non/filetype"
)

type Image interface {
	IsImageValid(fh *multipart.FileHeader) bool
	UploadImageToS3(fh *multipart.FileHeader) (string, error)
	DeleteImageFromS3(path_link string) error
}

type image struct{
	bucket string
}

func NewImage(bucket string) Image {
	return &image{bucket}
}

func(i *image) IsImageValid(fh *multipart.FileHeader) bool {

	file, err := fh.Open()

	if err != nil {
		return false
	}
	
	if fh.Size > 2 * 1024 * 1024 { // 2 MB
		return false
	}
	
	defer file.Close()
	
	buff := make([]byte, 512)

	if _, err := file.Read(buff); err != nil {
		return false
	}

	kind, _ := filetype.Match(buff)

	if kind == filetype.Unknown {
		return false
	}
	
	if kind.Extension != "jpg" && kind.Extension != "jpeg" && kind.Extension != "png" {
		return false
	}

	if kind.MIME.Type != "image"{
		return false
	}

	return true

}

func(i *image) UploadImageToS3(fh *multipart.FileHeader) (string, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	
	if err != nil {
		return "", err
	}
	
	file, err := fh.Open()

	if err != nil {
		return "", err
	}
	
	defer file.Close()

	client := s3.NewFromConfig(cfg)

	ext := filepath.Ext(fh.Filename)

	uploader := manager.NewUploader(client)

	uuid, _ := NewGoogleUUID().GenerateUUID()
	currentTime := time.Now().UnixNano()

	newFileName :=  uuid +  "-" + strconv.Itoa(int(currentTime)) + ext

	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(i.bucket),
		Key:    aws.String(newFileName),
		Body:   file,
	})

	if err != nil {
		return "", err
	}

	return result.Location, nil

}

func getFileName(path_link string) string {
	pattern := `/([^/]+)$`

	regex := regexp.MustCompile(pattern)

	matches := regex.FindStringSubmatch(path_link)

	if len(matches) < 2 {
		return ""
	}

	return matches[1]
}

func(i *image) DeleteImageFromS3(path_link string) error {
	
	filename := getFileName(path_link)

	if filename == "" {
		return nil
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	
	if err != nil {
		return err
	}

	client := s3.NewFromConfig(cfg)

	input := &s3.DeleteObjectInput{
		Bucket: aws.String(i.bucket),
		Key: 	aws.String(filename),
	}

	_, err = client.DeleteObject(context.TODO(), input)
	
	if err != nil {
		return err
	}

	return nil
}