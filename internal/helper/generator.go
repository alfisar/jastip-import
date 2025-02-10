package helper

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"jastip-import/domain"
	"jastip-import/internal/errorhandler"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// func GeneratePass for generate password bcrypt
func GeneratePass(data string) (passHashed string, err domain.ErrorData) {
	passHash, errData := bcrypt.GenerateFromPassword([]byte(data), bcrypt.MinCost)

	if errData != nil {
		err = errorhandler.ErrHashing(errData)
		return
	}

	passHashed = string(passHash)
	return
}

// func Verify for verify password
func Verify(passwordHash string, password string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return
}

func TimeGenerator() string {
	timeNow := time.Now().Format(time.RFC3339)
	return timeNow
}

// Hash data string single line with method sha256
func HashSha256(data string) (result []byte) {
	keysHash := sha256.New()
	keysHash.Write([]byte(data))
	result = keysHash.Sum(nil)

	return
}

// Decode data with method base64
func Decode(s string) []byte {
	data, _ := base64.StdEncoding.DecodeString(s)
	return data
}

// Decode data with AES-256-CBC Method
func DecryptAES256CBC(keys string, data string) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
			return
		}
	}()

	key := HashSha256(keys) // 32 bytes for AES-256

	encryptedData := data
	encryptedDataByte := Decode(encryptedData)
	if len(encryptedDataByte) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := encryptedDataByte[:aes.BlockSize]
	encryptedDataByte = encryptedDataByte[aes.BlockSize:]
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(encryptedDataByte) < aes.BlockSize {
		return "", fmt.Errorf("encryptedDataByte too short")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(encryptedDataByte, encryptedDataByte)

	// Unpad the decrypted data
	return string(Unpad(encryptedDataByte)), nil
}

// Encryption data with AES-256-CBC method
func EncryptAES256CBC(keys string, data []byte) (string, error) {

	private_key := HashSha256(keys)
	data = pad(data)

	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}

	block, err := aes.NewCipher(private_key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, len(data))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, data)

	encryptedData := append(iv, ciphertext...)
	return base64.StdEncoding.EncodeToString(encryptedData), nil
}

// Adding padding to data using PKCS#7
func pad(data []byte) []byte {
	padding := aes.BlockSize - (len(data) % aes.BlockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// Remove the padding data
func Unpad(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
