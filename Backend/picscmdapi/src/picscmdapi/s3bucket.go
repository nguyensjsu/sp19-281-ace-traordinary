package main

import (
	"log"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

//InsertIntoS3 insert image to s3 bucket
func InsertIntoS3(filename string, file multipart.File) string {
	res := ""
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewStaticCredentials(ACCESSKEYS3, SECRETKEYS3, "")},
	)
	uploader := s3manager.NewUploader(sess)
	result, s3err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(BUCKETNAME),
		Key:    aws.String(filename),
		Body:   file,
	})
	if s3err != nil {
		log.Fatalf("Unable to upload %q to %q, %v", filename, BUCKETNAME, err)
	}

	log.Println("Successfully uploaded image location" + result.Location)
	res = result.Location
	return res
}

/**
//DeleteFromS3 delete an image from S3 bucket
func DeleteFromS3(filename string) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewStaticCredentials(ACCESSKEYS3, SECRETKEYS3, "")},
	)
}

//DeleteMultipleObjectsFromS3 delete files from S3Bucket
func DeleteMultipleObjectsFromS3() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewStaticCredentials(ACCESSKEYS3, SECRETKEYS3, "")},
	)

}
**/
