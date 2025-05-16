package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	router "github.com/AbdallahAskar1/go-cloud-file-service/routes"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(os.Getenv("AWS_REGION")),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)

	bucketName := "aasksar"
	_, err = client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		if !isBucketAlreadyExists(err) {
			log.Fatal(err)
		}
	}
	fmt.Printf("Bucket %s is ready\n", bucketName)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "Hello world"})
	})

	router.RegisterRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}

func isBucketAlreadyExists(err error) bool {
	var s3Err *s3.ResponseError
	return errors.As(err, &s3Err)
}
