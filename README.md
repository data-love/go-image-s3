# go-image-s3

This repo will show some example for Image upload to aws S3 with golang.
There are different ways to do so and it depends on the case which is the best
solution.

## Dependencies

IF you do not have the dependencies aka imports on your local machine
=>
```
go get github.com/aws/aws-sdk-go/aws
  github.com/aws/aws-sdk-go/aws/awsutil
	github.com/aws/aws-sdk-go/aws/credentials
	github.com/aws/aws-sdk-go/aws/session
	github.com/aws/aws-sdk-go/service/s3
	github.com/joho/godotenv
```

## server-side-upload
[Documentation](server-side-upload/README.md)


## browser-based-uploads
[Documentation](browser-based-uploads/README.md)
