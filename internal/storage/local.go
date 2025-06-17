package storage

import "context"

type Storage interface {
	Upload(ctx context.Context, localFilePath string, remoteFileName string) error
}