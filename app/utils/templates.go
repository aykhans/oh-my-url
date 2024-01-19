package utils

import (
	"os"
	"path/filepath"
)

func GetTemplatePaths(filenames ...string) []string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	templatePath := filepath.Join(dir, "app", "templates")
	for i, filename := range filenames {
		filenames[i] = filepath.Join(templatePath, filename)
	}
	return filenames
}
