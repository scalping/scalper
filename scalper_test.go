package main

import (

	//"fmt"
	"fmt"
	"testing"

	ok "github.com/scalping/okcoin"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {

	pub, prv, err := loadKeys()
	assert.Nil(t, err)

	api, err := ok.NewWsAPI(pub, prv)
	assert.Nil(t, err)
	err = api.Connect("btc_cny", 5)
	assert.Nil(t, err)
	defer api.Close()

	err = api.Ping(true)
	assert.Nil(t, err)

	//channels := []string{"ok_btcusd_ticker", "ok_btcusd_depth"}
	//err = api.Send(NewReq("ok_btcusd_ticker", true), NewReq("ok_btcusd_trades_v1", true))
	assert.Nil(t, err)
	err = api.Send(&ok.Req{"ok_btccny_ticker", true, nil})
	//err = api.AddChannel("ok_btcusd_trades_v1")"ok_btcusd_ticker", "ok_btcusd_depth"
	assert.Nil(t, err)
	i := 1
	for i < 4 {
		fmt.Println("TEST!")
		ret, err := api.Read()
		assert.Nil(t, err)
		fmt.Println(string(ret))
		i++

	}

	err = api.Send(ok.NewReq("ok_spotcny_userinfo", true))
	assert.Nil(t, err)
	ret, err := api.Read()
	assert.Nil(t, err)
	fmt.Println(string(ret))

	fmt.Println("----------")

	err = api.Send(&ok.Req{"ok_btccny_trades_v1", true, nil})
	assert.Nil(t, err)
	i = 1
	for i < 5 {
		fmt.Println("***", i)
		ret, err := api.ReadResponses()
		assert.Nil(t, err)
		fmt.Printf("%#v\n\n", ret)
		i++
	}
}
