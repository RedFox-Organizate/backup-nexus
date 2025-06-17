package backup

import (
	"os"
)

type LocalSource struct {
	Folders []string
}

func NewLocalSource(folders []string) *LocalSource {
	return &LocalSource{Folders: folders}
}

func (ls *LocalSource) Validate() []string {
	var valid []string
	for _, path := range ls.Folders {
		info, err := os.Stat(path)
		if err == nil && info.IsDir() {
			valid = append(valid, path)
		}
	}
	return valid
}
