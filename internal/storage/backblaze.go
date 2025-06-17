package storage

import (
	"context"
	"os"

	"github.com/kurin/blazer/b2"
)

type BackblazeB2Storage struct {
	accountID  string
	appKey     string
	bucketName string
	client     *b2.Client
	bucket     *b2.Bucket
}

func NewBackblazeB2Storage(accountID, appKey, bucketName string) (*BackblazeB2Storage, error) {
	ctx := context.Background()
	client, err := b2.NewClient(ctx, accountID, appKey)
	if err != nil {
		return nil, err
	}

	bucket, err := client.Bucket(ctx, bucketName)
	if err != nil {
		return nil, err
	}

	return &BackblazeB2Storage{
		accountID:  accountID,
		appKey:     appKey,
		bucketName: bucketName,
		client:     client,
		bucket:     bucket,
	}, nil
}

func (b *BackblazeB2Storage) Upload(ctx context.Context, localFilePath string, remoteFileName string) error {
	f, err := os.Open(localFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	obj := b.bucket.Object(remoteFileName)
	w := obj.NewWriter(ctx) 

	defer func() {
		_ = w.Close()
	}()

	if _, err := w.ReadFrom(f); err != nil {
		return err
	}

	if err := w.Close(); err != nil {
		return err
	}

	return nil
}
