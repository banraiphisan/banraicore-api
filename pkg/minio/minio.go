package minio

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/tubfuzzy/banraicore-api/config"
	"log"
	"net/http"
	"net/url"
)

type MinioClient struct {
	*minio.Client
}

func sanitizeEndpoint(endpoint, port string) (string, error) {
	parsedURL, err := url.Parse(endpoint)
	if err != nil {
		return "", fmt.Errorf("❌ Invalid MinIO endpoint: %v", err)
	}

	sanitized := parsedURL.Host

	if parsedURL.Port() == "" {
		sanitized = fmt.Sprintf("%s:%s", sanitized, port)
	}

	return sanitized, nil
}

func New(cfg config.MinioConfig) (*MinioClient, error) {
	sanitizedEndpoint, err := sanitizeEndpoint(cfg.Endpoint, cfg.Port)
	if err != nil {
		return nil, err
	}
	options := minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
		Secure: cfg.UseSSL,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Proxy:           http.ProxyFromEnvironment,
		},
	}
	client, err := minio.New(sanitizedEndpoint, &options)
	if err != nil {
		log.Fatalf("❌ Failed to initialize MinIO client: %v", err)
		return nil, err
	}

	ctx := context.Background()

	exists, err := client.BucketExists(ctx, cfg.Bucket)
	if err != nil {
		return nil, fmt.Errorf("❌ Error checking bucket: %w", err)
	}

	if !exists {
		err = client.MakeBucket(ctx, cfg.Bucket, minio.MakeBucketOptions{})
		if err != nil {
			return nil, fmt.Errorf("❌ Error creating bucket: %w", err)
		}
		log.Printf("✅ MinIO bucket '%s' created", cfg.Bucket)
	}

	return &MinioClient{client}, nil
}

// GetClient returns the MinIO client instance
//func (m *MinioClient) GetClient() *minio.Client {
//	return m.Client
//}
