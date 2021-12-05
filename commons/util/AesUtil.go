package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

// Encryption Encrypt data to []byte
func Encryption(data []byte, iv []byte, key []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("length of key must = 32")
	}

	if len(iv) != 16 {
		return nil, errors.New("length of initialization vector must = 16")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(iv) != block.BlockSize() {
		return nil, errors.New("length of the initialization vector must = length of the key cipher")
	}

	streams := cipher.NewCTR(block, iv)

	dataBytes := make([]byte, len(data))
	streams.XORKeyStream(dataBytes, data)

	return dataBytes, nil
}

// Decryption Decrypt source to dst
func Decryption(source []byte, iv []byte, key []byte) ([]byte, error) {
	return Encryption(source, iv, key)
}

// EncryptionToString Encrypt data to string
func EncryptionToString(data string, iv string, key string) (string, error) {
	by, err := Encryption(StrToBytes(data), StrToBytes(iv), StrToBytes(key))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(by), nil
}

// DecryptionForString Decrypt source to dst
func DecryptionForString(source string, iv string, key string) (string, error) {
	str, err := base64.StdEncoding.DecodeString(source)
	if err != nil {
		return "", err
	}

	str, err = Decryption(str, StrToBytes(iv), StrToBytes(key))
	if err != nil {
		return "", err
	}
	return BytesToString(str), nil
}
