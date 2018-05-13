# server-side-uploads

## ENV File

Copy the env.sample from the main folder to .env in this subfolder (server-side-uploads)
Fill in your credentials there.
IMPORTANT: Never make them public.

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

## Docker usage
You need to be in the subfolder (server-side-uploads) since the . will pic up the
current directory.


```
docker build -t server-side-uploads .
docker run server-side-uploads
```

## Bucket on S3

Create a Bucket on s3 https://s3.console.aws.amazon.com/s3/home?region=us-east-1
This will allow you to create a bucket on the Region `us-east-1`.


## AWS BUCKET POLICIES
My Bucket Name is "go-image-s3". You need to replace that for you bucket name.
I want that everyone can have a look at my uploaded images.
Therefor i have to make it public. The `"Action": "s3:GetObject"` grants read permissions.

```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "AddPerm",
            "Effect": "Allow",
            "Principal": "*",
            "Action": "s3:GetObject",
            "Resource": "arn:aws:s3:::go-image-s3/*"
        }
    ]
}
```

Image will be later accessible at:
https://s3.amazonaws.com/go-image-s3/media/test.png

If you upload an other image with the same filename it will overwrite the current
one in aws. So it would make sense to use an additional UUID in the file name so we
make sure files get not overwritten.
