package test

import (
	"github.com/yuyenews/Beerus/network/http/commons"
	"log"
	"testing"
)

func TestBeeSession(t *testing.T) {
	session := new(commons.BeeSession)
	session.Secret = "12345678abcdefgh09876543alnkdjfh"
	session.InitializationVector = "12345678qwertyui"
	session.Timeout = 3000

	demo := Demo{}
	demo.Name = "Beerus"
	demo.Age = 18
	demo.Height = 180

	token, err := session.CreateToken(demo)
	if err != nil {
		log.Println("Test BeeSession FAIL: " + err.Error())
		return
	}

	log.Println("token: " + token)

	demo2 := Demo{}
	err = session.RestoreToken(token, &demo2)
	if err != nil {
		log.Println("Test BeeSession FAIL: " + err.Error())
		return
	}
	if demo2.Name != demo.Name || demo2.Age != demo.Age || demo2.Height != demo.Height {
		log.Println("Test BeeSession FAIL: Failed to restore token")
		return
	}

	log.Println("Test BeeSession SUCCESS")
}

type Demo struct {
	Name   string
	Age    int
	Height uint64
}
