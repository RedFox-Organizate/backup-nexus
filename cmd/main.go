package main

import (
	"backup-nexus/internal/backup"
	"backup-nexus/internal/config"
	"backup-nexus/internal/logger"
	"backup-nexus/internal/scheduler"
	"backup-nexus/internal/storage"
	"fmt"
	"os"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		fmt.Println("Failed to load .env:", err)
		return
	}
	logger.InitLogger()
	defer logger.CloseLogger()

	cfg, err := config.LoadConfig("configs/config.txt")
	if err != nil {
		logger.Error("Failed to load config.txt: " + err.Error())
		return
	}

	if len(os.Args) > 1 {
		cfg.BackupFolders = os.Args[1:]
		logger.Info("Overriding folders from CLI args")
	}

	logger.Info("Backup folders: " + logger.PrettyList(cfg.BackupFolders))
	logger.Info("Backup interval: " + logger.PrettyInt(cfg.IntervalSeconds) + " seconds")

	b2, err := storage.NewBackblazeB2Storage(cfg.BackblazeAccountID, cfg.BackblazeAppKey, cfg.BackblazeBucket)
	if err != nil {
		logger.Error("Failed to init Backblaze storage: " + err.Error())
		return
	}

	source := backup.NewLocalSource(cfg.BackupFolders)
	manager := backup.NewManager(source, b2)

	s := scheduler.NewScheduler(cfg.IntervalSeconds, func() error {
		return manager.Run()
	})
	s.Start()

	select {} // sonsuz bekleme
}
