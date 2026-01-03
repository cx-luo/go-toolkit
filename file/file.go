// Package file provides file operation utilities
package file

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
)

// ReadLines reads all lines from a file
func ReadLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// WriteLines writes lines to a file
func WriteLines(filePath string, lines []string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

// ReadFile reads the entire file content
func ReadFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

// WriteFile writes data to a file
func WriteFile(filePath string, data []byte) error {
	return os.WriteFile(filePath, data, 0644)
}

// Exists checks if a file or directory exists
func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// IsDir checks if the path is a directory
func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// IsFile checks if the path is a file
func IsFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// GetFileSize returns the size of a file
func GetFileSize(filePath string) (int64, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

// CopyFile copies a file from src to dst
func CopyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

// GetExt returns the file extension
func GetExt(filePath string) string {
	return filepath.Ext(filePath)
}

// GetBaseName returns the base name of a file
func GetBaseName(filePath string) string {
	return filepath.Base(filePath)
}

// GetDir returns the directory of a file path
func GetDir(filePath string) string {
	return filepath.Dir(filePath)
}
