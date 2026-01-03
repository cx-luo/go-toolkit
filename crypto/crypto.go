// Package crypto provides cryptographic utilities
package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

// MD5 returns the MD5 hash of a string
func MD5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// SHA1 returns the SHA1 hash of a string
func SHA1(text string) string {
	hash := sha1.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// SHA256 returns the SHA256 hash of a string
func SHA256(text string) string {
	hash := sha256.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}

// SHA512 returns the SHA512 hash of a string
func SHA512(text string) string {
	hash := sha512.Sum512([]byte(text))
	return hex.EncodeToString(hash[:])
}

// MD5Bytes returns the MD5 hash of bytes
func MD5Bytes(data []byte) string {
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}

// SHA1Bytes returns the SHA1 hash of bytes
func SHA1Bytes(data []byte) string {
	hash := sha1.Sum(data)
	return hex.EncodeToString(hash[:])
}

// SHA256Bytes returns the SHA256 hash of bytes
func SHA256Bytes(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

// SHA512Bytes returns the SHA512 hash of bytes
func SHA512Bytes(data []byte) string {
	hash := sha512.Sum512(data)
	return hex.EncodeToString(hash[:])
}

// HashString returns a hash string using the specified algorithm
func HashString(text string, algorithm string) (string, error) {
	switch algorithm {
	case "md5":
		return MD5(text), nil
	case "sha1":
		return SHA1(text), nil
	case "sha256":
		return SHA256(text), nil
	case "sha512":
		return SHA512(text), nil
	default:
		return "", fmt.Errorf("unsupported algorithm: %s", algorithm)
	}
}
