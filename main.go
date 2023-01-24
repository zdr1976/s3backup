package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	awsEndpoint = "http://localhost:9000"
	awsRegion   = "us-east-1"
	accessKey   = "admin"
	secretKey   = "P@ssw0rd123"
)

func listBuckets(s *S3Service) {
	buckets, err := s.ListBuckets()
	if err != nil {
		fmt.Println("Couldn't list buckets for your account:", err)
	}
	for _, bucket := range buckets {
		fmt.Printf("Bucket name: %s, Creation date: %v\n", *bucket.Name, bucket.CreationDate)
	}
}

func listObjects(s *S3Service, bucket string) {
	objects, err := s.ListBucketObjects(bucket)
	if err != nil {
		fmt.Println("Couldn't list objects in bucket:", err)
	}
	for _, object := range objects {
		fmt.Printf("Key: %s, ETag: %v, Size: %d, LastModified: %v\n", *object.Key, *object.ETag, object.Size, *object.LastModified)
		getObjectInfo(s, bucket, *object.Key)
	}
}

func getObjectInfo(s *S3Service, bucket string, key string) {
	attr, err := s.GetObjectAttributes(bucket, key)
	if err != nil {
		fmt.Println("Couldn't get objects attributes:", err)
	}

	fmt.Printf("  Key: %s, ContentType: %s, ContentLength: %d, ETag %s: , ListModified: %v\n", key, attr.ContentType, attr.ContentLength, attr.ETag, attr.LastModified)
}

func main() {
	customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		if awsEndpoint != "" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           awsEndpoint,
				SigningRegion: awsRegion,
			}, nil
		}

		// returning EndpointNotFoundError will allow the service to fallback to its default resolution
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	awsCredential := credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")

	awsCfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(awsRegion),
		config.WithEndpointResolver(customResolver),
		config.WithCredentialsProvider(awsCredential),
	)
	if err != nil {
		log.Fatalf("Cannot load the AWS configs: %s", err)
	}

	// Create the resource client
	client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	s := NewS3Service(client)

	// List buckets
	listBuckets(s)

	// List bucket
	bucket := "data"
	listObjects(s, bucket)

}
