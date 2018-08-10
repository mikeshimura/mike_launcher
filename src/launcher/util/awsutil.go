package util

import (
	"fmt"
	"os"
"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func S3Download(sbucket string, s3key string, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic("Failed to create file" + err.Error())
	}
	defer file.Close()
	downloader := s3manager.NewDownloader(session.New(&aws.Config{Region: aws.String("us-east-1")}))
	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(sbucket),
			Key:    aws.String(s3key),
		})
	if err != nil {
		fmt.Println("Failed to download file", err)
		return
	}
	fmt.Println("Downloaded file", file.Name(), numBytes, "bytes")

}

func ListKeys(bucket string) []*s3.Object{
	svc := s3.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})

    params := &s3.ListObjectsInput{
        Bucket: aws.String(bucket),
    }

    resp, _ := svc.ListObjects(params)
   return resp.Contents
}
