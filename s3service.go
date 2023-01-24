package main

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
)

type S3Service struct {
	Client *s3.Client
}

func NewS3Service(c *s3.Client) *S3Service {
	return &S3Service{
		Client: c,
	}
}

func (s3s S3Service) ListBuckets() ([]s3types.Bucket, error) {
	result, err := s3s.Client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	var buckets []s3types.Bucket
	if err != nil {
		log.Printf("Couldn't list buckets for your account. Here's why: %v\n", err)
	} else {
		buckets = result.Buckets
	}
	return buckets, err
}

func (s3s S3Service) BucketExists(bucketName string) (bool, error) {
	_, err := s3s.Client.HeadBucket(context.TODO(), &s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})

	exists := true

	if err != nil {
		var apiError smithy.APIError
		if errors.As(err, &apiError) {
			switch apiError.(type) {
			case *s3types.NotFound:
				log.Printf("Bucket %v is available.\n", bucketName)
				exists = false
				err = nil
			default:
				log.Printf("Either you don't have access to bucket %v or another error occurred. "+
					"Here's what happened: %v\n", bucketName, err)
			}
		}
	} else {
		log.Printf("Bucket %v exists and you already own it.", bucketName)
	}

	return exists, err
}

func (s3s S3Service) ListBucketObjects(bucketName string) ([]s3types.Object, error) {
	result, err := s3s.Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	var contents []s3types.Object
	if err != nil {
		log.Printf("Couldn't list objects in bucket %v. Here's why: %v\n", bucketName, err)
	} else {
		contents = result.Contents
	}
	return contents, err
}

func (s3s S3Service) GetObjectAttributes(bucketName string, key string) (*ObjectAttributes, error) {
	hoa, err := s3s.Client.HeadObject(context.TODO(), &s3.HeadObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})

	if err != nil {
		log.Printf("Couldn't get objects attributes %v. Here's why: %v\n", key, err)
		return nil, err
	}

	attr := &ObjectAttributes{
		ContentLength: hoa.ContentLength,
		ContentType:   *hoa.ContentType,
		ETag:          *hoa.ETag,
		LastModified:  *hoa.LastModified,
	}

	return attr, nil
}

func (s3s S3Service) DeleteObjects(bucketName string, objectKeys []string) error {
	var objectIds []s3types.ObjectIdentifier
	for _, key := range objectKeys {
		objectIds = append(objectIds, s3types.ObjectIdentifier{Key: aws.String(key)})
	}
	_, err := s3s.Client.DeleteObjects(context.TODO(), &s3.DeleteObjectsInput{
		Bucket: aws.String(bucketName),
		Delete: &s3types.Delete{Objects: objectIds},
	})
	if err != nil {
		log.Printf("Couldn't delete objects from bucket %v. Here's why: %v\n", bucketName, err)
	}
	return err
}
