package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
)

// Sess is variable
var Sess *session.Session

// SVC is variable
var SVC *s3.S3

// Bucket is variable
var bucket string = "on-board-project"

func init() {
	err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal(err)
	}
	awsID := os.Getenv("AWSID")
	awsKey := os.Getenv("AWSKEY")
	creds := credentials.NewStaticCredentials(awsID, awsKey, "")
	Sess, err = session.NewSession(&aws.Config{
		Credentials: creds,
		Region:      aws.String("ap-northeast-1"),
	})
	if err != nil {
		panic(err)
	}
	SVC = s3.New(Sess)
	fmt.Println("----------access---successfully------")
}

// UploadFileToBucket is repository of aws funcitons
// func UploadFileToBucket(filename string) (*s3manager.UploadOutput, error) {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	uploader := s3manager.NewUploader(Sess)
// 	res, err := uploader.Upload(&s3manager.UploadInput{
// 		Bucket: aws.String(bucket),
// 		Key:    aws.String(filename),
// 		Body:   file,
// 	})
// 	return res, err
// }

// func DeleteFileByBucket() {
// 	bucket := "bucket-name"
// 	obj := "file-to-delete"

// 	_, err = svc.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(bucket), Key: aws.String(obj)})

// 	err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
// 		Bucket: aws.String(bucket),
// 		Key:    aws.String(obj),
// 	})
// }
