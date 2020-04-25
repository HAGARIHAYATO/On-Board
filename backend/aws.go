package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Sess is variable
var Sess *session.Session

// SVC is variable
var SVC *s3.S3

// Bucket is variable
var bucket string = "on-board-pub"

func init() {
	// Product ENV
	Sess := session.Must(session.NewSession())
	SVC := ec2.New(
		Sess,
		aws.NewConfig().WithRegion("ap-northeast-1"),
	)
	fmt.Println(SVC.DescribeInstances(nil))
	// DEV ENV
	// err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// awsID := os.Getenv("AWSID")
	// awsKey := os.Getenv("AWSKEY")
	// creds := credentials.NewStaticCredentials(awsID, awsKey, "")
	// Sess, err = session.NewSession(&aws.Config{
	// 	Credentials: creds,
	// 	Region:      aws.String("ap-northeast-1"),
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// SVC = s3.New(Sess)
	fmt.Println("----------access---successfully------")
}

// UploadFileToBucket is repository of aws funcitons
func UploadFileToBucket(filename string) (*s3manager.UploadOutput, error) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	uploader := s3manager.NewUploader(Sess)
	res, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   file,
	})
	return res, err
}

// DeleteFileByBucket is repository of aws functions
func DeleteFileByBucket(filename string) error {
	trimedName := strings.Trim(filename, "https://on-board-pub.s3.ap-northeast-1.amazonaws.com/")
	_, err = SVC.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(bucket), Key: aws.String(trimedName)})
	err = SVC.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(trimedName),
	})
	return err
}
