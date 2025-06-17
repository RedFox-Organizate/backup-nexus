package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func FormatBackupFileName() string {
	now := time.Now()
	return fmt.Sprintf("backup-%d-%02d-%02d-%02d-%02d-%02d.zip",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
}


func EnsureDir(dirName string) error {
	err := os.MkdirAll(dirName, 0755)
	if err != nil {
		return err
	}
	return nil
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func GetAbsolutePath(path string) (string, error) {
	return filepath.Abs(path)
}
