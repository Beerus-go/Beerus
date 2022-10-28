package commons

import (
	"errors"
	"github.com/Beerus-go/Beerus/commons/util"
	"strconv"
	"strings"
	"time"
)

// BeeSession Session Management, Based on the aes algorithm
// The basic principle of creating a token is to convert the data into a json string, then splice a timeout to the end, then perform aes encryption, and then convert the encrypted cipher text to base64 output
// To restore the token, reverse the creation process, first convert the base64 string to cipher text, then decrypt the cipher text with aes, after decryption, get a string with a timeout at the end, split the string into json and timeout, and determine whether the timeout is over, if not, convert the json string to the specified type of data.
type BeeSession struct {
	Timeout              int64
	Secret               string
	InitializationVector string
}

// CreateToken Create a token based on the parameters passed in
// Parameters must be of type struct
func (bs BeeSession) CreateToken(data interface{}) (string, error) {
	err := bs.validVariables()
	if err != nil {
		return "", err
	}

	if bs.Timeout <= 0 {
		bs.Timeout = 86400000
	}

	// Converting data to json strings
	jsonStr, err := util.ToJSONString(data)
	if err != nil {
		return "", err
	}

	// Splice the timeout at the end of the json string
	timeOut := bs.Timeout + time.Now().UnixMilli()
	jsonStr = jsonStr + CarriageReturn + strconv.FormatInt(timeOut, 10)

	// AES encryption and conversion to base64 return
	dat, err := util.EncryptionToString(jsonStr, bs.InitializationVector, bs.Secret)
	if err != nil {
		return "", err
	}
	return dat, nil
}

// RestoreToken Restore the token to the original data
// The second parameter must be a pointer of type struct
func (bs BeeSession) RestoreToken(token string, dst interface{}) error {
	if token == "" {
		return errors.New("token is incorrect")
	}

	err := bs.validVariables()
	if err != nil {
		return err
	}

	// Restore the base64 and decrypt it to the original data by AES (json spliced with timeout)
	dstStr, err := util.DecryptionForString(token, bs.InitializationVector, bs.Secret)
	if err != nil {
		return err
	}

	// Splitting data into json and timeout
	index := strings.LastIndex(dstStr, CarriageReturn)
	if index < 0 {
		return errors.New("token is incorrect")
	}

	jsonStr := dstStr[:index]
	timeOutStr := dstStr[(index + len(CarriageReturn)):]

	timeOut, errMsg := strconv.ParseInt(timeOutStr, 10, 64)
	if errMsg != nil {
		return errors.New("token is incorrect" + errMsg.Error())
	}

	// If the timeout expires, the user is prompted
	if time.Now().UnixMilli() > timeOut {
		return errors.New("token is no longer valid")
	}

	// If it doesn't time out, the json string is converted to the specified struct
	err = util.ParseStruct(jsonStr, dst)
	if err != nil {
		return err
	}

	return nil
}

// validVariables Verify that the secret key and initialization vector are empty
func (bs BeeSession) validVariables() error {
	if bs.Secret == "" {
		return errors.New("you need to set a secret key first before you can use BeeSession")
	}

	if bs.InitializationVector == "" {
		return errors.New("you need to set a initialization vector first before you can use BeeSession")
	}

	return nil
}
