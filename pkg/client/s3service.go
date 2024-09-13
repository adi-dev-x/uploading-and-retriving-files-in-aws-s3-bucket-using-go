package client

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"time"

	"s3service/pkg/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconf "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var s3Client *s3.Client

type S3exposeGateway interface {
	UploadFile(file []byte, key string) error
	RetriveFile(file string) (string, error)
	Init()
}

type S3pathway struct {
	Conf config.Config
}

func (s *S3pathway) Init() {
	awsCfg, err := awsconf.LoadDefaultConfig(context.TODO(),
		awsconf.WithRegion("us-east-2"),
		awsconf.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			s.Conf.AccesKey, s.Conf.SecretKey, "")),
	)

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	s3Client = s3.NewFromConfig(awsCfg)
}
func (s *S3pathway) UploadFile(file []byte, key string) error {
	bucketName := s.Conf.BucketName

	_, err := s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   bytes.NewReader(file),

		ContentLength: aws.Int64(int64(len(file))),
		ContentType:   aws.String("application/octet-stream"),
	})
	if err != nil {
		return err
	}
	fmt.Println("wnded the service")
	return nil
}
func (s *S3pathway) RetriveFile(file string) (string, error) {
	presigner := s3.NewPresignClient(s3Client)
	bucketName := s.Conf.BucketName
	req, err := presigner.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket:                     aws.String(bucketName),
		Key:                        aws.String(file),
		ResponseContentDisposition: aws.String("inline"),
	}, s3.WithPresignExpires(15*time.Minute))
	if err != nil {
		return "", err

	}
	return req.URL, nil
}
