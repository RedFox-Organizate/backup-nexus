package config

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	BackupFolders      []string
	IntervalSeconds    int
	BackblazeAccountID string
	BackblazeAppKey    string
	BackblazeBucket    string
}

func LoadEnv() error {
	return godotenv.Load()
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var folders []string
	interval := 3600 //default

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if strings.HasPrefix(line, "interval=") {
			val := strings.TrimPrefix(line, "interval=")
			sec, err := strconv.Atoi(val)
			if err != nil {
				return nil, errors.New("Invalid interval in config.txt")
			}
			interval = sec
		} else {
			folders = append(folders, line)
		}
	}

	if len(folders) == 0 {
		return nil, errors.New("No backup folders specified in config.txt")
	}

	return &Config{
		BackupFolders:      folders,
		IntervalSeconds:    interval,
		BackblazeAccountID: os.Getenv("BACKBLAZE_ACCOUNT_ID"),
		BackblazeAppKey:    os.Getenv("BACKBLAZE_APP_KEY"),
		BackblazeBucket:    os.Getenv("BACKBLAZE_BUCKET_NAME"),
	}, nil
}
