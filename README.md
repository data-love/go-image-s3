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

Documentation from AWS [https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-UsingHTTPPOST.html](https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-UsingHTTPPOST.html)

## server-side-upload
[Documentation](server-side-upload/README.md)


## browser-based-uploads (pre-signed-url [https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/s3-example-presigned-urls.html](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/s3-example-presigned-urls.html))
[Documentation](browser-based-uploads/README.md)
