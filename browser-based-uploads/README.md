# browser-based-uploads
In this scenario we will create a link where the user is able to post his image.

## ENV File

Copy the env.sample from the main folder to .env in this subfolder (server-side-uploads)
Fill in your credentials there.
IMPORTANT: Never make them public.


## S3 CORS

```
<CORSConfiguration>
  <CORSRule>
    <AllowedOrigin>*</AllowedOrigin>
    <AllowedMethod>GET</AllowedMethod>
    <AllowedMethod>PUT</AllowedMethod>
    <MaxAgeSeconds>3000</MaxAgeSeconds>
    <AllowedHeader>Authorization</AllowedHeader>
  </CORSRule>
</CORSConfiguration>
```


## Docker usage
You need to be in the subfolder (server-side-uploads) since the . will pic up the
current directory.


```
docker build -t browser-based-uploads .
docker run browser-based-uploads
```

## Upload Image via pre-signed url.

If you specify the ContentType in the PutObject you need to curl
```
curl -v -H 'Content-Type: image/png' -T ./image.png '<signed url>'
```

## Bucket Policy Enhancement
You might want restrict your users from uploading all kinds of files and/or sizes.
That must be configured in the Bucket Policy.


## EXAMPLE (without ContentType header):

Put URL:
```
curl -v -T ./image.png https://go-image-s3.s3.amazonaws.com/9ba25d0e-0298-441c-a71e-15ba93a18805?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Content-Sha256=UNSIGNED-PAYLOAD&X-Amz-Credential=AKIAJMC27RDOO4HO5EIQ%2F20180513%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20180513T084549Z&X-Amz-Expires=900&X-Amz-SignedHeaders=content-type%3Bhost&X-Amz-Signature=aa253d3cb939b1098fa0ed9940a2115cd362d848f9cdffa5b8eede1e7f1be003
```

Get URL:
```
https://go-image-s3.s3.amazonaws.com/9ba25d0e-0298-441c-a71e-15ba93a18805
```
