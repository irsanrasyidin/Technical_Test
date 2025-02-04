package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

type AppError struct {
	ErrorCode    int
	ErrorMessage string
}

func (appErr *AppError) Error() string {
	return fmt.Sprintf("%v - %v", appErr.ErrorCode, appErr.ErrorMessage)
}

func generateSalt(size int) ([]byte, error) {
	salt := make([]byte, size)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func HashPassword(password string) (string, error) {
	salt, err := generateSalt(16) 
	if err != nil {
		return "", err
	}

	time := uint32(3)           
	memory := uint32(64 * 1024) 
	threads := uint8(2)         
	keyLen := uint32(32)        

	hash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLen)

	// Encode hasilnya dalam Base64
	saltBase64 := base64.StdEncoding.EncodeToString(salt)
	hashBase64 := base64.StdEncoding.EncodeToString(hash)

	// Format standar penyimpanan Argon2id
	encodedHash := fmt.Sprintf("$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s", memory, time, threads, saltBase64, hashBase64)

	return encodedHash, nil
}
