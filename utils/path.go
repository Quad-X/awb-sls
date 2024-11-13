package utils

import "path/filepath"

// Join multiple paths
func Join(paths ...string) string {
	return filepath.Join(paths...)
}

// Returns filename from path
func GetFilename(path string) string {
	return filepath.Base(path)
}
