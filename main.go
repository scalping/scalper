package main

import (
	"errors"
	//"fmt"

	ini "github.com/vaughan0/go-ini"
)

func loadKeys() (pub, prv string, err error) {
	var file ini.File
	file, err = ini.LoadFile("keys.ini")
	if err != nil {
		return
	}
	var ok bool
	pub, ok = file.Get("keys", "api")
	if !ok {
		err = errors.New("api key missed")
		return
	}
	prv, ok = file.Get("keys", "secret")
	if !ok {
		err = errors.New("secret key missed")
		return
	}
	return
}
