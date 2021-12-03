package util

import "encoding/base64"

// Encryption TODO Encrypt data to []byte
func Encryption(data interface{}, key string) ([]byte, error) {
	return nil, nil
}

// Decryption TODO Decrypt source to dst
func Decryption(source []byte, dst interface{}, key string) error {
	return nil
}

// EncryptionToString Encrypt data to string
func EncryptionToString(data interface{}, key string) (string, error) {
	by, err := Encryption(data, key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(by), nil
}

// DecryptionForString Decrypt source to dst
func DecryptionForString(source string, dst interface{}, key string) error {
	str, err := base64.StdEncoding.DecodeString(source)
	if err != nil {
		return err
	}

	err = Decryption(str, dst, key)
	if err != nil {
		return err
	}
	return nil
}
