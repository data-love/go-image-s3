package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	id := os.Getenv("AWS_ACCESS_KEY_ID")
	key := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_S3_REGION")
	bucket := os.Getenv("AWS_S3_BUCKET")
	token := ""

	creds := credentials.NewStaticCredentials(id, key, token)
	_, err = creds.Get()
	if err != nil {
		fmt.Printf("bad credentials: %s", err)
	}
	cfg := aws.NewConfig().WithRegion(region).WithCredentials(creds)
	uuid := uuid.Must(uuid.NewV4()).String()
	svc := s3.New(session.New(), cfg)
	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(uuid),
		// ContentType: aws.String("image/png"), This is optional, but in my case i want to vary between jpeg, jpg and png
	})
	urlStr, err := req.Presign(15 * time.Minute)

	if err != nil {
		log.Println("Failed to sign request", err)
	}

	log.Println("The URL is", urlStr)
}
