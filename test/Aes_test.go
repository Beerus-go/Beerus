package test

import (
	"github.com/yuyenews/Beerus/commons/util"
	"log"
	"testing"
)

func TestAes(t *testing.T) {

	key := "12345678abcdefgh09876543alnkdjfh"
	iv := "hellowordwertyu8"

	data := "helloWord"

	res, err := util.EncryptionToString(data, iv, key)

	if err != nil {
		log.Println(err.Error())
	}

	res, err = util.DecryptionForString(res, iv, key)
	if err != nil {
		log.Println("Aes Test FAIL: " + err.Error())
	}
	if data != res {
		log.Println("Aes Test FAIL: Failed to decrypt")
	} else {
		log.Println("Aes Test SUCCESS")
	}
}
