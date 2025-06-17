package backup

import (
	"backup-nexus/internal/compress"
	"backup-nexus/internal/logger"
	"backup-nexus/internal/storage"
	"backup-nexus/internal/utils"
	"context"
	"os"
	"path/filepath"
)

type Manager struct {
	Source  *LocalSource
	Storage storage.Storage
}

func NewManager(source *LocalSource, storage storage.Storage) *Manager {
	return &Manager{
		Source:  source,
		Storage: storage,
	}
}

func (m *Manager) Run() error {
	validFolders := m.Source.Validate()
	if len(validFolders) == 0 {
		logger.Warn("No valid folders to backup.")
		return nil
	}

	if err := utils.EnsureDir("tmp"); err != nil {
		return err
	}

	backupName := utils.FormatBackupFileName()
	outputPath := filepath.Join("tmp", backupName)

	logger.Info("Compressing folders...")
	if err := compress.CompressFolders(validFolders, outputPath); err != nil {
		return err
	}
	logger.Success("Folders compressed successfully.")

	logger.Info("Uploading to storage...")
	err := m.Storage.Upload(context.Background(), outputPath, backupName)
	if err != nil {
		logger.Error("Upload failed: " + err.Error())
		return err
	}
	logger.Success("Upload completed.")

	_ = os.Remove(outputPath)
	return nil
}
