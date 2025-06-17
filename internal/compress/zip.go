package compress

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func CompressFolders(folders []string, outputZip string) error {
	zipFile, err := os.Create(outputZip)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, folder := range folders {
		err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			relPath, err := filepath.Rel(filepath.Dir(folder), path)
			if err != nil {
				return err
			}

			if info.IsDir() {
				_, err := zipWriter.Create(relPath + "/")
				return err
			}

			fileToZip, err := os.Open(path)
			if err != nil {
				return err
			}
			defer fileToZip.Close()

			writer, err := zipWriter.Create(relPath)
			if err != nil {
				return err
			}

			_, err = io.Copy(writer, fileToZip)
			return err
		})

		if err != nil {
			return err
		}
	}

	return nil
}