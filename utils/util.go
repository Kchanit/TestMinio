package utils

import (
	"log"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Credentials struct {
	AccessKeyID     string
	SecretAccessKey string
}

func UploadInvoice() {
	// endpoint := "https://minio.ratchaphon1412.co"
	endpoint := "localhost:9000"
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY")
	secretAccessKey := os.Getenv("MINIO_SECRET_KEY")
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	// Initialize default config
	if err != nil {
		log.Fatalln("Failed to initialize minio client", err)
	}

	log.Printf("minioClient: %v\n", minioClient)

	bucketName := "pixelmanstorage"
	location := "us-east-1"
	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln("Failed to create bucket", err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	objectName := "invoice-" + time.Now().Format("2006-01-02")
	filePath := "./test.pdf"
	contentType := "image/pdf"

	// Upload the test file with FPutObject
	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln("Failed to upload file", err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
	// save file name to database
}
