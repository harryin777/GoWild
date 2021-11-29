package utils

import "os"

func CheckFilePath(filePath string) {
	if _, err := os.Stat(filePath); err != nil {
		_ = os.MkdirAll(filePath, 0777)
	}
}
