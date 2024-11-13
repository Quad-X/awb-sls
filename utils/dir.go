package utils

import "os"

// Returns temporary directory path
func GetTempDir() string {
	return os.Getenv("TMP_DIR")
}
